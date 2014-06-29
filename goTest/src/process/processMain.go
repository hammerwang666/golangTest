package process

import (
	"loginTest"
	"log"
	"shopTest"
	"userDate"
	"reflect"
	"result"
)



func ProcessLoginAddShop() bool{
//统计测试结果用,添加流程信息
	result.PreTotalReport.ProcessNumber++
	result.ProcessNum++
	result.PreResult.Process.Name="ProcessLoginAddShop"
	result.PreResult.Process.Number=result.ProcessNum
	  for k,v:=range userDate.UserMsg {
		  result.PreTotalReport.TotalTestPeople++	//每一个用户加一次
		  //获得登陆的用户
		  userDate.UserNum = k+1
		  userDate.LoginUserAccounce = string(reflect.ValueOf(v).Field(0).String())
		  userDate.LoginUserPassword = string(reflect.ValueOf(v).Field(1).String())

		  //统计结果用的,分别有登陆用户的信息
		 result.PreResult.Person.Name=userDate.LoginUserAccounce
		 result.PreResult.Person.Number=userDate.UserNum
         


		  //测试用户登陆
		  result.PreTotalReport.TotalTest++//每增加一个test要增加一个
		  loginResult := loginTest.LoginTest()
		  if loginResult == false {
			  log.Println("loginRult false")
			  return false
		  }
		  //每个test之后都要将CaseNum归零一次
		  result.CaseNum=0



          //测试添加商店
		  result.PreTotalReport.TotalTest++//每增加一个test要增加一个
		  addStore := shopTest.CreatShopTest();
		  if addStore == false {
			  log.Println("add  false")
			  return false
		  }
		  //每个test之后都要将CaseNum归零一次
		  result.CaseNum=0
		  //在最后要把test归零,用于用户循环用

		  result.TestNum=0
	  }

	//流程结尾要append Result,并清空
	result.TotalReported = append(result.TotalReported, result.PreTotalReport)
	result.PreTotalReport.TotalTestPeople = 0
	result.PreTotalReport.TotalTest = 0
	result.PreTotalReport.PreTestTotalCase = 0
	result.PreTotalReport.PreTestTrueCaseNum = 0
	result.PreTotalReport.PreTestFalseCaseNum = 0

	return true

}

func ProcessLogin() bool{
	//统计测试结果用,添加流程信息
	result.PreTotalReport.ProcessNumber++
	result.ProcessNum++
	result.PreResult.Process.Name="ProcessLogin"
	result.PreResult.Process.Number=result.ProcessNum
	for k,v:=range userDate.UserMsg {
		//获得登陆的用户
		userDate.UserNum = k+1
		userDate.LoginUserAccounce = string(reflect.ValueOf(v).Field(0).String())
		userDate.LoginUserPassword = string(reflect.ValueOf(v).Field(1).String())
		result.PreTotalReport.TotalTestPeople++	//每一个用户加一次
		//统计结果用的,分别有登陆用户的信息
		result.PreResult.Person.Name = userDate.LoginUserAccounce
		result.PreResult.Person.Number = userDate.UserNum



		//测试用户登陆
		result.PreTotalReport.TotalTest++//每增加一个test要增加一个
		loginResult := loginTest.LoginTest();
		if loginResult == false {
			log.Println("loginRult false")
			return false
		}
		//每个test之后都要将CaseNum归零一次
		result.CaseNum = 0
		//在最后要把test归零,用于用户循环用
		result.TestNum=0

	}
	//流程结尾要append Result,并清空
	result.TotalReported = append(result.TotalReported, result.PreTotalReport)
	result.PreTotalReport.TotalTestPeople = 0
	result.PreTotalReport.TotalTest = 0
	result.PreTotalReport.PreTestTotalCase = 0
	result.PreTotalReport.PreTestTrueCaseNum = 0
	result.PreTotalReport.PreTestFalseCaseNum = 0
	return true
}
