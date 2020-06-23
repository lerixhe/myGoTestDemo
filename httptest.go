package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// rsp, _ := http.Get("http://127.0.0.1:2343")
	// rsp, _ := http.Get("http://190.216.58.51:8043/test/ips/heathcheck")
	// rsp, _ := http.Get("http://baidu.com")
	tr := &http.Transport{
		// 开启http长连接与压缩
		DisableKeepAlives:  false,
		DisableCompression: false,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   180 * time.Second,
	}
	req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1:2343", nil)
	rsp, _ := client.Do(req)
	ioutil.ReadAll(rsp.Body)
	fmt.Println("是否被压缩？", rsp.Uncompressed)
	// fmt.Println(string(body))
}
