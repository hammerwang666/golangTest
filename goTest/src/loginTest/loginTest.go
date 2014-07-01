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
	result.TesDetailResult("LoginTest")
	//Login url
	userParam := url.Values{}
	userParam.Set("username",userDate.LoginUserAccounce)
	userParam.Set("password", userDate.LoginUserPassword)
	callbackDate:=&commonFun.ResponseStruct{}
	err, _ := commonFun.HttpUrlFunc("POST", "/loginHandle1", userParam,callbackDate)
	if err != nil {
		log.Println("login err",err)
		return   false
	}
	result := commonFun.Expected(callbackDate.Code, 0)
	if result == false {
		log.Println("Expected login  fail ",result)
		return  false
	}

	return true

}
