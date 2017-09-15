package logger

import (
	"../../model"
	"go.uber.org/zap"
)

var newlogger *zap.Logger

func Init(){
	logger,err := zap.NewProduction()
	if err != nil {
		panic(err)
		return
	}
	newlogger=logger
}

func Debug(funcName,msg string){
	newlogger.Debug(msg)
}

func Info(funcName,msg string){
	newlogger.Info(msg)
	logInf:=model.LogInfo{
		Level:"Info",
		FuncName:funcName,
		Content:msg,
	}
	logInf.Insert()
}

func Warm(funcName,msg string){
	newlogger.Warn(msg)
	logInf:=model.LogInfo{
		Level:"Warm",
		FuncName:funcName,
		Content:msg,
	}
	logInf.Insert()
}

func Error(funcName,msg string){
	newlogger.Error(msg)
	logInf:=model.LogInfo{
		Level:"Error",
		FuncName:funcName,
		Content:msg,
	}
	logInf.Insert()
}




