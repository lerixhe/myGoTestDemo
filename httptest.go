package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	rsp, _ := http.Get("http://127.0.0.1:8042/test/info/station?pageSize=3&pageNo=1")
	body, _ := ioutil.ReadAll(rsp.Body)
	fmt.Println("是否被压缩？", rsp.Uncompressed)
	fmt.Println(string(body))

}
