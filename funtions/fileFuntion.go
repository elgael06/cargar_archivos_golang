package funtions

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/elgael06/curso_1/structs"
)

func ReaderFile(file multipart.File) string {
	var buf bytes.Buffer
	defer file.Close()
	io.Copy(&buf, file)
	contents := buf.String()
	buf.Reset()
	return contents
}

func SplitString(content string, sep string) []string {
	resList := strings.Split(content, sep)
	return resList
}

func GetHeaderFileToJSON(data []string) structs.FileDataStruct {
	fmt.Println("filas->", len(data))
	headers := SplitString(data[0], ",")
	fmt.Println("headers->", len(headers))
	lentArray := (len(data) - 1)
	fmt.Println("lentArray->", lentArray)

	contents := make([][]string, lentArray)

	for i, s := range data {
		if i > 0 {
			// contents[i] = SplitString(s, ",")
			contetRow := SplitString(s, ",")
			fmt.Println("contetRow->", i, contetRow)
			contents[i-1] = contetRow
		}
	}

	Dic := structs.FileDataStruct{
		Header:  headers,
		Content: contents,
	}
	fmt.Println("dic->", Dic)
	return Dic
}
