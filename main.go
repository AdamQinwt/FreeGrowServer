package main

import (
	"net/http"
)

func init() {
	CreateTable()
	p := Params{
		Name:    "water",
		Time:    "20180718-14:10",
		Current: 50,
		Min:     0,
		Max:     100,
	}
	p.InsParams()
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8088",
	}
	http.HandleFunc("/water", Water)
	http.HandleFunc("/data", Data)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/", Login)
	http.HandleFunc("/login_usr", User_Data)
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./html"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	server.ListenAndServe()
}
