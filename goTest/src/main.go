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
	err = datePrepare.InnitData()
	if err != nil {
		log.Fatal("InnitDate err", err.Error())
		return
	}


	//get loginTeacher
	err, userDate.LogTeacherMsg = datePrepare.SelectTeacher()
	if err != nil {
		log.Fatal("GetTeacherDate err", err.Error())
		return
	}
	//get loginAdmin
	err,userDate.LoginAdminMsg=datePrepare.SelectAdmin()
	if err != nil {
		log.Fatal("GetAdminDate err", err.Error())
		return
	}

     //process1    admin  add  storeCat
    addStoreCat:=process.AdminAddStoreCat()
	if addStoreCat == false {
		log.Println("process of add storeCat fail")
		result.ShowResult()
		//delete   test  date
		err = datePrepare.DeleteTeacher()
		if err != nil {
			log.Println("delete date  fail", err)
			return

		}
		return
	}


	//process 2
	addShop := process.ProcessLoginAddShop()
	if addShop == false {
		log.Println("process of add shop fail")
		result.ShowResult()
		//delete   test  date
		err = datePrepare.DeleteTeacher()
		if err != nil {
			log.Println("delete date  fail", err)
			return

		}
		return
	}


	//show the result
	result.ShowResult()


	//delete   test  date
	err = datePrepare.DeleteTeacher()
	if err != nil {
		log.Println("delete date  fail", err)
		return

	}
}
