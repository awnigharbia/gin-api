package main

import (
	"fmt"
	"ginApp/database"
	"ginApp/modules/auth"
	"ginApp/modules/upload"
	"ginApp/modules/users"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	db, err := database.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return 
	}

	defer db.Close()

	router := gin.Default()
	
	
	users.Users(router)
	auth.Routes(router)
	upload.Routes(router)

	router.Run(":8080")
}