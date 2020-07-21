package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter(srv *Service) *gin.Engine {
	r := gin.Default()
	proxy := r.Group("proxy")
	{
		// 随机获取一个可用代理
		proxy.GET("/get", func(c *gin.Context) {
			rsp, err := srv.GetOneProxy(c)
			if err != nil {
				if err == NoFound {
					c.AbortWithStatus(http.StatusNotFound)
					return
				}
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			c.JSON(http.StatusOK, rsp)
		})
		// 获取所有可用代理
		proxy.GET("/getall", func(c *gin.Context) {
			rsp, err := srv.GetAllProxy(c)
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			c.JSON(http.StatusOK, rsp)
		})
		// 删除代理
		proxy.DELETE("/delete/:id", func(c *gin.Context) {
			type ProxyID struct {
				ID int `uri:"id" binding:"required"`
			}
			var pID = &ProxyID{}
			if err := c.ShouldBindUri(pID); err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			if err := srv.DeleteOneProxy(c, pID.ID); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			c.JSON(http.StatusOK, nil)
		})
	}

	return r
}
