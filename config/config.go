package config

import (
	"fmt"
	"html/template"

	"gopkg.in/mgo.v2"
)

var Tpl *template.Template
var Col *mgo.Collection

func init() {
	Tpl = template.Must(template.ParseGlob("templates/*.html"))
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		fmt.Println("error in connection :", err)
		return
	}
	session.SetMode(mgo.Monotonic, true)
	Col = session.DB("testing").C("employee")
}
