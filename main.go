package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	dsn := "root:password@tcp(object_store)/store"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		panic(err)
	}

	r := gin.Default()
	r.POST("/user", createUser)
	r.Run()
}

type createUserArgs struct {
	Username  string
	FirstName string
	LastName  string
	Email     string
}

func createUser(c *gin.Context) {
	var args createUserArgs
	if err := c.ShouldBindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	query := fmt.Sprintf(`INSERT INTO Users (Username, FirstName, LastName, Email) VALUES ("%s", "%s", "%s", "%s");`,
		args.Username,
		args.FirstName,
		args.LastName,
		args.Email,
	)

	_, err := db.Exec(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "success"})
}
