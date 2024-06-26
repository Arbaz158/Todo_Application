package handlers

import (
	"net/http"
	"strconv"
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
	count, _ := config.DB.C("employee").Find(filter).Count()
	if count > 0 {
		c.String(http.StatusBadRequest, "Username or Email already exists")
		return
	}

	err := config.DB.C("employee").Insert(&emp)
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
	err := config.DB.C("employee").Find(filter).One(&emp)
	if err != nil {
		c.String(http.StatusUnauthorized, "Employee Not Found")
		return
	}

	empType := emp.EmployeeType
	if empType == "" {
		c.String(http.StatusUnauthorized, "Invalid Credentials")
	}

	accToken, _, err := authentication.GenerateTokenAndRefreshToken(emp)
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

func Add(c *gin.Context) {
	config.Tpl.ExecuteTemplate(c.Writer, "addstuff.html", nil)
}

func AddStuff(c *gin.Context) {

	if c.Request.Method != "POST" {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	cookie, err := c.Cookie("token")
	if err != nil {
		c.String(http.StatusBadRequest, "error in getting cookie")
		return
	}
	token := cookie
	userName, err := authentication.VerifyToken(token)
	if err != nil {
		c.String(http.StatusBadRequest, "error in verifying token")
		return
	}
	var emp model.Employee
	filter := bson.M{"username": userName}
	err = config.DB.C("employee").Find(filter).One(&emp)
	if err != nil {
		c.String(http.StatusBadRequest, "error in getting employee type")
		return
	}
	if emp.EmployeeType == "manager" {
		var st model.Stuff
		desktop := c.PostForm("desktop")
		monitor := c.PostForm("monitor")
		cpu := c.PostForm("monitor")
		quantity := c.PostForm("quantity")
		st.Desktop = desktop
		st.Monitor = monitor
		st.CPU = cpu
		quan, err := strconv.Atoi(quantity)
		if err != nil {
			c.String(http.StatusBadRequest, "error in conversion")
		}
		st.Quantity = quan
		err = config.DB.C("stuff").Insert(&st)
		if err != nil {
			c.String(http.StatusBadRequest, "error while inserting data")
		}
	} else {
		c.String(http.StatusUnauthorized, "employee is not authorized for this operation")
	}
}

func GetStuff(c *gin.Context) {

}

func UpdateStuff(c *gin.Context) {

}

func DeleteStuff(c *gin.Context) {

}
