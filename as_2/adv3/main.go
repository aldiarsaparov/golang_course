package main
// aldiar saparov
import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/aldik/crud_go/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/api/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/api/user/{id}", handlers.CreateUser).Methods("PUT")
	r.HandleFunc("/api/user/{id}", handlers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
