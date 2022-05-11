package config

import (
	"encoding/json"
	"fmt"

	"github.com/BurntSushi/toml"
)

// GlobalConfigInstance 全局配置
var GlobalConfigInstance = &GlobalConfig{}

// GlobalConfig 全局配置
type GlobalConfig struct {

	// 用户账号配置
	UserAccountManagerConfig MysqlConfig
}

// MysqlConfig mysql配置
type MysqlConfig struct {
	MysqlHost     string
	MysqlPort     string
	MysqlUser     string
	MysqlPassword string
	DB            string
	TimeOut       string
	CharSet       string
}

// Init 初始化读取文件加载配置
func Init() {
	_, err := toml.DecodeFile("conf/config.toml", GlobalConfigInstance)
	if err != nil {
		panic(err)
	}
	bytes, _ := json.Marshal(GlobalConfigInstance)
	str := "--------------------------------------------------------------------------------------"
	fmt.Printf("init config susscess,config is %s\n"+str+"\n\n", string(bytes))
}
