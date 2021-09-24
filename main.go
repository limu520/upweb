package main

import (
	//"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func geturl(url string) string {
	req, _ := http.NewRequest("POST", url, nil)
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)

}

func getservice(id string) string {
	aurl := geturl("http://172.30.0.11")
	// fmt.Println(aurl)
	upurl := aurl[68:498]

	resp, _ := http.PostForm("http://172.30.0.11:80/eportal/userV2.do?method=getServices", url.Values{"username": {id}, "search": {upurl}})

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	sname := ""
	if strings.Contains(string(body), "移动pppoe") {
		sname = url.QueryEscape("移动pppoe")
	} else if strings.Contains(string(body), "联通专线") {
		sname = string(url.QueryEscape("联通专线"))
	} else if strings.Contains(string(body), "电信专线") {
		sname = string(url.QueryEscape("电信专线"))
	} else {
		fmt.Println("您是否未购买校园网套餐，或者本程序存在BUG")
	}
	//fmt.Println(sname)
	return string(sname)
}

func upweb(id string, passwd string, service int) {
	aurl := geturl("http://172.30.0.11")
	// fmt.Println(aurl)
	upurl := aurl[69:498]
	rurl := url.QueryEscape(upurl)
	// fmt.Println(rurl)
	if service == 1 {
		name := getservice(id)
		// fmt.Println(name)
		resp, _ := http.PostForm("http://172.30.0.11:80/eportal/InterFace.do?method=login", url.Values{"userId": {id}, "password": {passwd}, "service": {name}, "queryString": {rurl}, "operatorPwd": {""}, "operatorUserId": {""}, "validcode": {""}, "passwordEncrypt": {"false"}})
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	} else if service == 2 {
		sname := string(url.QueryEscape("校园网"))
		resp, _ := http.PostForm("http://172.30.0.11:80/eportal/InterFace.do?method=login", url.Values{"userId": {id}, "password": {passwd}, "service": {sname}, "queryString": {rurl}, "operatorPwd": {""}, "operatorUserId": {""}, "validcode": {""}, "passwordEncrypt": {"false"}})
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	} else {
		fmt.Println("参数错误")
	}

}

func main() {
	var id string
	var passwd string
	var service int
	flag.StringVar(&id, "i", " ", "用户名（学号）")
	flag.StringVar(&passwd, "p", " ", "用户密码")
	flag.IntVar(&service, "s", 2, "用户服务，自动选择购买的服务为1,接校园网内网为2")
	flag.Parse()

	upweb(id, passwd, service)
}
