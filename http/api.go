package http

import (
	"net/http"
	"proxy-pool/stroage"

	"github.com/gin-gonic/gin"
)

// InitHttp 初始化路由
func InitHttp(s stroage.Stroage) *gin.Engine {
	r := gin.Default()
	proxy := r.Group("proxy")
	{
		// 随机获取一个可用代理
		proxy.GET("/get", func(c *gin.Context) {
			rsp, err := s.Get(c)
			if err != nil {
				if err == stroage.ErrNoFound {
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
			rsp, err := s.GetAll(c)
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			c.JSON(http.StatusOK, rsp)
		})
		// 删除代理
		proxy.DELETE("/delete/:proxy", func(c *gin.Context) {
			var person struct {
				Proxy string `uri:"proxy" binding:"required"`
			}
			if err := c.ShouldBindUri(&person); err != nil {
				c.JSON(400, gin.H{"msg": err.Error()})
				return
			}
			if err := s.Delete(c, person.Proxy); err != nil {
				if err == stroage.ErrNoFound {
					c.AbortWithStatus(http.StatusNotFound)
					return
				}
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			c.JSON(http.StatusOK, nil)
		})
	}

	return r
}
