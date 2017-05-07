package forms

import "github.com/gin-gonic/gin"

func (p *Plugin) getFormCancel(c *gin.Context, l string) (gin.H, error) {
	return gin.H{}, nil
}

type fmCancel struct {
	Email string `form:"title" binding:"required,max=255"`
}

func (p *Plugin) postFormCancel(c *gin.Context, l string, o interface{}) (interface{}, error) {
	return gin.H{}, nil
}
