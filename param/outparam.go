package param

/**
 * @Author miraclebay
 * @Date 23:19 2022/4/1
 * @note
 **/

//返回的Image列表
type Image struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	ImageId    string `json:"imageid"`
	Created    string `json:"created"`
	Size       string `json:"size"`
}

type Container struct {
	ContainerId string   `json:"containerid"`
	ImageName   string   `json:"imagename"`
	Cmd         string   `json:"cmd"`
	Created     string   `json:"created"`
	Status      string   `json:"status"`
	Ports       string   `json:"ports"`
	Names       []string `json:"names"`
	Running     bool     `json:"running"`
}

type Service struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}
