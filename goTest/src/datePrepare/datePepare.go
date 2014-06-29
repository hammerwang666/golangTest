package datePrepare
import (
   "database/sql"
	"log"
)


var  conn *sql.DB

type UserStruct   struct{
  userAccount  string
  userPassword   string
}
//插入用户
func InnitData()error{
	 conn,_=NewTestConn()
	//userDate
	userSql1:="INSERT INTO rcp_sys_user (TID,U_ACCOUNT, U_NAME, U_PASSWORD, U_CREATE_TIME, TYPE, IS_TRUE_USER, IS_SON_ACCOUNT,STATUS) VALUES(?,?,?,?,?,?,?,?,?)"

	if _,err :=conn.Exec(userSql1,"4","teacherTest1", "teacher1", "123456", "2014-6-7 22:58:08", "teacher", "0", "0","NORMAL");err!=nil{
		log.Fatal("insertTeacher1 error", err.Error())
		return err

	}

	if _,err :=conn.Exec(userSql1,"5","teacherTest2", "teacher2", "123456", "2014-6-7 22:58:08", "teacher", "0", "0","NORMAL");err!=nil{
		log.Fatal("insertTeacher2 error", err.Error())
		return err
	}

	return  nil



}

//获得用户
func SelectTeacher() (error,[]UserStruct){
	conn,_=NewTestConn()
	u_teacher:=UserStruct{}
	u_teacherArr:=[]UserStruct{}
	user1:="select u_account,u_password from rcp_sys_user where type=? "
    rows,err:=conn.Query(user1,"teacher")
   if err!=nil{

	   log.Fatal("getteacher query  error", err.Error())
	   return err,nil
   }else if err==nil{
	   for  rows.Next(){
		   err=rows.Scan(&u_teacher.userAccount,&u_teacher.userPassword)
		   if err !=nil{
			   log.Fatal("GetTeacher Scan error", err.Error())
			   return err,nil
		   }
		   u_teacherArr=append(u_teacherArr,u_teacher)
	   }
   }
	return  nil,u_teacherArr


}

//删除用户
func  DeleteTeacher()error{
	conn,_=NewTestConn()
	defer conn.Close()
	deleteTeacher:="delete from rcp_sys_user where type='teacher'"
	deleteStore:="delete from rcp_store where u_id>'3'  and  u_id<'6' "
	deleteUserStore:="delete from rcp_user_store_competence where user_tid>'3'and user_tid<'6'"
	if _, err := conn.Exec(deleteTeacher); err != nil {
		log.Fatal("deleteTeacher error",err.Error())
		return err
	}
	if _,err:=conn.Exec(deleteStore);err!=nil{
		log.Fatal("delete rcp_store error",err.Error())
	return err
    }
	if _,err:=conn.Exec(deleteUserStore);err!=nil{
		log.Fatal("delete rcp_user_store_competence error",err.Error())
		return err
	}


	return nil
}
