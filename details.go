package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func NewParams(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("user")
	if err != nil {
		return
	}
	t, err := template.ParseFiles("html/newParams.html")
	if err != nil {
		fmt.Fprintf(w, "parse template error: %s", err.Error())
		return
	}
	w.Header().Set("Content-type", "text/html")
	t.Execute(w, nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("user")
	if err != nil {
		return
	}
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	//fmt.Println(string(body))
	var t Params
	json.Unmarshal(body, &t)
	//fmt.Println(t)
	t.InsParams()
}
func Update(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("user")
	if err != nil {
		fmt.Fprintf(w,"__Login")
		return
	}
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	//fmt.Println(string(body))
	var t Params
	json.Unmarshal(body, &t)
	err = t.UpdateParams(false,false, true, true)
	if err == nil {
		fmt.Fprintf(w, "Success")
	} else {
		fmt.Fprintf(w, "Failure!"+err.Error())
	}
}
func GetParams(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("user")
	if err != nil {
		fmt.Fprintf(w,"__Login")
		return
	}
	name := r.URL.Query().Get("name")
	b := GetParamsByName(name)
	t, err := template.ParseFiles("html/info.html")
	if err != nil {
		fmt.Printf("parse template error: %s", err.Error())
		return
	}
	w.Header().Set("Content-type", "text/html")
	t.Execute(w, b)
}
