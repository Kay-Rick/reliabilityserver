package consul

import "github.com/gin-gonic/gin"

/**
 * @Author miraclebay
 * @Date $
 * @note
 **/

func Routers(e *gin.Engine) {
	e.GET("/api/v1/allservices", GetAllServicesHandler)
	e.GET("/api/services/:name/instance", GetServiceInstances)
	e.GET("/api/services/:name/:id/health-checks", GetServiceInstances)
}
