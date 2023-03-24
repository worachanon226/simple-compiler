package entities

type ClientCode struct {
	Lang string `json:"lang"`
	Code string `json:"code"`
}

type ClientOutput struct {
	Output string `json:"output"`
}
