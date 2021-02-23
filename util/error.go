package util

import "log"

func CheckError(err error)  {
	if err!=nil {
		log.Fatalf("出现了错误",err)
	}
}
