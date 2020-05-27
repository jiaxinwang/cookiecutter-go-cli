package auth

import (
	"github.com/sirupsen/logrus"
	"github.com/xyproto/pstore"
)

// Perm ...
var Perm *pstore.Permissions

// Init ...
func Init(dsn, db string) {
	var err error
	if Perm, err = pstore.NewWithDSN(dsn, db); err != nil {
		logrus.WithError(err).Fatal()
		return
	}

	userstate := Perm.UserState()
	userstate.AddUser("major", "tom", "major@tom.com")
	userstate.SetAdminStatus("major")
}
