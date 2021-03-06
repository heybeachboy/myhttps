package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"myhttps/core"
	"myhttps/memory"
	"net/http"
)

const (
	ServerCrt = "../tls/server.crt"
	ServerCsr = "../tls/service.csr"
	ServerKey = "../tls/server.key"
)

func Home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!!!\n")
}

func GetAllData(w http.ResponseWriter, r *http.Request) {
     list := memory.MemStorage.GetAllData()
     by, _ := json.Marshal(&list)
     w.Write(by)
     w.Header().Set("Content-Type","application/json")
     w.WriteHeader(http.StatusOK)
}
func CheckTestData(w http.ResponseWriter, r *http.Request) {
	//log.Println(r.Method)
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var input []string

	by, err2 := ioutil.ReadAll(r.Body)
	if err2 != nil {
		w.Write([]byte(err2.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err3 := json.Unmarshal(by, &input)
	if err3 != nil {
		w.Write([]byte(err3.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, _ := json.Marshal(core.CheckData(input))
	w.Write(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func main() {
	log.Println("Start https service...")
	http.HandleFunc("/", Home)
	http.HandleFunc("/data",GetAllData)
	http.HandleFunc("/check", CheckTestData)
	if err := http.ListenAndServeTLS(":9090", ServerCrt, ServerKey, nil); err != nil {
		log.Fatal("Start https server error:", err.Error())
	}
	defer memory.MemStorage.Flush() //清除内存数据

}
