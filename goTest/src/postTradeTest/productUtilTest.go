package postTradeTest

import(
	"result"
	"net/url"
	"strconv"
	"userDate"
	"categoryTest"
	"log"
	"commonFun"
)

type Resp struct {
	ErrorTip  string
	ErrorCode  int
	Data      interface{}
}
func ProductReleaseTest()bool{

	result.TesDetailResult("ProductReleaseTest")
	//case 1
	err, categoryId:=categoryTest.GetCategoryId()
	if err!=nil{
		log.Println("get categoryId fail ",err)
		return false
	}
	values:=url.Values{
		"sourceName":{"测试商品"+strconv.Itoa(userDate.UserNum)},
		"sourcePrice":{"9999.99"},
        "sourceNum":{"10"},
		"sourceIntro":{"sourceIntro"},
		"sourceCategory":{strconv.Itoa(categoryId)},
		"sourcePic":{"www.pic.com/bmp.jpg"},
		"sourceLogistics":{"123"},
		"sourceNo":{"123"},
		"sourceType":{"123"},
		"sourceDetailType":{"0"},
		"sourceDetailList":{""},
         "sourceWeight":{"28.88"},
		"sourceLocation":{"北京|东城"},
		"sourceShowcase":{"0"},

	}
	resParse:=&Resp{}
	err,_=commonFun.HttpUrlFunc("POST","/productRelease",values,resParse)
	if err!=nil{
		log.Println(err)
		return  false
	}
	result := commonFun.Expected(resParse.ErrorCode, 0)
	if result == false {
		return false
	}
	return  true
}
