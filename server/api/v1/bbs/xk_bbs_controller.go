package bbsApi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	bbsModel "github.com/flipped-aurora/gin-vue-admin/server/model/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	bbsService "github.com/flipped-aurora/gin-vue-admin/server/service/bbs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type XkBbsApi struct{}

// 根据id查询
func (xkBbsApi *XkBbsApi) GetXkBbs(c *gin.Context) {
	var xkBbs bbsModel.XkBbs
	//绑定参数
	err := c.ShouldBindQuery(&xkBbs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	xkBbsService := new(bbsService.XkBbsService)

	data, err := xkBbsService.GetXkBbs(xkBbs.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}
