package command

import (
	daemon "github.com/sevlyar/go-daemon"
	"go.uber.org/dig"
	"ikdev/go-web/config"
	"ikdev/go-web/exception"
	"ikdev/go-web/http"
	"log"
)

var appConf config.Conf

// Start Go-Web server
func RunServer(container *dig.Container) {
	http.StartServer(container)

}

func RunDaemon(container *dig.Container) {

	err := container.Invoke(func(conf config.Conf) {

		appConf = conf

		// Simple way to check is a string contains only digits

		cntxt := &daemon.Context{
			PidFileName: "storage/log/go-web.pid",
			PidFilePerm: 0754,
			LogFileName: "storage/log/go-webd.log",
			LogFilePerm: 0754,
			Umask:       027,
		}

		daemonBg, daemonError := cntxt.Reborn()

		if daemonError != nil {
			log.Fatal("Troubles to start daemon ", daemonError, daemonError.Error())
		}

		if daemonBg != nil {
			return
		}

		defer cntxt.Release()

		http.StartServer(container)

	})

	if err != nil {
		exception.ProcessError(err)
	}

}
