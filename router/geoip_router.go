package router

import (
	geoip_service "service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitTestRouter(geoip_router_group *gin.RouterGroup) {
	geoip_router_group.GET("/get_geoip", getGepipInfo)
}

func getGepipInfo(c *gin.Context) {
	var reqIP string
	typeV := 1
	var requestUrl string
	//IP := c.RemoteIP()
	reqIP, ok := c.GetQuery("ip")
	typeStr, ok2 := c.GetQuery("type")
	requestUrl, ok3 := c.GetQuery("requesturl")

	// str = c.Query("wd")
	// str := c.DefaultQuery("wd","acwing")
	// str , ok := c.GetQuery("wd")
	if !ok || reqIP == "" {
		reqIP = c.ClientIP()
	}

	if !ok2 || typeStr == "" {
		typeV = 1
	} else {
		typeV, _ = strconv.Atoi(typeStr)
		//typeV = typeVt
	}

	if !ok3 {
		requestUrl = ""
	}
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	c.JSON(200, geoip_service.GeoipInfor(reqIP, typeV, requestUrl))
}
