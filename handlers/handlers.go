package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/todo_application/authentication"
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
	if c.Request.Method != "POST" {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	uname := c.PostForm("username")
	password := c.PostForm("password")

	var emp model.Employee
	filter := bson.M{
		"$and": []bson.M{
			{"username": uname},
			{"password": password},
		},
	}
	err := config.Col.Find(filter).One(&emp)
	if err != nil {
		c.String(http.StatusUnauthorized, "Employee Not Found")
		return
	}

	empType := emp.EmployeeType
	if empType == "" {
		c.String(http.StatusUnauthorized, "Invalid Credentials")
	}

	accToken, _, err := authentication.GenerateTokenAndRefreshToen(emp)
	if err != nil {
		c.String(http.StatusBadRequest, "Error in Generating tokens")
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    accToken,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
}

func Log(c *gin.Context) {
	config.Tpl.ExecuteTemplate(c.Writer, "login.html", nil)
}
