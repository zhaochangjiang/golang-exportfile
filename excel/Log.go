package data

import (
	log4go "github.com/alecthomas/log4go"
)

//Log 日志写入
func Log() {

	//输出到控制台,级别为DEBUG
	log4go.AddFilter("stdout", log4go.DEBUG, log4go.NewConsoleLogWriter())

	//输出到文件,级别为DEBUG,文件名为test.log,每次追加该原文件
	log4go.AddFilter("file", log4go.DEBUG, log4go.NewFileLogWriter("test.log", false))

	//l4g.LoadConfiguration("log.xml")//使用加载配置文件,类似与java的log4j.propertites
	log4go.Debug("the time is now :%s -- %s", "213", "sad")

	defer log4go.Close() //注:如果不是一直运行的程序,请加上这句话,否则主线程结束后,也不会输出和log到日志文件
}
