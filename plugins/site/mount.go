package site

import "github.com/gin-gonic/gin"

// Mount mount web points
func (p *Plugin) Mount(rt *gin.Engine) {

	rt.GET("/", p.Wrap.HTML("site/home", p.getHome))
	rt.GET("/dashboard", p.Jwt.MustSignInMiddleware, p.Wrap.HTML("site/dashboard", p.getDashboard))
	rt.GET("/install", p.mustDatabaseEmpty, p.Wrap.HTML("form", p.getInstall))
	rt.POST("/install", p.mustDatabaseEmpty, p.Wrap.FORM(&fmInstall{}, p.postInstall))

	ag := rt.Group("/admin", p.Jwt.MustAdminMiddleware)

	ag.GET("/users", p.Wrap.HTML("site/admin/users/index", p.indexAdminUsers))

	ag.GET("/locales", p.Wrap.HTML("site/admin/locales/index", p.getAdminLocales))
	ag.GET("/locales/edit", p.Wrap.HTML("form", p.editAdminLocales))
	ag.POST("/locales", p.Wrap.FORM(&fmLocale{}, p.saveAdminLocales))
	ag.DELETE("/locales/:code", p.Wrap.JSON(p.deleteAdminLocales))

	asg := ag.Group("/site")
	asg.GET("/status", p.Wrap.HTML("site/admin/status", p.getAdminSiteStatus))
	asg.GET("/info", p.Wrap.HTML("form", p.getAdminSiteInfo))
	asg.POST("/info", p.Wrap.FORM(&fmSiteInfo{}, p.postAdminSiteInfo))
	asg.GET("/author", p.Wrap.HTML("form", p.getAdminSiteAuthor))
	asg.POST("/author", p.Wrap.FORM(&fmSiteAuthor{}, p.postAdminSiteAuthor))
	asg.GET("/seo", p.Wrap.HTML("site/admin/seo", p.getAdminSiteSeo))
	asg.POST("/seo", p.Wrap.FORM(&fmSiteSeo{}, p.postAdminSiteSeo))
	asg.GET("/smtp", p.Wrap.HTML("site/admin/smtp", p.getAdminSiteSMTP))
	asg.POST("/smtp", p.Wrap.FORM(&fmSiteSMTP{}, p.postAdminSiteSMTP))

	rt.GET("/notices", p.Wrap.HTML("site/notices/index", p.indexNotices))
	ag.GET("/notices", p.Wrap.HTML("site/notices/manage", p.indexAdminNotices))
	ag.GET("/notices/new", p.Wrap.HTML("form", p.newNotice))
	ag.POST("/notices", p.Wrap.FORM(&fmNotice{}, p.createNotice))
	ag.GET("/notices/edit/:id", p.Wrap.HTML("form", p.editNotice))
	ag.POST("/notices/:id", p.Wrap.FORM(&fmNotice{}, p.updateNotice))
	ag.DELETE("/notices/:id", p.Wrap.JSON(p.destroyNotice))

	rt.GET("/leave-words/new", p.Wrap.HTML("auth/users/non-sign-in", p.newLeaveWord))
	rt.POST("/leave-words", p.Wrap.FORM(&fmLeaveWord{}, p.createLeaveWord))
	ag.GET("/leave-words", p.Wrap.HTML("site/leave-words/index", p.indexLeaveWords))
	ag.DELETE("/leave-words/:id", p.Wrap.JSON(p.destroyLeaveWord))

	ag.GET("/friend-links", p.Wrap.HTML("site/friend-links/index", p.indexFriendLinks))
	ag.GET("/friend-links/new", p.Wrap.HTML("form", p.newFriendLink))
	ag.POST("/friend-links", p.Wrap.FORM(&fmFriendLink{}, p.createFriendLink))
	ag.GET("/friend-links/edit/:id", p.Wrap.HTML("form", p.editFriendLink))
	ag.POST("/friend-links/:id", p.Wrap.FORM(&fmFriendLink{}, p.updateFriendLink))
	ag.DELETE("/friend-links/:id", p.Wrap.JSON(p.destroyFriendLink))

	ag.GET("/links", p.Wrap.HTML("site/links/index", p.indexLinks))
	ag.GET("/links/new", p.Wrap.HTML("form", p.newLink))
	ag.POST("/links", p.Wrap.FORM(&fmLink{}, p.createLink))
	ag.GET("/links/edit/:id", p.Wrap.HTML("form", p.editLink))
	ag.POST("/links/:id", p.Wrap.FORM(&fmLink{}, p.updateLink))
	ag.DELETE("/links/:id", p.Wrap.JSON(p.destroyLink))

	ag.GET("/cards", p.Wrap.HTML("site/cards/index", p.indexCards))
	ag.GET("/cards/new", p.Wrap.HTML("form", p.newCard))
	ag.POST("/cards", p.Wrap.FORM(&fmCard{}, p.createCard))
	ag.GET("/cards/edit/:id", p.Wrap.HTML("form", p.editCard))
	ag.POST("/cards/:id", p.Wrap.FORM(&fmCard{}, p.updateCard))
	ag.DELETE("/cards/:id", p.Wrap.JSON(p.destroyCard))
}
