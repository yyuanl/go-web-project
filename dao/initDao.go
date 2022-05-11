package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"web_project/config"
)

var once sync.Once

type DbConnections struct {
	UserAccountManagerDB *gorm.DB
}

var DbInstance DbConnections
var gDB *gorm.DB

// InitDb 连接数据库 单例执行
func InitDb() {
	once.Do(
		func() {
			DbInstance.UserAccountManagerDB = InitConnectionDb(GetMySqlDsn(config.GlobalConfigInstance.UserAccountManagerConfig))
		})
}

// InitConnectionDb 初始化连接数据库
func InitConnectionDb(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	return db
}

// GetMySqlDsn 组装mysql连接信息
func GetMySqlDsn(mysqlConfig config.MysqlConfig) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=5s",
		mysqlConfig.MysqlUser,
		mysqlConfig.MysqlPassword,
		mysqlConfig.MysqlHost,
		mysqlConfig.MysqlPort,
		mysqlConfig.DB)
	return dsn
}

// Begin 事务开启
func (s *DbConnections) Begin() *gorm.DB {
	return gDB.Begin()
}

// Commit 事务提交
func (s *DbConnections) Commit() *gorm.DB {
	return gDB.Commit()
}
