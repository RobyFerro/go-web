package conf

import gwf "github.com/RobyFerro/go-web-framework"

// This is the main configuration of Go-Web
// You can implement this method if wanna implement more configuration.
type Conf struct {
	gwf.Conf
	Name string
}

var AppConfiguration = Conf{
	Name: "Your first app",
}
