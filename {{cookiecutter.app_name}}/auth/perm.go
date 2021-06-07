package auth

import (
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/util/l"
	"github.com/xyproto/pstore"
)

// Perm ...
var Perm *pstore.Permissions

// Init ...
func Init(dsn, db string) {
	var err error
	if Perm, err = pstore.NewWithDSN(dsn, db); err != nil {
		l.S.Panic(err)
		return
	}

	userstate := Perm.UserState()
	userstate.AddUser("major", "tom", "major@tom.com")
	userstate.SetAdminStatus("major")
}
