package request

import "github.com/gin-gonic/gin"

type IDParam struct {
	ID string `form:"id" uri:"id"`
}

func (p *IDParam) Bind(c *gin.Context) (err error) {
	err = c.ShouldBindUri(&p)
	if p.ID != "" {
		return
	}
	err = c.ShouldBind(&p)
	return
}
