package controller

import (
	"github.com/gin-gonic/gin"
)

type RestController struct {
	BaseController

	restinboundController *RestInboundController
	settingController     *SettingController
}

func NewRestController(g *gin.RouterGroup) *RestController {
	a := &RestController{}
	a.initRouter(g)
	return a
}

func (a *RestController) initRouter(g *gin.RouterGroup) {
	g = g.Group("/rest")
	g.Use(a.checkRestToken)

	g.GET("/", a.index)
	g.GET("/inbounds", a.inbounds)
	g.GET("/setting", a.setting)

	a.restinboundController = NewRestInboundController(g)
	a.settingController = NewSettingController(g)
}

func (a *RestController) index(c *gin.Context) {
	html(c, "index.html", "pages.index.title", nil)
}

func (a *RestController) inbounds(c *gin.Context) {
	html(c, "inbounds.html", "pages.inbounds.title", nil)
}

func (a *RestController) setting(c *gin.Context) {
	html(c, "setting.html", "pages.setting.title", nil)
}
