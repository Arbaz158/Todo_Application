package config

import (
	"fmt"
	"html/template"

	"gopkg.in/mgo.v2"
)

var Tpl *template.Template
var DB *mgo.Database

func init() {
	Tpl = template.Must(template.ParseGlob("templates/*.html"))
	session, err := mgo.Dial("mongodb://172.31.43.162:27017")
	if err != nil {
		fmt.Println("error in connection :", err)
		return
	}
	session.SetMode(mgo.Monotonic, true)
	DB = session.DB("testing")
}
