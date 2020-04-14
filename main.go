package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
)

const (
	//hello
	hello string = "Welcome to use this AUTO_REP19V.\nIt is based in GOLANG and written by starx.\nYou'll agree the term of use by you start this program."
	input_bl string = "Please Enter the User Block."
	input_id string = "Please Enter the User ID."
	input_pa string = "Please Enter the User Password."
	//
	get string = "GET"
	post string = "POST"
	//
	ubl string = "user_block"
	uid string = "user_id"
	upa string = "user_pa"
	uim string = "user_imei"
	uuid string = "uu_id"
	//
)

var (
	//api
	yoya_login_g string = "https://cz.yoya.com/user_block/do?action=cz/h5/login&start=login&start=login&user_code=user_id&app_type=Android&action=cz/h5/login&user_pwd=user_pa&client_code=user_imei"
	yoya_base_g string = "https://jkjc.yoya.com/user_block/do?action=tzjc/h5/gzbdreport&start=getPreviousReport&user_id=uu_id"
	yoya_report_p string = "https://jkjc.yoya.com/user_block/do?action=tzjc/h5/gzbdreport&start=saveReport"
	//login
	userIm string = "AUTO_REP19V"
	userBl string
	userId string
	userPa string
	userUid string
	//info
	userRealId   string = "000000000000000000"
	userRealAddr string = "Paradise"
	userTel      string = "911"
	userBackTime string = "GUESS?"
	//report
	postData = map[string]string{"type":"1","user_id":"","id_type":"1","identity_code":"real_id","address":"user_real_addr","telephone":"user_tel","back_time":"user_back_time","go_where":"None","contact_type":"1","es":"1","health_status":"2","is_diagnosis":"","is_fever":"0","temperature":"36","is_cough":"0","isolate":"0","isolate_type":"","isolate_time":"","remark":"AUTO_REP_19V_GOLANG"}
	//cookies
	cookies []*http.Cookie
	cookieJar, _ = cookiejar.New(nil)
	//
	req1Json string
	data_g string
	data_u string
	base_d string
)


func pe(e error)  {
	fmt.Println(e)
	os.Exit(1)
}

func welcome()  {
	fmt.Println(hello)
}

func getInput()  {
	for i := 0;i < 3;i++ {
		reader := bufio.NewReader(os.Stdin)
		switch i {
		case 0:
			fmt.Println(input_bl)
			input,_ := reader.ReadString('\n')
			userBl = input
		case 1:
			fmt.Println(input_id)
			input,_ := reader.ReadString('\n')
			userId = input
		case 2:
			fmt.Println(input_pa)
			input,_ := reader.ReadString('\n')
			userPa = input
		}
	}
}

func applies(){
	fmt.Println("Applying..")
	yoya_login_g = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(yoya_login_g,ubl, userBl),uid, userId),upa, userPa),uim, userIm),"\n","")
	yoya_base_g = strings.ReplaceAll(strings.ReplaceAll(yoya_base_g,ubl, userBl),"\n","")
	yoya_report_p = strings.ReplaceAll(strings.ReplaceAll(yoya_report_p,ubl,userBl),"\n","")

}

func newRequest(method string,durl string,body io.Reader) *http.Request {
	req,err := http.NewRequest(method,durl,body)
	if err != nil {
		pe(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36")
	req.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Connection","keep-alive")
	req.Header.Add("Accept-Language" ,"zh-CN,zh;q=0.9,en;q=0.8")
	return req
}

func loadCookies(req *http.Request){
	if cookies != nil {
		fmt.Println("Cookies:",cookies)
		for i := range cookies {
			req.AddCookie(cookies[i])
		}
	}
}

func updateCookies(resp *http.Response)  {
	cookies = resp.Cookies()
}

func getResp(client *http.Client,req *http.Request) (string,int,error) {
	resp,err := client.Do(req)
	if err != nil {
		return "",0,err
	}
	updateCookies(resp)
	defer resp.Body.Close()
	resBytes,err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		pe(err2)
	}
	return string(resBytes),resp.StatusCode,nil

}

func urlRequest(mode int,durl string,data io.Reader) (string,int,error) {
	client := &http.Client {
		Jar: cookieJar,
	}
	switch mode {
	// mode 0=get 1=post data=None
	case 0:
		fmt.Println("Get:",durl)
		req := newRequest(get,durl,nil)
		//loadCookies(req)
		resp,code,err := getResp(client,req)
		if err != nil {
			pe(err)
		} else {
			return resp,code,nil
		}
	case 1:
		fmt.Println("Post:",durl)
		req := newRequest(post,durl,data)
		//loadCookies(req)
		resp,code,err := getResp(client,req)
		if err != nil {
			pe(err)
		} else {
				return resp,code,nil
			}

	}
	return "",0,nil
}

func tryLogin()  {
	res,code,err := urlRequest(0,yoya_login_g,nil)
	if err != nil || code !=200 || gjson.Get(res,"code").Int() != 200 {
		fmt.Println(res)
		pe(err)
	}
	req1Json = res
	fmt.Println("Login Success!")
}

func updateInfo()  {
	data_g = gjson.Get(req1Json,"data").String()
	data_u = gjson.Get(data_g,"user").String()
	userUid = gjson.Get(data_u,"user_id").String()
}

func secHello()  {
	fmt.Println("Last_Login_Time:",gjson.Get(data_u,"last_login_time"),"\nName:",gjson.Get(req1Json,"data.user.user_name"),"\nBlock:",gjson.Get(data_g,"siteInfo.site_code"))
}

func prePare()  {
	yoya_base_g = strings.ReplaceAll(yoya_base_g,uuid,userUid)
	res,code,err := urlRequest(0,yoya_base_g,nil)
	if err != nil || code != 200 {
		pe(err)
	}
	base_d = gjson.Get(res,"data").String()
	userRealId = gjson.Get(base_d,"identity_code").String()
	userRealAddr = gjson.Get(base_d,"address").String()
	userTel = gjson.Get(base_d,"telephone").String()
	userBackTime = gjson.Get(base_d,"back_time").String()
	postData["user_id"] = userUid
	postData["identity_code"] = userRealId
	postData["address"] = userRealAddr
	postData["telephone"] = userTel
	postData["back_time"] = userBackTime
}

func postRep()  {
	// ordered by sequence

	//keys := make([]string,len(postData))
	//
	//for k,_ := range postData {
	//	keys = append(keys,k)
	//}
	//sort.Strings(keys)
	fmt.Println(postData)

	fp,er := json.Marshal(postData)
	if er != nil {
		pe(er)
	}
	res,code,err := urlRequest(1,yoya_report_p,bytes.NewReader(fp) )
	if err != nil || code != 200 {
		pe(err)
	}
	fmt.Println(res)
}

func main()  {
	welcome()
	getInput()
	applies()
	tryLogin()
	updateInfo()
	secHello()
	prePare()
	postRep()
}
