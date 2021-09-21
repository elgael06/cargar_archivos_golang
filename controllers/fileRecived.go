package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FileRecivedCVS(w http.ResponseWriter, r *http.Request) {
	file, h, err := r.FormFile("cvs")
	fmt.Println("h->", h)
	fmt.Println("file->", file)
	fmt.Println("err->", err)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("{message:'ok'}")
}
