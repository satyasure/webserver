package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/user"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/satyasure/webserver/server/action/users"
)

var db *gorm.DB

func init() {
	initializeRDSConn()
	validateRDS()
}

func initializeRDSConn() {
	//read from configuration file
	dat, _ := ioutil.ReadFile("/etc/server.conf")
	fmt.Printf("%s", string(dat))
	m := make(map[string]string)
	json.Unmarshal(dat, &m)
	user := m["user"]         //admin
	password := m["password"] //PJwuu-MbCsEXdU__
	netloc := m["netloc"]     //my-cool-project-db-instance.cozpurlif6yt.us-west-2.rds.amazonaws.com:3306
	database := m["database"] //users

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, netloc, database)
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("%s", err)
	}
}

func validateRDS() {
	//If the users table does not already exist, create it
	if !db.HasTable("users") {
		db.CreateTable(&user.User{})
	}
}

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("/deployment/public", false)))
	r.POST("/users", createUserHandler)
	r.DELETE("/users/:id", deleteUserHandler)
	r.GET("/users/:id", getUserHandler)
	r.GET("/users", listUsersHandler)
	r.OPTIONS("/users", optionsUserHandler)
	r.OPTIONS("/users/:id", optionsUserHandler)
	r.Run()
}

func optionsUserHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE")
	c.Header("Access-Control-Allow-Headers", "origin, content-type, accept")
}
func createUserHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	var req users.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := users.CreateUser(db, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

func deleteUserHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	id := c.Param("id")
	req := users.DeleteUserRequest{ID: id}
	res, err := users.DeleteUser(db, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func listUsersHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	limit := 10
	if c.Query("limit") != "" {
		newLimit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			limit = 10
		} else {
			limit = newLimit
		}
	}
	if limit > 50 {
		limit = 50
	}
	req := users.ListUsersRequest{Limit: uint(limit)}
	res, _ := users.ListUsers(db, &req)
	c.JSON(http.StatusOK, res)
}

func getUserHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	id := c.Param("id")
	req := users.GetUserRequest{ID: id}
	res, _ := users.GetUser(db, &req)
	if res.User == nil {
		c.JSON(http.StatusNotFound, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
