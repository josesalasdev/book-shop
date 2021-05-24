// controllers.ping

package controllers

import (
	"github.com/gin-gonic/gin"
)

type pingMessage struct {
	Message string `json:message`
}

// Ping godoc
// @Summary Health api.
// @Description Health Api.
// @Accept  json
// @Produce  json
// @Success 200 {object} pingMessage
// @Router /ping/ [get]
func Ping(c *gin.Context) {
	c.JSON(
		200,
		pingMessage{"pong"},
	)
}
