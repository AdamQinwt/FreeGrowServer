package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	tableParamsName = " params "

	//fields
	fid      = "id"
	fname    = "name"
	ftime    = "time"
	fcurrent = "current"
	fmax     = "max"
	fmin     = "min"
)

type Params struct {
	Id      int
	Name    string `json:"Name"`
	Time    string `json:"Time"`
	Current int    `json:"Current"`
	Max     int    `json:"Max"`
	Min     int    `json:"Min"`
}

func DropTable() error {
	q := "Drop TABLE " + tableParamsName + "  ;"
	//fmt.Println(q)
	return Execute(q)
}

func CreateTable() error {
	q := "CREATE TABLE if not exists " + tableParamsName + " ("
	q += " id INTEGER primary key auto_increment, "
	q += fname + " text, "
	q += ftime + " text, "
	q += fcurrent + " integer,"
	q += fmax + " integer,"
	q += fmin + " integer"
	q += ");"
	//fmt.Println(q)
	return Execute(q)
}

func RTOAParams(r Record) []Params {
	result := make([]Params, 0, 5)
	for _, v := range r {
		id, _ := strconv.ParseInt(v[fid], 10, 32)
		current, _ := strconv.ParseInt(v[fcurrent], 10, 32)
		max, _ := strconv.ParseInt(v[fmax], 10, 32)
		min, _ := strconv.ParseInt(v[fmin], 10, 32)
		result = append(result, Params{
			Id:      int(id),
			Name:    v[fname],
			Time:    v[ftime],
			Current: int(current),
			Max:     int(max),
			Min:     int(min),
		})
	}
	return result
}

func FindAll() []Params {
	q := "select * from" + tableParamsName
	//fmt.Println(q)
	r := GetRecord(q)
	//fmt.Println(r)
	c := RTOAParams(r)

	if len(c) != 0 {
		return c
	} else {
		return nil
	}
}

func ParamsToJSON(u []Params) []byte {
	jsonFile, err := os.Create("post.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return nil
	}
	defer jsonFile.Close()
	jsonWriter := io.Writer(jsonFile)
	encoder := json.NewEncoder(jsonWriter)
	err = encoder.Encode(&u)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return nil
	}
	return nil
}

func (params *Params) ToJSON() []byte {
	output, err := json.Marshal(&params)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return output
}

func (params *Params) InsParams() error {
	q := "Insert into " + tableParamsName
	q += " ( "
	q += fname + " , "
	q += ftime + " , "
	q += fcurrent + " , "
	q += fmax + " , "
	q += fmin
	q += " ) values("
	q += "'" + params.Name + "'" + ","
	q += "'" + params.Time + "'" + ","
	q += strconv.Itoa(params.Current) + ","
	q += strconv.Itoa(params.Max) + ","
	q += strconv.Itoa(params.Min)
	q += ")"
	//fmt.Println(q)
	err := Execute(q)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func DelParams(id int) error {
	err := Execute("delete from " + tableParamsName + " where Id = " + strconv.Itoa(id))
	return err
}

//Parameters decide whether
func (params *Params) UpdateParams(current bool, time bool, max bool, min bool) error {
	q := "update " + tableParamsName
	q += " set "
	first := true
	if current {
		q += fcurrent + " = " + strconv.Itoa(params.Current)
		first = false
	}
	if time {
		if !first {
			q += " , "
		}
		q += ftime + " = '" + params.Time + "'"
		first = false
	}
	if max {
		if !first {
			q += " , "
		}
		q += fmax + " = " + strconv.Itoa(params.Max)
		first = false
	}
	if min {
		if !first {
			q += " , "
		}
		q += fmin + " = " + strconv.Itoa(params.Min)
	}

	q += " where "
	q += fname + " = '" + params.Name + "'"
	//fmt.Println(q)
	err := Execute(q)
	return err
}
func GetParamsByName(name string) Params {
	q := "select * from" + tableParamsName
	q += " where " + fname + " = '" + name + "'"
	//fmt.Println(q)
	r := GetRecord(q)
	//fmt.Println(r)
	c := RTOAParams(r)

	if len(c) != 0 {
		return c[0]
	} else {
		return Params{
			Id:   0,
			Name: name,
		}
	}
}
