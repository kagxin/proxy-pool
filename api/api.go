package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kagxin/golib/web/gin/response"
)

// InitRouter 初始化路由
func InitRouter(srv *Service) *gin.Engine {
	r := gin.Default()
	proxy := r.Group("proxy")
	{
		proxy.GET("/get", func(c *gin.Context) {
			rsp, err := srv.GetOneProxy(c)
			if err != nil {
				response.JSON(c, nil, err)
				return
			}
			response.JSON(c, OK, rsp)
		})
		proxy.GET("/getall", func(c *gin.Context) {
			rsp, err := srv.GetAllProxy(c)
			if err != nil {
				response.JSON(c, nil, err)
				return
			}
			response.JSON(c, OK, rsp)
		})
		proxy.DELETE("/delete/:id", func(c *gin.Context) {
			type ProxyID struct {
				ID int `uri:"id" binding:"required"`
			}
			var pID = &ProxyID{}
			if err := c.ShouldBindUri(pID); err != nil {
				response.JSON(c, nil, RequestError)
				return
			}
			if err := srv.DeleteOneProxy(c, pID.ID); err != nil {
				response.JSON(c, nil, err)
				return
			}
			response.JSON(c, OK, nil)
		})
	}

	return r
}
