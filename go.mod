module go_blog

go 1.17

require (
	github.com/astaxie/beego v1.12.3
	github.com/boombuler/barcode v1.0.1
	github.com/gin-gonic/gin v1.7.4
	github.com/go-ini/ini v1.64.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/robfig/cron v1.2.0
	github.com/unknwon/com v1.0.1
	github.com/xuri/excelize/v2 v2.4.1
	gorm.io/driver/mysql v1.2.0
	gorm.io/gorm v1.22.3
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
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/richardlehane/mscfb v1.0.3 // indirect
	github.com/richardlehane/msoleps v1.0.1 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/ugorji/go/codec v1.2.6 // indirect
	github.com/xuri/efp v0.0.0-20210322160811-ab561f5b45e3 // indirect
	golang.org/x/crypto v0.0.0-20211115234514-b4de73f9ece8 // indirect
	golang.org/x/net v0.0.0-20210726213435-c6fcb2dbf985 // indirect
	golang.org/x/sys v0.0.0-20211116061358-0a5406a5449c // indirect
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
