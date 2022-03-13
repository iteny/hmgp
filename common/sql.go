package common

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Sql() *gorm.DB {
	return DB
}
func (c *BaseCtrl) Sql() *gorm.DB {
	return DB
}
func init() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/hmgp?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "hm_", // 表名前缀，`Article` 的表名应该是 `it_articles`
			SingularTable: true,  // 使用单数表名，启用该选项，此时，`Article` 的表名应该是 `it_article`
		},
	})
	// DB, err := gorm.Open(mysql.New(mysql.Config{
	// 	DSN:                       "root:123456@tcp(127.0.0.1:3306)/hmgp?charset=utf8&parseTime=True&loc=Local", // data source name
	// 	DefaultStringSize:         256,                                                                          // default size for string fields
	// 	DisableDatetimePrecision:  true,                                                                         // disable datetime precision, which not supported before MySQL 5.6
	// 	DontSupportRenameIndex:    true,                                                                         // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
	// 	DontSupportRenameColumn:   true,                                                                         // `change` when rename column, rename column not supported before MySQL 8, MariaDB
	// 	SkipInitializeWithVersion: false,
	// 	NamingStrategy: schema.NamingStrategy{
	// 		TablePrefix:   "it_", // 表名前缀，`Article` 的表名应该是 `it_articles`
	// 		SingularTable: true,  // 使用单数表名，启用该选项，此时，`Article` 的表名应该是 `it_article`
	// 	},
	// }), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	sqlDB, err := DB.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
