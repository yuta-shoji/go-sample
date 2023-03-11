package Repository

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"server/DB"
	"server/Model"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	db := DB.ConnectToDB()

	todos := Model.Todos{}
	db.Find(&todos)
	json.NewEncoder(w).Encode(todos)
}

func GetDetailTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		panic("failed to get path variable")
	}

	db := DB.ConnectToDB()

	todo := Model.Todo{}
	db.First(&todo, id)
	fmt.Println(todo.Title)

	if todo.ID != 0 {
		json.NewEncoder(w).Encode(todo)
	} else {
		json.NewEncoder(w).Encode(nil)
	}
}

func hoge() {
	dsn := "host=localhost dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Model.Todo{})
	fmt.Println("Successfully migrated")

	var count int64
	db.Model(&Model.Todo{}).Count(&count)
	if count == 0 {
		db.Create(&Model.Todo{Title: "todo01", Description: "description01", Status: "USNTARTED"})
		db.Create(&Model.Todo{Title: "todo02", Description: "description02", Status: "IN_PROGRESS"})
		db.Create(&Model.Todo{Title: "todo03", Description: "description03", Status: "FINISHED"})
		db.Create(&Model.Todo{Title: "todo04", Description: "description04", Status: "USNTARTED"})
		db.Create(&Model.Todo{Title: "todo05", Description: "description05", Status: "FINISHED"})
	}

	var todos []Model.Todo
	db.Where(Model.Todo{Status: "FINISHED"}).First(&todos)
	fmt.Println(todos)
}
