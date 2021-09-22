package structs

type FileDataStruct struct {
	Header  []string   `json:header`
	Content [][]string `json:content`
}
