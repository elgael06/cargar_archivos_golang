package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elgael06/curso_1/funtions"
)

func FileRecivedCVS(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(20000)
	file, info, err := r.FormFile("cvs")
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	fmt.Println("name->", info.Filename)

	var dataFile string = funtions.ReaderFile(file)
	content := funtions.SplitString(dataFile, "\n")
	orderData := funtions.GetHeaderFileToJSON(content)
	json.NewEncoder(w).Encode(orderData)
}
