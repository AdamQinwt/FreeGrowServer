package main

const (

	//additional defs
	VoidString = "&"
	VoidInt    = -1
)

type User struct {
	Id   int
	Name string `json:"Name"`
	PWD  string `json:"PWD"`
}

func VerifyUser(u User) bool {
	if u.Name != "root" {
		return false
	}
	if u.PWD != "root" {
		return false
	}
	return true
}
