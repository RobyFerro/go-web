package http

import (
	"go.uber.org/dig"
	"ikdev/go-web/app/config"
	"ikdev/go-web/exception"
	"net"
	"net/http"
	"strconv"
)

//#include <unistd.h>
//#include <errno.h>
import (
	"C"
)

var ServiceContainer *dig.Container
var appConf config.Conf

// Start HTTP server
func StartServer(sc *dig.Container) {
	ServiceContainer = sc

	err := sc.Invoke(func(s *http.Server, conf config.Conf) {
		// Declaring global app configuration
		appConf = conf
		webListener, _ := net.Listen("tcp4", ":"+strconv.Itoa(conf.Server.Port))

		// Disabled by default because is compatible only with Linux OS
		// changeRunningUser()

		if appConf.Server.Ssl {
			if httpErr := s.ServeTLS(webListener, appConf.Server.SslCert, appConf.Server.SslKey); httpErr != nil {
				exception.ProcessError(httpErr)
			}

		} else {
			if httpsErr := s.Serve(webListener); httpsErr != nil {
				exception.ProcessError(httpsErr)
			}
		}

	})

	if err != nil {
		exception.ProcessError(err)
	}
}

// Enable this method if you'll need to change the Go-Web running user.
// This option is commented because is compatible only with linux system
/*func changeRunningUser() {
	var numUID int
	var numGID int

	isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
	if len(appConf.Server.RunUser) > 0 && len(appConf.Server.RunGroup) > 0 {

		// Check the way UID is written (digits or string)
		if strings.IndexFunc(appConf.Server.RunUser, isNotDigit) == 0 {

			// If UID is a string, we lookup it to an int
			_uid, _ := user.Lookup(appConf.Server.RunUser)
			numUID, _ = strconv.Atoi(_uid.Uid)

		} else {
			numUID, _ = strconv.Atoi(appConf.Server.RunUser)

		}

		// Check the way GID is written (digits or string)
		if strings.IndexFunc(appConf.Server.RunGroup, isNotDigit) == 0 {

			// If UID is a string, we lookup it to an int
			_gid, _ := user.LookupGroup(appConf.Server.RunGroup)
			numGID, _ = strconv.Atoi(_gid.Gid)
		} else {
			numGID, _ = strconv.Atoi(appConf.Server.RunGroup)

		}

		// Check if running os is linux
		// You'll allow to change the running user only on Linux OS
		log.Printf("Changing running user to %d:%d\n", uint32(numUID), uint32(numGID))

		// Set running GID through setgid "C"  syscall
		cerr, errno := C.setgid(C.__gid_t(numGID))
		if cerr != 0 {
			log.Fatalln("Unable to set GID due to error:", errno)
		}

		// Set running GID through setuid "C"  syscall
		cerr, errno = C.setuid(C.__uid_t(numUID))
		if cerr != 0 {
			log.Fatalln("Unable to set UID due to error:", errno)
		}

	}

}*/
