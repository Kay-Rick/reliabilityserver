package healthcheck

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/api/checkhealth/ping", PingHandler)

	e.GET("/api/router/getimages", ConImageGetHandler)    //参数结构体
	e.GET("/api/router/runcontainer", DockerOtherHandler) //参数结构体
	e.POST("/api/router/createandruncontainer", CACHandler)
	e.GET("/api/router/getallcontainers", ConImageGetHandler)
	e.GET("/api/router/stopcontainer", DockerOtherHandler)
	e.GET("/api/router/removecontainer", DockerOtherHandler)
	e.GET("/api/router/removeimage", DockerOtherHandler)
	e.POST("/api/router/receiveandloadimage", RALHandler)
	e.GET("/api/router/devicestatus", DockerOtherHandler)
	e.GET("/api/router/devicedspstatus", DockerOtherHandler)
	e.GET("/api/router/devicefpgastatus", DockerOtherHandler)

}
