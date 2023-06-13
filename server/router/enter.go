package router

import (
	bbsRouter "github.com/flipped-aurora/gin-vue-admin/server/router/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	XkBbs   bbsRouter.XkBbsRouter
}

var RouterGroupApp = new(RouterGroup)
