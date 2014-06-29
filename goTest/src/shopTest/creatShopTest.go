package shopTest

import (
	"net/url"
	"strconv"
	"commonFun"
	"log"
	"userDate"
	"result"
)

func CreatShopTest() bool{


	//增加测试结果,一个测试写一个
	result.TestNum++
	result.PreResult.Test.Name="CreatShopTest"
	result.PreResult.Test.Number=result.TestNum

	shopMsg := url.Values{}
	shopMsg.Set("storeName", "老罗英文教育书店"+strconv.Itoa(userDate.UserNum))
	shopMsg.Set("shopCategoryId", "1")
	shopMsg.Set("storeIntro", "test store")
	err, date2 := commonFun.HttpUrlFunc("POST", "/register-store", shopMsg)

	if err != nil {
		log.Println(err)
		return  false
	}
	result := commonFun.Expected(date2.Code, 0)
	if result == false {
		return  false
	}


	shopMsg1 := url.Values{}
	shopMsg1.Set("storeName", "")
	shopMsg1.Set("shopCategoryId", "")
	shopMsg1.Set("storeIntro", "test store")
	err, date2 = commonFun.HttpUrlFunc("POST", "/register-store", shopMsg1)

	if err != nil {
		log.Println(err)
		return  false
	}
//	result = commonFun.Expected(date2.Message, " ")
	result = commonFun.Expected(date2.Message, "商店名为空")
	if result == false {
		return  false
	}

	return true
}

