package auth

import (
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/util/logger"
	"github.com/xyproto/pstore"
)

var log *zap.SugaredLogger

func init() {
	log = logger.S.Named("auth")
}


// Perm ...
var Perm *pstore.Permissions

// Init ...
func Init(dsn, db string) {
	var err error
	if Perm, err = pstore.NewWithDSN(dsn, db); err != nil {
		log.Panic(err)
		return
	}

	userstate := Perm.UserState()
	userstate.AddUser("major", "tom", "major@tom.com")
	userstate.SetAdminStatus("major")
}
