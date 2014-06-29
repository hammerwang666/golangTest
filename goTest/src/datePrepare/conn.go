package datePrepare

import (
	"database/sql"
	"config"
//	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/Centny/Cny4go/log"
)
var  connConfig  string =""
var  db *sql.DB
func SetConnConfig(s string)(){
	connConfig=s
}
func GetNewConn()(*sql.DB,error){
    if  db!=nil{
//      log.I("Close the old db  open the  new one :%v \n",connConfig)
		db.Close()
   }
	var err  error
	db ,err=sql.Open("mysql",connConfig)
	if err !=nil{
		log.E("GetTestDB db error:", err.Error())
		return  nil,err
	}
	return  db ,nil
}
func NewTestConn()(*sql.DB,error){
    var  conn *sql.DB
	dbConfig:=config.GetConfig("DB_CONFIG")
//	fmt.Println("test db is :")
	SetConnConfig(dbConfig.(string))
	conn,_=GetNewConn()
	return conn, nil
}
