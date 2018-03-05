package controllers

import (
	"github.com/gin-gonic/gin"
)

type GlobalNotiSetting struct{}

// One ...
func (ctrl GlobalNotiSetting) One(c *gin.Context) {
	c.JSON(200, gin.H{
		"doNotDisturb":         gin.H{"enabled": true, "from": 1300, "to": 2200},
		"perNotiMethodSetting": gin.H{"notiMethodId": "123", "type": "email", "identifier": "a@b.com"},
	})
}
