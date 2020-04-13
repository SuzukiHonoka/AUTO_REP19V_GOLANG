package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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
	user_im string = "AUTO_REP19V"
	user_bl string
	user_id string
	user_pa string
	//report
	rtype int = 1
	postData = map[string]string{"type":string(rtype)}
	//cookies
	cookie []*http.Cookie
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
			user_bl = input
		case 1:
			fmt.Println(input_id)
			input,_ := reader.ReadString('\n')
			user_id = input
		case 2:
			fmt.Println(input_pa)
			input,_ := reader.ReadString('\n')
			user_pa = input
		}
	}
}

func applies(){
	fmt.Println("Applying..")
	yoya_login_g = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(yoya_login_g,ubl,user_bl),uid,user_id),upa,user_pa),uim,user_im),"\n","")
	yoya_base_g = strings.ReplaceAll(strings.ReplaceAll(yoya_base_g,ubl,user_bl),"\n","")

}

func newRequest(method string,durl string,body io.Reader) *http.Request {
	req,_ := http.NewRequest(method,durl,body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36")
	req.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Connection","keep-alive")
	req.Header.Set("Accept-Language" ,"zh-CN,zh;q=0.9,en;q=0.8")
	return req
}

func loadCookies(req *http.Request){
	if cookie != nil {
		for i := range cookie {
			req.AddCookie(cookie[i])
		}
	}
}

func updateCookies(resp *http.Response)  {
	cookie = resp.Cookies()
}

func getResp(client http.Client,req *http.Request) (string,int,error) {
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

func urlRequest(mode int,durl string) (string,int,error) {
	var client http.Client
	switch mode {
	// mode 0=get 1=post data=None
	case 0:
		fmt.Println("Get:",durl)
		req := newRequest(get,durl,nil)
		loadCookies(req)
		resp,code,err := getResp(client,req)
		if err != nil {
			pe(err)
		} else {
			return resp,code,nil
		}
	case 1:
		fP,err := json.Marshal(postData)
		if err != nil {
			pe(err)
		} else {
			req := newRequest(post,durl,bytes.NewBuffer(fP))
			loadCookies(req)
			resp,code,err := getResp(client,req)
			if err != nil {
				pe(err)
			} else {
				return resp,code,nil
			}
		}
	}
	return "",0,nil
}

func main()  {
	welcome()
	getInput()
	applies()
	res,_,err := urlRequest(0,yoya_login_g)
	if err != nil {
		pe(err)
	}
	fmt.Println(res)
}
