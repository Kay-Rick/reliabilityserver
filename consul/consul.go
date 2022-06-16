package consul

import (
	"bytes"
	"encoding/json"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"math/rand"
	"net/http"
	"rs/config"
	param2 "rs/param"
	"time"
)

var smsMap, consulServerAddress, nodeStatus, reliability = config.ConfigInit()

var downedServices map[string]bool = make(map[string]bool)

//服务健康检查没通过的次数
var servicedownedtimes map[string]byte = make(map[string]byte)

//健康检查过程中需要重新启动的服务，这些服务已经在连续三轮健康检查中被标记为critical
var restartServiceChecks map[string]consulapi.AgentCheck = make(map[string]consulapi.AgentCheck)

func GetReliability() bool {
	return reliability
}

func ConsulClientInit() *consulapi.Client {
	var config = consulapi.DefaultConfig()
	fmt.Println(consulServerAddress)
	config.Address = consulServerAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Panicln("consul client error: ", err)
	}
	return client
}

func StatusCheck(client *consulapi.Client) {
	//从Agent获取所有的服务
	services, err := client.Agent().Services()
	if err != nil {
		fmt.Println("client agent get services error!")
		return
	}

	//serviceLocation指定一个服务在哪些节点上运行，interface指代的类型是map[string]方便查找和删除， key为服务name非ID
	serviceLocation := make(map[string]map[string]bool)

	//获取健康检查情况
	checks, err := client.Agent().Checks()
	if err != nil {
		fmt.Println("client agent get services check error!")
		return
	}

	/**
	将所有节点的状态标记为正常:true
	*/
	for ip, _ := range nodeStatus {
		nodeStatus[ip] = true
	}

	for _, check := range checks {
		/**
		记录相同服务名运行的所有IP，一个节点不能运行两个相同服务名的服务
		*/
		serviceId := check.ServiceID
		serviceName := check.ServiceName
		ip := services[serviceId].Address
		if _, ok := serviceLocation[serviceName]; !ok {
			serviceLocation[serviceName] = make(map[string]bool)
		}
		serviceLocation[serviceName][ip] = true

		//如果该服务已经被重启过了，则不会再加入重启列表
		if _, ok := downedServices[serviceId]; check.Status != "passing" && !ok {
			//服务异常，健康检查没通过
			nodeStatus[ip] = false //将节点状态标记为下线
			if _, ok := servicedownedtimes[serviceId]; !ok {
				servicedownedtimes[serviceId] = 0
			} else if servicedownedtimes[serviceId] < 3 {
				servicedownedtimes[serviceId] += 1
			}
			if servicedownedtimes[serviceId] == 3 {
				//连续三次健康检查未通过，将该服务加入重启服务列表，并从servicedownedtimes map中删除，防止下次重新启动
				restartServiceChecks[serviceId] = *check
			}
		}
	}

	rand.Seed(time.Now().Unix())
	var restartedServices []string

	if len(restartServiceChecks) != 0 {
		fmt.Printf("\n---------------------------------------------------\n")
	}

	for sid, rsc := range restartServiceChecks {
		fmt.Println("异常服务ID: " + sid)
		serviceName := rsc.ServiceName

		//候选节点
		var nodesCandidate []string
		for ip, status := range nodeStatus {
			//当该节点运行正常并且没有相同服务名的服务在运行,则将该节点加入候选节点列表
			if _, ok := serviceLocation[serviceName][ip]; !ok && status {
				nodesCandidate = append(nodesCandidate, ip)
			}
		}

		//如果此时没有正常节点，则本轮健康检查退出
		if len(nodesCandidate) == 0 {
			fmt.Println("Warning: the candidate nodes is none, health check exit")
			continue
		}

		//随机选择一个正常节点
		index := rand.Intn(len(nodesCandidate))
		chooseIP := nodesCandidate[index]
		param := param2.CARInParam{
			ImageName:     smsMap[serviceName].Image,
			Cmd:           []string{smsMap[serviceName].Cmd, chooseIP, "1"},
			Network:       "host",
			Tdy:           true,
			Openstdin:     true,
			ContainerName: "",
		}
		fmt.Printf("发送数据：%v\n", param)
		bytesData, _ := json.Marshal(param)

		res, err := http.Post("http://"+chooseIP+":8001/createandruncontainer", "application/json;charset=utf-8", bytes.NewBuffer([]byte(bytesData)))
		if err != nil {
			fmt.Println("http post error", err.Error())
		}
		defer res.Body.Close()
		resCode := res.StatusCode
		fmt.Println("send http request: http://"+chooseIP+":8001/createandruncontainer: status-> ", resCode)

		downedServices[rsc.ServiceID] = true //将该服务保存到已下线节点
		delete(servicedownedtimes, rsc.ServiceID)
		restartedServices = append(restartedServices, sid)
		if _, ok := serviceLocation[serviceName]; !ok {
			serviceLocation[serviceName] = make(map[string]bool)
		}
		serviceLocation[serviceName][chooseIP] = true
	}
	for _, sid := range restartedServices {
		delete(restartServiceChecks, sid)
	}
}
