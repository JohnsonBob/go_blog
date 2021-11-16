package models

import (
	"fmt"
	"go_blog/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"strings"
	"time"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec, err := setting.Cfg.GetSection(setting.DATABASE)
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = sec.Key(setting.TYPE).String()
	dbName = sec.Key(setting.NAME).String()
	user = sec.Key(setting.USER).String()
	password = sec.Key(setting.PASSWORD).String()
	host = sec.Key(setting.HOST).String()
	tablePrefix = sec.Key(setting.TABLE_PREFIX).String()

	dsnT := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(dsnT, user, password, host, dbName)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                dbType,
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,                       // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
	})
	if err != nil {
		log.Println("数据库连接失败", err)
	}
	db.Logger = db.Logger.LogMode(logger.Info)
	mysqlDB, err := db.DB()
	if err != nil {
		log.Println(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	mysqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	mysqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	mysqlDB.SetConnMaxLifetime(time.Hour)
}

func CloseDB() {
	defer db.Distinct()
}
