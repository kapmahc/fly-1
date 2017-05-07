package forms

import "github.com/gin-gonic/gin"

// Mount mount web points
func (p *Plugin) Mount(rt *gin.Engine) {
	fg := rt.Group("/forms")
	fg.GET("/", p.Wrap.HTML("forms/index", p.indexForms))
	fg.GET("/apply/:id", p.Wrap.HTML("form", p.getFormApply))
	fg.POST("/apply/:id", p.Wrap.FORM(&fmApply{}, p.postFormApply))
	fg.GET("/cancel/:id", p.Wrap.HTML("form", p.getFormCancel))
	fg.POST("/cancel/:id", p.Wrap.FORM(&fmCancel{}, p.postFormCancel))

	ag := fg.Group("/", p.Jwt.MustAdminMiddleware)
	ag.GET("/manage", p.Wrap.HTML("forms/manage", p.indexForms))
	ag.GET("/report/:id", p.Wrap.HTML("forms/report", p.getFormReport))
	ag.GET("/export/:id", p.Wrap.Do(p.getFormExport))
	ag.GET("/new", p.Wrap.HTML("form", p.newForm))
	ag.POST("/", p.Wrap.FORM(&fm{}, p.createForm))
	ag.GET("/edit/:id", p.Wrap.HTML("form", p.editForm))
	ag.POST("/edit/:id", p.Wrap.FORM(&fm{}, p.updateForm))
	ag.DELETE("/:id", p.Wrap.JSON(p.destroyForm))
}
