package structs

type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}
