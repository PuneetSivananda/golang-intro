package structs

type Reponse struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}