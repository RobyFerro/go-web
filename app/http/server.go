package http

import (
	"github.com/RobyFerro/go-web/app/config"
	"github.com/RobyFerro/go-web/exception"
	"go.uber.org/dig"
	"log"
	"net"
	"net/http"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

var ServiceContainer *dig.Container
var appConf config.Conf

// Start HTTP server
func StartServer(sc *dig.Container) {
	ServiceContainer = sc

	if err := sc.Invoke(func(s *http.Server, conf config.Conf) {
		appConf = conf
		webListener, _ := net.Listen("tcp4", ":"+strconv.Itoa(conf.Server.Port))

		changeRunningUser()
		if appConf.Server.Ssl {
			if httpErr := s.ServeTLS(webListener, appConf.Server.SslCert, appConf.Server.SslKey); httpErr != nil {
				exception.ProcessError(httpErr)
			}
		} else {
			if httpsErr := s.Serve(webListener); httpsErr != nil {
				exception.ProcessError(httpsErr)
			}
		}
	}); err != nil {
		exception.ProcessError(err)
	}
}

// Change running user
// This method works only on Linux systems
// If you'd like to run go-web on Windows or OSX system you should avoid the following code
func changeRunningUser() {
	var numUID int
	var numGID int

	isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
	if len(appConf.Server.RunUser) > 0 && len(appConf.Server.RunGroup) > 0 {
		// Check the way UID is written (digits or string)
		if strings.IndexFunc(appConf.Server.RunUser, isNotDigit) == 0 {
			// If UID is a string, we lookup it to an int
			_uid, _ := user.Lookup(appConf.Server.RunUser)

			if _uid == nil {
				return
			}

			if numUID, err := strconv.Atoi(_uid.Uid); err != nil {
				exception.ProcessError(err)
			} else {
				changeUID(numUID)
			}

		} else {
			numUID, _ = strconv.Atoi(appConf.Server.RunUser)
		}

		// Check the way GID is written (digits or string)
		if strings.IndexFunc(appConf.Server.RunGroup, isNotDigit) == 0 {
			// If UID is a string, we lookup it to an int
			_gid, _ := user.LookupGroup(appConf.Server.RunGroup)

			if _gid == nil {
				return
			}

			if numGID, err := strconv.Atoi(_gid.Gid); err != nil {
				exception.ProcessError(err)
			} else {
				changeGID(numGID)
			}

		} else {
			numGID, _ = strconv.Atoi(appConf.Server.RunGroup)
		}

		log.Printf("Changing running user to %d:%d\n", uint32(numUID), uint32(numGID))
	}
}

// Change user ID
func changeUID(uid int) {
	if syscall.Getuid() == uid {
		return
	}

	if err := syscall.Setuid(uid); err != nil {
		exception.ProcessError(err)
	}
}

// Change group ID
func changeGID(gid int) {
	if syscall.Getgid() == gid {
		return
	}

	if err := syscall.Setgid(gid); err != nil {
		exception.ProcessError(err)
	}
}
