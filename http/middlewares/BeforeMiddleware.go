package middlewares

import (
    "github.com/AboutGoods/go-utils/log"
    "github.com/gin-gonic/gin"
)

func Log(c *gin.Context) {
    log.ResetContext()
    log.AddToContext("path", c.Request.RequestURI)
    log.AddToContext("method", c.Request.Method)
    log.AddToContext("host", c.Request.Host)
    log.Debug("Request received")
}
