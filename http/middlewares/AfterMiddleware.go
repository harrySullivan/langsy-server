package middlewares

import (
    "App/http/controllers"
    "App/utils"
    
    "time"
    "errors"

    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"

)

func LogLatency(c *gin.Context) {
    t := time.Now()
    c.Next()

    latency := time.Since(t)

    log.WithField("time", latency).Info("route latency")
}

func AuthRequired(c *gin.Context) {
	token := c.GetHeader("X-Auth-Token")
	if len(token) > 0 {
		userId, err := controllers.ParseOrError(token)
		if err == nil {
			c.Set("UserID", userId)
			c.Next()
			return
		}
	}


	utils.HttpError(c, 400, errors.New("No Authorized to make request"))
	return
}
