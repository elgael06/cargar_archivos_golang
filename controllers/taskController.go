package controller

import (
	"encoding/json"
	"net/http"
)

func getAllTask(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("")
}
