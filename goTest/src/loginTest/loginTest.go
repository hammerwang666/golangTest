package loginTest

import (
	"log"
	"commonFun"
	"net/url"
	"userDate"
	"result"
)

func LoginTest() bool {
	//增加测试结果,一个测试写一个
	result.TestNum++
	result.PreResult.Test.Name="LoginTest"
	result.PreResult.Test.Number=result.TestNum
	//Login url
	userParam := url.Values{}
	userParam.Set("username",userDate.LoginUserAccounce)
	userParam.Set("password", userDate.LoginUserPassword)
	err, date1 := commonFun.HttpUrlFunc("POST", "/loginHandle1", userParam)
	if err != nil {
		log.Println("login err",err)
		return   false
	}
	result := commonFun.Expected(date1.Code, 0)
	if result == false {
		log.Println("Expected login  fail ",result)
		return  false
	}

	return true

}
