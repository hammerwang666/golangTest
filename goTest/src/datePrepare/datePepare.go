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
	if _,err :=conn.Exec(userSql1,"1","admin", "admin", "123456", "2014-6-7 22:58:08", "admin", "0", "0","NORMAL");err!=nil{
		log.Fatal("insertAdmin error", err.Error())
		return err
	}
	if _,err :=conn.Exec(userSql1,"2","teacherTest1", "teacher1", "123456", "2014-6-7 22:58:08", "teacher", "0", "0","NORMAL");err!=nil{
		log.Fatal("insertTeacher1 error", err.Error())
		return err

	}

	if _,err :=conn.Exec(userSql1,"3","teacherTest2", "teacher2", "123456", "2014-6-7 22:58:08", "teacher", "0", "0","NORMAL");err!=nil{
		log.Fatal("insertTeacher2 error", err.Error())
		return err
	}

	return  nil



}

//获得老师
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
//获得管理员
func SelectAdmin()(error,[]UserStruct){
	conn,_=NewTestConn()
	u_admin:=UserStruct{}
	u_adminArr:=[]UserStruct{}
	user1:="select u_account,u_password from rcp_sys_user where type=? "
	rows,err:=conn.Query(user1,"admin")
	if err!=nil{

		log.Fatal("getAdminquery  error", err.Error())
		return err,nil
	}else if err==nil{
		for  rows.Next(){
			err=rows.Scan(&u_admin.userAccount,&u_admin.userPassword)
			if err !=nil{
				log.Fatal("GetTeacher Scan error", err.Error())
				return err,nil
			}
			u_adminArr=append(u_adminArr,u_admin)
		}
	}
	return  nil,u_adminArr


}
//删除用户
func  DeleteDate()error{
	conn,_=NewTestConn()
	defer conn.Close()
	deleteTeacher:="delete from rcp_sys_user where tid<'4'"
	deleteStore:="delete from rcp_store where   u_id<'4' "
	deleteUserStore:="delete from rcp_user_store_competence where user_tid<'6'"
	deleteStoreCat:="delete from rcp_shopcategory"
	deletecategory:="delete from rcp_category"
	deleteResorceNew:="delete from rcp_resource_new"
	delete_r_cat_res:="delete from rcp_r_cat_res"
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

	if _,err:=conn.Exec(deleteStoreCat);err!=nil{
		log.Fatal("delete  rcp_shopcategory error",err.Error())
		return err
	}
	if _,err:=conn.Exec(deletecategory);err!=nil{
		log.Fatal("delete  rcp_category error",err.Error())
		return err
	}
    if _,err:=conn.Exec(deleteResorceNew);err!=nil{
		log.Fatal("delete  rcp_resource_new error",err.Error())
		return err
	}
	if _,err:=conn.Exec(delete_r_cat_res);err!=nil{
		log.Fatal("delete  rcp_r_cat_res error",err.Error())
		return err
	}
	return nil
}


