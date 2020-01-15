package controller

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
)

type UserCheck struct {
	BaseController
}

//
func (c *UserCheck) Check() {

	uid := os.Getuid()
	gid := os.Getgid()

	uidString, _ := user.LookupId(strconv.Itoa(uid))
	gidString, _ := user.LookupGroupId(strconv.Itoa(gid))
	outString := fmt.Sprintf("You are running as user \"%s\" (%d) with group \"%s\" (%d)", uidString.Name, uid, gidString.Name, gid)

	_, _ = c.Response.Write([]byte(outString))

}
