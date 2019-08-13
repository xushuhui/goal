/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package tests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	// ...
}
func TestRegister(t *testing.T) {
	// ...
}
func mockServer() *httptest.Server {
	//API调用处理函数
	sendJson := func(rw http.ResponseWriter, r *http.Request) {
		u := struct {
			Name string
		}{
			Name: "张三",
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(rw).Encode(u)
	}
	//适配器转换
	return httptest.NewServer(http.HandlerFunc(sendJson))
}

func TestSendJSON(t *testing.T) {
	//创建一个模拟的服务器
	server := mockServer()
	defer server.Close()
	//Get请求发往模拟服务器的地址
	resq, err := http.Get(server.URL)
	if err != nil {
		t.Fatal("创建Get失败")
	}
	defer resq.Body.Close()

	log.Println("code:", resq.StatusCode)
	jsons, err := ioutil.ReadAll(resq.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("body:%s\n", jsons)
}
