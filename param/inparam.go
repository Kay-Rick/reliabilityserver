package param

/**
 * @Author miraclebay
 * @Date 23:13 2022/4/1
 * @note
 **/

//CreateAndRunContainer入参
type CARInParam struct {
	ImageName     string   `json:"imagename"`
	Cmd           []string `json:"cmd"`
	Network       string   `json:"network"`
	Tdy           bool     `json:"tdy"`
	Openstdin     bool     `json:"openstdin"`
	ContainerName string   `json:"containername"`
}
