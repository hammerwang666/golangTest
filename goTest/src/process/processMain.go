package process

import (
	"loginTest"
	"log"
	"shopTest"
	"userDate"
	"result"
	"reflect"
	"storeCatTest"
)
//add  in  every process  begin
func AddPreTestResult(arg_processName string){
	//统计测试结果用,添加流程信息
	result.PreTotalReport.ProcessNumber++
	result.ProcessNum++
	result.PreResult.Process.Number=result.ProcessNum
	result.PreResult.Process.Name=arg_processName
}
//add in every process  end
func  DeletePreTestResult(){
	//每个流程结尾要append Result,并清空
	result.TotalReported = append(result.TotalReported, result.PreTotalReport)
	result.PreTotalReport.TotalTestPeople = 0
	result.PreTotalReport.TotalTest = 0
	result.PreTotalReport.PreTestTotalCase = 0
	result.PreTotalReport.PreTestTrueCaseNum = 0
	result.PreTotalReport.PreTestFalseCaseNum = 0
}

//add  in user cycle
func  AddPrePersonInfo(arg_personNum  int){
	userDate.UserNum = arg_personNum+1
	result.PreTotalReport.TotalTestPeople++    //每一个用户加一次
	//统计结果用的,分别有登陆用户的信息
	result.PreResult.Person.Name = userDate.LoginUserAccounce
	result.PreResult.Person.Number = userDate.UserNum

}


//process  admin add storeCat
 func AdminAddStoreCat() bool{
	 AddPreTestResult("AdminAddStoreCat")
	 for k,v:=range userDate.LoginAdminMsg{
		 //获得登陆的用户
		 userDate.LoginUserAccounce = string(reflect.ValueOf(v).Field(0).String())
		 userDate.LoginUserPassword = string(reflect.ValueOf(v).Field(1).String())
		 AddPrePersonInfo(k)
		 //测试用户登陆
		 result.PreTotalReport.TotalTest++//每增加一个test要增加一个
		 loginResult := loginTest.LoginTest()
		 if loginResult == false {
			 log.Println("loginRult false")
			 return false
		 }
		 result.CaseNum = 0  //每个test之后都要将CaseNum归零一次

		 //测试添加商店类型
		 result.PreTotalReport.TotalTest++//每增加一个test要增加一个
		 addStoreCatResult:=storeCatTest.CreatStoreCatTest()
		 if addStoreCatResult == false {
			 log.Println("add  storeCat false")
			 return false
		 }
		 result.CaseNum = 0  //每个test之后都要将CaseNum归零一次
		 //在最后要把test归零
		 result.TestNum=0
	 }
	 //流程结尾要append Result,并清空
	 DeletePreTestResult()
	 return true
 }

//process  teacher add  shop
func ProcessLoginAddShop() bool{
//统计测试结果用,添加流程信息
	AddPreTestResult("ProcessLoginAddShop")

	  for k,v:=range userDate.LogTeacherMsg {
		  //获得登陆的用户
		  userDate.LoginUserAccounce = string(reflect.ValueOf(v).Field(0).String())
		  userDate.LoginUserPassword = string(reflect.ValueOf(v).Field(1).String())
		  AddPrePersonInfo(k)


		  //测试用户登陆
		  result.PreTotalReport.TotalTest++//每增加一个test要增加一个
		  loginResult := loginTest.LoginTest()
		  if loginResult == false {
			  log.Println("loginRult false")
			  return false
		  }
		  result.CaseNum = 0  //每个test之后都要将CaseNum归零一次

		  //测试添加商店
		  result.PreTotalReport.TotalTest++//每增加一个test要增加一个
		  addStore := shopTest.CreatShopTest();
		  if addStore == false {
			  log.Println("add  false")
			  return false
		  }
		  result.CaseNum = 0//每个test之后都要将CaseNum归零一次

		  //在最后要把test归零
		  result.TestNum=0
	  }

	//流程结尾要append Result,并清空
	DeletePreTestResult()
	return true

}




