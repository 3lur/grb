package request

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"github.com/zwhypls/grb/pkg/response"
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
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确")
		return false
	}

	errs := handler(obj, c)

	if len(errs) > 0 {
		response.ValidationError(c, errs)
		return false
	}
	return true
}
