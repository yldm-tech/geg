/**
 * @Author: Evan<suzukaze.hazuki2020@gmail.com>
 * @Description: auth认证中间件
 * @File:  auth
 * @Version: 1.0.0
 * @Date: 2021/11/30 14:50
 */

package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		//token := c.GetHeader("Authorization")
		//if len(token) == 0 {
		//	log.Println("auth check fail: token is not found, please check Authorization on your header.")
		//	model.ResFailure(c, http.StatusUnauthorized, "auth check fail: token is not found, please check Authorization on your header")
		//	c.Abort()
		//	return
		//}

		if true {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, "auth check fail: token is error, please check Authorization on your header")
		}
		c.Abort()
	}
}
