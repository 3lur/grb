package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Manager struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (m *Manager) SignIn(c *gin.Context) {
	var json Manager
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}
