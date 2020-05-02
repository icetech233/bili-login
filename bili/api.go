package bili

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const get_login_url = "https://passport.bilibili.com/qrcode/getLoginUrl"
const get_login_info = "https://passport.bilibili.com/qrcode/getLoginInfo"

// Get Login Url 获取二维码的 地址
func GetLoginUrl() LoginUrlModel {
	resp, err := http.Get(get_login_url)
	if err != nil {
		fmt.Println("get Login Url :", err)
	}
	bodyRep, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	//fmt.Println(string(bodyRep))
	var baseApi BiliApiBaseModel
	err = json.Unmarshal(bodyRep, &baseApi)
	if err != nil {
		fmt.Println("Unmarshal failed :", err)
	}
	//log.Println(baseApi.Data.OauthKey)
	return baseApi.Data
}

// Get Login Info 官方网站是三秒一个循环
func GetLoginInfo(oauthKey string) ([]byte, http.Header) {
	data := make(url.Values)
	data["oauthKey"] = []string{oauthKey}
	data["gourl"] = []string{"https://www.bilibili.com/"}
	resp, err := http.PostForm(get_login_info, data)
	if err != nil {
		fmt.Println("getLoginInfo", err)
	}
	headerRep := resp.Header.Clone()
	bodyRep, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return bodyRep, headerRep
}
