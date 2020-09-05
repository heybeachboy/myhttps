package main

import (
	"encoding/json"
	"io/ioutil"
	"myhttps/core"
	"myhttps/memory"

	"io"
	"log"
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

type U struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Tel  int    `json:"tel"`
}

func User(w http.ResponseWriter, r *http.Request) {
	u := new(U)
	u.Name = "LeLe"
	u.Age = 20
	u.Tel = 15086846070
	data, _ := json.Marshal(u)
	w.Write(data)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func CheckTestData(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
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
	http.HandleFunc("/user", User)
	http.HandleFunc("/check", CheckTestData)
	if err := http.ListenAndServeTLS(":9090", ServerCrt, ServerKey, nil); err != nil {
		log.Fatal("Start https server error:", err.Error())
	}
	defer memory.MemStorage.Flush() //清除内存数据

}
