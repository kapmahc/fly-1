package auth

import "github.com/gin-gonic/gin"

// Mount mount web points
func (p *Plugin) Mount(rt *gin.Engine) {
	ung := rt.Group("/users")
	ung.GET("/", p.Wrap.HTML("auth/users/index", p.indexUsers))
	ung.GET("/sign-up", p.Wrap.HTML("auth/users/non-sign-in", p.getUsersSignUp))
	ung.POST("/sign-up", p.Wrap.FORM(&fmSignUp{}, p.postUsersSignUp))
	ung.GET("/sign-in", p.Wrap.HTML("auth/users/non-sign-in", p.getUsersSignIn))
	ung.POST("/sign-in", p.Wrap.FORM(&fmSignIn{}, p.postUsersSignIn))
	ung.GET("/confirm", p.Wrap.HTML("auth/users/non-sign-in", p.getUsersEmail("confirm")))
	ung.GET("/confirm/:token", p.Wrap.Redirect(p.getUsersConfirmToken))
	ung.POST("/confirm", p.Wrap.FORM(&fmEmail{}, p.postUsersConfirm))
	ung.GET("/unlock", p.Wrap.HTML("auth/users/non-sign-in", p.getUsersEmail("unlock")))
	ung.GET("/unlock/:token", p.Wrap.Redirect(p.getUsersUnlockToken))
	ung.POST("/unlock", p.Wrap.FORM(&fmEmail{}, p.postUsersUnlock))
	ung.GET("/forgot-password", p.Wrap.HTML("auth/users/non-sign-in", p.getUsersEmail("forgot-password")))
	ung.POST("/forgot-password", p.Wrap.FORM(&fmEmail{}, p.postUsersForgotPassword))
	ung.GET("/reset-password/:token", p.Wrap.HTML("auth/users/non-sign-in", p.getUsersResetPassword))
	ung.POST("/reset-password", p.Wrap.FORM(&fmResetPassword{}, p.postUsersResetPassword))

	umg := rt.Group("/users", p.Jwt.MustSignInMiddleware)
	umg.GET("/info", p.Wrap.HTML("form", p.getUsersInfo))
	umg.POST("/info", p.Wrap.FORM(&fmInfo{}, p.postUsersInfo))
	umg.GET("/change-password", p.Wrap.HTML("form", p.getUsersChangePassword))
	umg.POST("/change-password", p.Wrap.FORM(&fmChangePassword{}, p.postUsersChangePassword))
	umg.GET("/logs", p.Wrap.HTML("auth/users/logs", p.getUsersLogs))
	umg.DELETE("/sign-out", p.Wrap.JSON(p.deleteUsersSignOut))
}
