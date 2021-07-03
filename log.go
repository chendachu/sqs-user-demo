package main

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var (
	Logger    = logrus.New()
	ErrLogger = logrus.New()
)

func init() {
	Logger.SetReportCaller(true)
	ErrLogger.SetReportCaller(true)

	// 日期格式化
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	Logger.SetFormatter(customFormatter)
	ErrLogger.SetFormatter(customFormatter)

	logger := &lumberjack.Logger{
		// 日志输出文件路径
		Filename: "D:\\log\\logrus-demo.log",
		// 日志文件最大 size, 单位是 MB
		MaxSize: 500, // megabytes
		// 最大过期日志保留的个数
		MaxBackups: 3,
		// 保留过期文件的最大时间间隔,单位是天
		MaxAge: 3, //days
		// 是否需要压缩滚动日志, 使用的 gzip 压缩
		Compress: true, // disabled by default
	}
	w1 := io.MultiWriter(logger, os.Stdout)
	Logger.SetOutput(w1)


	errLogger := &lumberjack.Logger{
		// 日志输出文件路径
		Filename: "D:\\log\\logrus-demo-err.log",
		// 日志文件最大 size, 单位是 MB
		MaxSize: 500, // megabytes
		// 最大过期日志保留的个数
		MaxBackups: 200,
		// 保留过期文件的最大时间间隔,单位是天
		MaxAge: 3, //days
		// 是否需要压缩滚动日志, 使用的 gzip 压缩
		Compress: true, // disabled by default
	}
	w2 := io.MultiWriter(errLogger, os.Stdout)
	ErrLogger.SetOutput(w2)
}
