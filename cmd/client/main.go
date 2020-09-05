package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	CheckUrl      = "https://localhost:9090/check"
	GetAllDataUrl = "https://localhost:9090/data"
)

var action int
var data string

func init()  {
	flag.IntVar(&action,"action",1,"1 get all test data,2 check test data")
	flag.StringVar(&data,"data","","commit test data and delimit by ,")
	flag.Parse()
}
/**
 *post test data to server
 */
func checkTestData(list []string) {
	cli := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}}
	data, _ := json.Marshal(list)
	resp, err := cli.Post(CheckUrl, "application/json;charset=UTF-8", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("code: %d request : %s error:%s", resp.StatusCode, CheckUrl, err.Error())
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

}

/**
 *get all data
 */
func getAllData() {
	cli := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}}
	resp, err := cli.Get(GetAllDataUrl)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("code: %d request : %s error:%s", resp.StatusCode, CheckUrl, err.Error())
		return
	}
	io.Copy(os.Stdout, resp.Body)

}

func main() {
	log.Println("action:",action,"data:",data)
	if action == 1 {
		getAllData()
		return
	}

	list := strings.Split(strings.Trim(data, " "), ",")
	if len(list) < 1 || data == "" {
		log.Println("check data at least one")
		return
	}
	checkTestData(list)

}
