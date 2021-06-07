package v1

import (
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/auth"
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/controller/response"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/util/l"

)

// Signup ...
func Signup(c *gin.Context) {
	requestID := c.MustGet("requestID")
	param := struct {
		Name     string `valid:"alphanum,required,stringlength(6|12)"`
		Password string `valid:"stringlength(6|128)"`
		Email    string `valid:"email,optional"`
	}{}

	if err := c.ShouldBindJSON(&param); err != nil {
		response.ClientErr(c, err.Error())
		return
	}

	if result, err := govalidator.ValidateStruct(param); err != nil {
		response.ClientErr(c, err.Error())
		return
	}

	userState := auth.Perm.UserState()
	if userState.HasUser(param.Name) {
		response.ClientErr(c, "user already exists")
		return
	}
	userState.AddUser(param.Name, param.Password, param.Email)
	response.Response(c, 0, "", nil)
}