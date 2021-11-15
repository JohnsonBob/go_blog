module go_blog

go 1.17

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/go-ini/ini v1.64.0
	github.com/unknwon/com v1.0.1
	gorm.io/driver/mysql v1.1.3
	gorm.io/gorm v1.22.2
)

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/ugorji/go/codec v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20211108221036-ceb1ce70b4fa // indirect
	golang.org/x/sys v0.0.0-20211113001501-0c823b97ae02 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/JohnsonBob/go_blog/conf => ./pkg/conf
	github.com/JohnsonBob/go_blog/middleware => ./middleware
	github.com/JohnsonBob/go_blog/models => ./models
	github.com/JohnsonBob/go_blog/pkg/setting => ./setting
	github.com/JohnsonBob/go_blog/routers => ./routers
)