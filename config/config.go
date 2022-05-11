package config

//
//import (
//	"flag"
//	"fmt"
//	"os"
//
//	"git.code.oa.com/trpc-go/trpc-go"
//	"git.code.oa.com/trpc-go/trpc-go/log"
//)
//
//var (
//	buildstamp = ""
//	gitversion = ""
//	gitbranch  = ""
//	goversion  = ""
//	sysuname   = ""
//	ConfigPath = "../conf/" //配置文件路径
//)
//
////信息
//func PrintVersion(c bool) {
//	f := func(c bool, format string, a ...interface{}) {
//		if c {
//			log.Infof(format, a)
//		} else {
//			fmt.Printf(format, a)
//		}
//	}
//	f(c, "\tGitVersion: %v \n", gitversion)
//	f(c, "\tGitBranch: %v \n", gitbranch)
//	f(c, "\tBuildStamp: %v \n", buildstamp)
//	f(c, "\tGoVersion: %v \n", goversion)
//	f(c, "\tSysUname: %v \n", sysuname)
//
//}
//
//// 更改 增加 可执行程序版本信息
//func InitFlag() {
//	verOpt := flag.Bool("v", false, "Print application version")
//	ConfigPathOpt := flag.String("c", "", "config file path")
//	flag.Parse()
//	if *verOpt {
//		PrintVersion(false)
//		os.Exit(0) //输出版本信息,退出
//	}
//	//更改配置文件
//	if *ConfigPathOpt != "" {
//		ConfigPath = *ConfigPathOpt
//		trpc.ServerConfigPath = ConfigPath + "trpc_go.yaml"
//		log.Infof("change ServerConfigPath : %v \n", trpc.ServerConfigPath)
//	}
//	PrintVersion(true)
//}
