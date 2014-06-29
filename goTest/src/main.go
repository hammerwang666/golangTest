package main

import (
	"datePrepare"
	"log"
	"userDate"
    "process"
    "result"
)

func main() {

	 var err error
//insert  loginUserMsg
	 err=datePrepare.InnitData()
	 if err!=nil{
	  log.Fatal("InnitDate err",err.Error())
		return
	  }


//get loginUserMsg
	err,userDate.UserMsg=datePrepare.SelectTeacher();
	if err!=nil{
		log.Fatal("GetUserDate err",err.Error())
	     return
	}


//process 1
     addShop:=process.ProcessLoginAddShop()
	  if addShop==false{
		  log.Println("process of add shop fail")
		  result.ShowResult()
		  return
	  }

//process 2
	loginPro:=process.ProcessLogin()
	if loginPro==false{
		log.Println("process of login fail")
		result.ShowResult()
		return
	}




//show the result
result.ShowResult()


//delete   test  date
	err=datePrepare.DeleteTeacher()
	if err!=nil{
		log.Println("delete date  fail",err)
		return

	}
}
