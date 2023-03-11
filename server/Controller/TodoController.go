package Controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server/Repository"
)

func TodoController() {
	router := mux.NewRouter()
	router.HandleFunc("/", TopPage)
	router.HandleFunc("/todos", Repository.GetTodos)
	router.HandleFunc("/todos/{id}", Repository.GetDetailTodo)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func TopPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the HomePage!")
	fmt.Println("Endpoint Hit: Page")
}
