package request

import "github.com/astaxie/beego/validation"

// 表单验证基础接口
type RequestBaseI interface {
	Valid(string, map[string]interface{}) (map[string]string, bool)
}

// 表单验证基础结构体
type RequestBaseS struct{}

// 表单验证 - 开始校验
func (this RequestBaseS) Valid(method string, data map[string]interface{}) (map[string]string, bool) {
	// valid := validation.Validation{}
	// valid.Required(data["name"], "name").Message("用户名不能为空")
	// return this.Rtn(valid)
	return map[string]string{}, true
}

// 表单验证 - 返回错误信息
func (this RequestBaseS) Rtn(valid validation.Validation) (map[string]string, bool) {
	tmpMap := map[string]string{}
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			tmpMap[err.Key] = err.Message
		}
		return tmpMap, false
	}
	return tmpMap, true
}
