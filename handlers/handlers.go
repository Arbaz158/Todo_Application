package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/todo_application/config"
	"github.com/todo_application/model"
	"gopkg.in/mgo.v2/bson"
)

func Home(c *gin.Context) {
	config.Tpl.ExecuteTemplate(c.Writer, "home.html", nil)
}

func Sign(c *gin.Context) {
	config.Tpl.ExecuteTemplate(c.Writer, "signup.html", nil)
}

func SignUp(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	uname := c.PostForm("username")
	password := c.PostForm("password")
	etype := c.PostForm("employeetype")
	email := c.PostForm("email")

	var emp model.Employee
	emp.UserName = uname
	emp.Password = password
	emp.Email = email
	emp.EmployeeType = etype

	filter := bson.M{"$or": []bson.M{
		{"email": email},
		{"username": uname},
	}}
	count, _ := config.Col.Find(filter).Count()
	if count > 0 {
		c.String(http.StatusBadRequest, "Username or Email already exists")
		return
	}

	err := config.Col.Insert(&emp)
	if err != nil {
		c.String(http.StatusBadRequest, "error in inserting employee data")
		return
	}
}

func Login(c *gin.Context) {

}

func Log(c *gin.Context) {

}
