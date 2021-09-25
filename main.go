package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/elgael06/curso_1/controllers"
)

// func index(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "pages/index.html")
// }

func main() {
	port := os.Getenv("PORT")
	fs := http.FileServer(http.Dir("./pages"))

	// http.HandleFunc("/", index)
	http.HandleFunc("/api/tasks", controllers.GetAllTask)
	http.HandleFunc("/api/getFile", controllers.FileRecivedCVS)
	http.Handle("/", fs)

	if port == "" {
		port = "8000"
	}

	fmt.Println("server run on PORT:", port)

	log.Fatal(http.ListenAndServe((":" + port), nil))
}
