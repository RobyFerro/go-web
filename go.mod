module github.com/RobyFerro/go-web

go 1.16

// Only for development environment
//replace github.com/RobyFerro/go-web-framework => <your local gwf path>

require (
	github.com/RobyFerro/go-web-framework v0.5.8-beta
	github.com/auth0/go-jwt-middleware v0.0.0-20200810150920-a32d7af194d1
	github.com/brianvoe/gofakeit/v4 v4.3.0
	github.com/common-nighthawk/go-figure v0.0.0-20200609044655-c4b36f998cf2
	github.com/denisenkom/go-mssqldb v0.0.0-20200206145737-bbfc9a55622e // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20200819071622-59b6a186f8dd
	github.com/go-redis/redis/v7 v7.4.0
	github.com/gorilla/sessions v1.2.1
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/now v1.1.1 // indirect
	github.com/klauspost/compress v1.10.11 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/labstack/gommon v0.3.0
	github.com/lib/pq v1.3.0 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.9.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/tidwall/pretty v1.0.1 // indirect
	go.mongodb.org/mongo-driver v1.5.1
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200826173525-f9321e4c35a6 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
	golang.org/x/tools v0.0.0-20200827010519-17fd2f27a9e3 // indirect
	gopkg.in/yaml.v2 v2.3.0
)
