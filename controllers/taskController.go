package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elgael06/curso_1/funtions"
)

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	var rq = r.Method
	fmt.Println(rq)
	if rq == "GET" {
		var allTas = funtions.GetAllTask()
		json.NewEncoder(w).Encode(allTas)
	} else {
		http.Error(w, "Error URL o metodo no definidos", 404)
	}
}
