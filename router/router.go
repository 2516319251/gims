package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.GET("/", handler)
	return r
}

func handler(c *gin.Context) {
	// 升级协议
	conn, err := w.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "upgrader failed",
			"error": err.Error(),
		})
		return
	}

	for {
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			break
		}

		if err = conn.WriteMessage(msgType, data); err != nil {
			conn.Close()
			break
		}
	}

}
