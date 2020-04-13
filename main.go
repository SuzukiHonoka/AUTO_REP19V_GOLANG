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
	get string = "get"
	post string = "post"
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

func getResp(client http.Client,req *http.Request) (*http.Response,error) {
	resp,err := client.Do(req)
	if err != nil {
		return nil,err
	}
	updateCookies(resp)
	defer resp.Body.Close()
	return resp,nil

}

func urlRequest(mode int,durl string) *http.Response {
	var client http.Client
	switch mode {
	// mode 0=get 1=post data=None
	case 0:
		fmt.Println("Get:",durl)
		req := newRequest(get,durl,nil)
		loadCookies(req)
		resp,err := getResp(client,req)
		if err != nil {
			pe(err)
		} else {
			return resp
		}
	case 1:
		fP,err := json.Marshal(postData)
		if err != nil {
			pe(err)
		} else {
			req := newRequest(post,durl,bytes.NewBuffer(fP))
			loadCookies(req)
			resp,err := getResp(client,req)
			if err != nil {
				return resp
			} else {
				pe(err)
			}
		}
	}
	return nil
}

func main()  {
	welcome()
	getInput()
	applies()
	fr,_ := ioutil.ReadAll(urlRequest(0,yoya_login_g).Body)
	fmt.Println(string(fr))
}
