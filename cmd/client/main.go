package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)
const(
	Url = "https://localhost:9090/check"
	)
var param = []string{"sas","swq","rtet","12","89","77"}
func main() {
	log.Println("Start http client...")
	cli := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}}
	data,_ := json.Marshal(param)
	resp, err := cli.Post(Url,"application/json;charset=UTF-8", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("code: %d request : %s error:%s",resp.StatusCode,Url,err.Error())
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout,resp.Body)

}
