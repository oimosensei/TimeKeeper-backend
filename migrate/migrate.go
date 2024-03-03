package main

import (
	"TverClip/db"
	"TverClip/model"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Migration Done")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{}, &model.Attendance{})
}
