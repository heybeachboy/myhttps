package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)
const(
	ServerCrt = "../tls/server.crt"
	ServerCsr = "../tls/service.csr"
	ServerKey = "../tls/server.key"
)
func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("method:",r.Method)
	io.WriteString(w,"Hello world!!!\n")

}

type U struct {
	 Name string `json:"name"`
	 Age  int `json:"age"`
	 Tel  int `json:"tel"`
}
func User(w http.ResponseWriter, r *http.Request) {
	 u := new(U)
	 u.Name = "LeLe"
	 u.Age = 20
	 u.Tel = 15086846070
	 data, _ := json.Marshal(u)
     w.Write(data)
     w.Header().Set("ContentType","application/json")
     w.WriteHeader(http.StatusOK)
}

func main() {
	 log.Println("Start https service...")
	 http.HandleFunc("/",Home)
	 http.HandleFunc("/user",User)
	 if err := http.ListenAndServeTLS(":9090",ServerCrt,ServerKey,nil);err != nil {
	 	log.Fatal("Start https server error:",err.Error())
	 }

}
