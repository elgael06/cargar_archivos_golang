package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FileRecivedCVS(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(20000)
	file, info, err := r.FormFile("cvs")
	fmt.Println("err->", err)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	var buf bytes.Buffer
	defer file.Close()
	fmt.Println("name->", info.Filename)
	fmt.Println("file->", file)

	io.Copy(&buf, file)

	contents := buf.String()
	fmt.Println(contents)
	buf.Reset()
	json.NewEncoder(w).Encode(contents)
}
