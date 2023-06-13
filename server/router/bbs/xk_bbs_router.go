package bbsRouter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/bbs"
	"github.com/gin-gonic/gin"
)

type XkBbsRouter struct {
}

func (e *XkBbsRouter) InitXkBbsRouter(Router *gin.RouterGroup) {
	//xkBbsRouter := Router.Group("xkBbs").Use(middleware.OperationRecord())
	xkBbsRouterWithoutRecord := Router.Group("bbs")
	var xkBbsApi = new(bbsApi.XkBbsApi)
	{
		xkBbsRouterWithoutRecord.GET("xkBbs", xkBbsApi.GetXkBbs) // 获取单一客户信息
	}
}
