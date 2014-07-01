package categoryTest
import (
	"result"
	"net/url"
	"strconv"
	"userDate"
	"commonFun"
	"log"
	"datePrepare"
	"database/sql"
)
var conn  *sql.DB
func AddCategoryTest()bool{
	//增加测试结果,一个测试写一个
	result.TesDetailResult("AddCategoryTest")
	//case 1
	values := url.Values{
	"Tid":{""},
	"Name":{"商品测试类型"+strconv.Itoa(userDate.UserNum)},
	"Pid":{"0"},
	"Action":{"create"},
	"Sort":{""},
    }

	callbackDate:=&commonFun.ResponseStruct{}
	err,_:=commonFun.HttpUrlFunc("POST","/createUpdateCategory",values,callbackDate)
	if err!=nil{
		log.Println(err)
		return  false
	}
	result:=commonFun.Expected(callbackDate.Code, 0)
	if result == false {
		return false
	}

	return true
}


func GetCategoryId()(error,int){
	conn,_=datePrepare.NewTestConn()
	var categoryId  int
	getCategoryIdSql:="select tid from rcp_category where name=?"
	rows,err:=conn.Query(getCategoryIdSql,"商品测试类型1")
	if err!=nil{
		log.Fatal("getCategoryId  query error",err.Error())
		return err,-1
	}else{
       for rows.Next(){
		   err=rows.Scan(&categoryId)
		   if err!=nil{
			   log.Fatal("getCategoryId  scan error",err.Error())
			   return err,-1
		   }
	   }
	}
	return nil,categoryId
}
