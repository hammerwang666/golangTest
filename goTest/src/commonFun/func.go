package commonFun

import (
	"strings"
	"net/url"
	"net/http"
	"log"
	"os"
	"io/ioutil"
    "encoding/json"
	"result"
	"commonParam"
)

type ResponseStruct struct {
	Code        int    `json:"code"`
	Message     string    `json:"msg"`
	Data        interface{} `json:"data"`
	Description string  `json:"description"`
}

var serviceUrl =commonParam.Url
type Jar struct{
	cookies[]*http.Cookie
}

func (jar *Jar) SetCookies(u *url.URL, cookies[]*http.Cookie) {
	jar.cookies = cookies
}
func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
	return jar.cookies
}

var jar = new(Jar)
var client = http.Client{nil, nil, jar}
func HttpUrlFunc(method string, url  string , values url.Values, structType interface {}) (error,int){
	url = serviceUrl+url
	b := strings.NewReader(values.Encode())
	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		log.Fatal("http fail")
		return err,-1
	}
	if "POST" == method {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Fatal error ", err.Error())
		os.Exit(0)
		return err,-1
	}
//	cookies := res.Cookies()
//	for _, cookie := range cookies {
//		fmt.Println("cookie:", cookie)
//	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	resp_body:=string(body)
	if 0==len(resp_body){
		log.Fatalln("Get The Worng Response:EmPTY")
	}
//	var response *ResponseStruct = &ResponseStruct{}

	err= json.Unmarshal([]byte(resp_body), &structType)
   return nil,0
}

//expect for   case
func Expected(actaual interface {},expected interface {}) bool{
	result.CaseNum++
	result.PreResult.CaseResult.Number=result.CaseNum
	result.PreTotalReport.PreTestTotalCase++
     if actaual!=expected{
		 result.PreTotalReport.PreTestFalseCaseNum++
		 result.PreResult.CaseResult.Result="false"
		 result.TotalResult=append(result.TotalResult,result.PreResult)
		 result.TotalReported = append(result.TotalReported, result.PreTotalReport)//find the  false
		 return false
	 }
    result.PreTotalReport.PreTestTrueCaseNum++
	result.PreResult.CaseResult.Result="true"
	result.TotalResult=append(result.TotalResult,result.PreResult)
	return true
}



