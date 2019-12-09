package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

func init() {

	//logfileLoc := beego.AppConfig.DefaultString("LogfileLoc", "/opt/logs/ylm-service/app.log")
	logfileLoc := "/opt/logs/ylm-service/test.log"
	//loglevel:=beego.AppConfig.DefaultString("loglevel", "LogLevelInfo")
	//BeegoLogLevel, XormLogLevel, _ = utils.ConvertNameToLogLevel("info")

	// 日志配置
	//isResetLog := beego.AppConfig.DefaultBool("ResetLog", false)
	//if isResetLog {
	//	// 将console打印去掉
	//	logs.Reset()
	//}
	logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"%s"}`, logfileLoc))
	logs.EnableFuncCallDepth(true)
	// 对log进行了一层封装 0 对应log.go中的writeMsg函数内部。1 对应log.XXX的函数内部。 如果是使用logs而不是logger对象来调用，那么2对应的是logs间接层函数。
	// 如果是使用logs那么3才是对应真正的函数。如果对logs进行了一次封装，那么 4 才对应真正调用出错的函数内部，所以这里设置的是4
	// 整个程序内部要统一，否则这个机制就不灵了。1) 统一使用logs 2) 统一封装一层调用
	logs.SetLogFuncCallDepth(2)
	logs.Async()

	//drivers.Engine.SetLogger(utils.GoLog2BeegoLog{})
	//drivers.Engine.Logger().SetLevel(XormLogLevel)
	//drivers.Engine.ShowSQL()
}
