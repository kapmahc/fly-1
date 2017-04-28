package site

import "github.com/gin-gonic/gin"

func (p *Plugin) getInstall(c *gin.Context, l string) (gin.H, error) {
	return gin.H{}, nil
}

type fmInstall struct {
	Title                string `form:"title" binding:"required"`
	SubTitle             string `form:"subTitle" binding:"required"`
	Email                string `form:"email" binding:"required,email"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Plugin) postInstall(c *gin.Context, l string, o interface{}) (interface{}, error) {
	return gin.H{}, nil
}
