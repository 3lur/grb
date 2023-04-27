package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type validateFunc func(interface{}, *gin.Context) map[string]string

func validate(data interface{}, rules, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}
	return govalidator.New(opts).ValidateStruct()
}

func Validate(c *gin.Context, obj interface{}, handler validateFunc) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		// todo
		return false
	}

	errs := handler(obj, c)

	if len(errs) > 0 {
		// todo
		return false
	}
	return true
}
