package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"os"
)
const(
	Url = "https://localhost:9090/"
)
func main() {
	log.Println("Start http client...")
	cli := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}}
	resp, err := cli.Get(Url)
	if err != nil {
		log.Fatalf("request : %s error:%s",Url,err.Error())
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout,resp.Body)

}
