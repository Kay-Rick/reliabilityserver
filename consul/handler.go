package consul

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

/**
 * @Author miraclebay
 * @Date $
 * @note
 **/

var r = gin.Default()

func CommonHttpGetOperation(resp *http.Response, err error, c *gin.Context) {
	if err != nil || resp.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var jsondata interface{}
	err = json.Unmarshal(body, &jsondata)
	if err != nil {
		fmt.Println("json unmarshal failed: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, jsondata)
}

func GetAllServicesHandler(c *gin.Context) {
	//返回数据中心列表
	resp, err := http.Get("http://localhost:8500/v1/internal/ui/services?dc=dc1")
	CommonHttpGetOperation(resp, err, c)
}

func GetServiceInstances(c *gin.Context) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8500/v1/health/service/%s?dc=dc1", c.Param("name")))
	CommonHttpGetOperation(resp, err, c)
}
