package structs

type FileDataStruct struct {
	Header  []string   `json:header`
	Content [][]string `json:content`
}

type PriceStruct struct {
	Rybiye string `json:rybiye`
	Loin   string `json:loin`
	Chuck  string `json:chuck`
}

type ExportData struct {
	Subject   string      `json:rybiye`
	Sugestion string      `json:Sugestion`
	Tags      string      `json:Tags`
	Response  string      `json:Response`
	Replies   string      `json:Replies`
	CreatedAt string      `json:createdAt`
	Prices    PriceStruct `json:Prices`
}
