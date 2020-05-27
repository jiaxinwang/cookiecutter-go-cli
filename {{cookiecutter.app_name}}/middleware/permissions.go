package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xyproto/pstore"
)

var perm *pstore.Permissions

// Init ...
func Init(dsn, db string) {
	var err error
	if perm, err = pstore.NewWithDSN(dsn, db); err != nil {
		logrus.WithError(err).Fatal()
		return
	}

	userstate := perm.UserState()
	userstate.AddUser("major", "tom", "major@tom.com")
	userstate.SetAdminStatus("major")
}

// Permission ...
func Permission(c *gin.Context) {
	if perm.Rejected(c.Writer, c.Request) {
		c.AbortWithStatus(http.StatusForbidden)
		fmt.Fprint(c.Writer, "Permission denied!")
		return
	}
	c.Next()
}
