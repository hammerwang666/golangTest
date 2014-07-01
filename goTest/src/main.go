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
		//delete   test  date
		err = datePrepare.DeleteDate()
		if err != nil {
			log.Println("delete date  fail", err)
			return

		}
		return
	}


	//get loginTeacher
	err, userDate.LogTeacherMsg = datePrepare.SelectTeacher()
	if err != nil {
		log.Fatal("GetTeacherDate err", err.Error())
		//delete   test  date
		err = datePrepare.DeleteDate()
		if err != nil {
			log.Println("delete date  fail", err)
			return

		}
		return
	}
	//get loginAdmin
	err,userDate.LoginAdminMsg=datePrepare.SelectAdmin()
	if err != nil {
		log.Fatal("GetAdminDate err", err.Error())
		//delete   test  date
		err = datePrepare.DeleteDate()
		if err != nil {
			log.Println("delete date  fail", err)
			return

		}
		return
	}

     //process1    admin  add  storeCat , cateGory
    addStoreCat:=process.AdminAddStoreCat()
	if addStoreCat == false {
		log.Println("process of add storeCat fail")
		result.ShowResult()
		//delete   test  date
		err = datePrepare.DeleteDate()
		if err != nil {
			log.Println("delete date  fail", err)
			return

		}
		return
	}


	//process 2   teacher  add   shop
	addShop := process.ProcessLoginAddShop()
	if addShop == false {
		log.Println("process of add shop fail")
		result.ShowResult()
		//delete   test  date
		err = datePrepare.DeleteDate()
		if err != nil {
			log.Println("delete date  fail", err)
			return

		}
		return
	}


	//show the result
	result.ShowResult()


	//delete   test  date
	err = datePrepare.DeleteDate()
	if err != nil {
		log.Println("delete date  fail", err)
		return

	}
}
