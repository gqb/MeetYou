package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"MeetYou/model"
	"MeetYou/model/auth"
)

func main() {
	dsn := "root:root1234@tcp(127.0.0.1:3306)/meet_you?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("open database error: " + err.Error())
		return
	}
	err = db.AutoMigrate(&auth.Account{})
	if err != nil {
		fmt.Println("auto migrate error")
		return
	}

	app := iris.New()

	app.Post("/login", func(ctx iris.Context) {

		fmt.Println("login invoked")
		var account auth.Account
		ctx.ReadJSON(&account)
		fmt.Print("Get login param userName: " + account.UserName + " ; password : " + account.Password)
		db.Create(&account)
		response := model.BaseResponse[string]{Code: iris.StatusOK, Message: "ok", Data: "login token"}

		ctx.JSON(response)
	})

	app.Get("/query", func(ctx iris.Context) {
		var accounts []auth.Account
		result := db.Find(&accounts)
		if result.Error == nil {
			ctx.JSON(model.BaseResponse[[]auth.Account]{Code: iris.StatusOK, Message: "", Data: accounts})
		} else {
			ctx.JSON(model.BaseResponse[any]{Code: iris.StatusInternalServerError, Message: result.Error.Error(), Data: nil})
		}
	})

	app.Listen(":9999")
}
