package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func Water(w http.ResponseWriter, r *http.Request) {
	//verify user
	//fmt.Println(r.Header["Cookie"])
	t, err := template.ParseFiles("html/water.html")
	if err != nil {
		fmt.Fprintf(w, "parse template error: %s", err.Error())
		return
	}
	w.Header().Set("Content-type", "text/html")
	t.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/login.html")
	if err != nil {
		fmt.Fprintf(w, "parse template error: %s", err.Error())
		return
	}
	w.Header().Set("Content-type", "text/html")
	t.Execute(w, nil)
}

func User_Data(w http.ResponseWriter, r *http.Request) {
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	//fmt.Println(string(body))
	var t User
	json.Unmarshal(body, &t)
	ok := VerifyUser(t)
	if ok {
		c := http.Cookie{
			Name:  "user",
			Value: t.Name,
		}
		//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Add("Set-Cookie", c.String())
		//w.WriteHeader(http.StatusOK)
		//http.SetCookie(w, &c)
		fmt.Fprintf(w, "Success")
		//fmt.Println("Success")
	} else {
		//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		//w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "__Error")
	}

}

func Data(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("user")
	if err != nil {
		fmt.Fprint(w, "__Login")
		return
	}
	name := r.URL.Query()["name"][0]
	p := GetParamsByName(name)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		panic(err)
	}
}
