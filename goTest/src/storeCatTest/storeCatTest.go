package storeCatTest

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

var conn *sql.DB
func CreatStoreCatTest()bool{
	//增加测试结果,一个测试写一个
	result.TestNum++
	result.PreResult.Test.Name="CreatStoreTest"
	result.PreResult.Test.Number=result.TestNum
	//case 1
	assStoreCat:=url.Values{}
    assStoreCat.Set("Name","商品测试类型"+strconv.Itoa(userDate.UserNum))
	err,date:=commonFun.HttpUrlFunc("POST","/storeCat/create",assStoreCat)
	if err!=nil{
		log.Println(err)
		return  false
	}
	result := commonFun.Expected(date.Code, 0)
	if result == false {
		return false
	}
	return true
}

func GetStoreCatId()(error,int){
	conn,_=datePrepare.NewTestConn()
	var storeCatId int
	getStoreIdSql:="select tid from rcp_shopcategory  where name=?"
	rows,err:=conn.Query(getStoreIdSql,"商品测试类型1")
	if err!=nil{
		log.Fatal("getStoreCatId  query  error", err.Error())
		return err,-1
	}else {
		for rows.Next(){
			err=rows.Scan(&storeCatId)
			if err!=nil{
				log.Fatal("getStoreCatId Scan err",err.Error())
				return err,-1
			}

		}
	}
	return nil,storeCatId
}
