package forms

import "github.com/gin-gonic/gin"

func (p *Plugin) getFormApply(c *gin.Context, l string) (gin.H, error) {
	return gin.H{}, nil
}

type fmApply struct {
	Username string `form:"username" binding:"required,max=255"`
	Email    string `form:"title" binding:"required,max=255"`
	Phone    string `form:"phone" binding:"required,max=255"`
	Value    string `form:"value" binding:"required"`
}

func (p *Plugin) postFormApply(c *gin.Context, l string, o interface{}) (interface{}, error) {
	return gin.H{}, nil
}
