package api

import "net/http"

type Response struct {
	Data  any    `json:"data"`
	Links []Link `json:"links"`
}

type Links struct {
	Links []Link `json:"links"`
}

type Link struct {
	Rel    string `json:"rel"`
	Action string `json:"action"`
	Uri    string `json:"uri"`
}

func (re *Links) SelfGet(r *http.Request) {
	re.Links = append(re.Links,
		Link{
			Rel:    "self",
			Action: http.MethodGet,
			Uri:    r.Host + r.RequestURI},
	)
}

func (re *Links) SelfPut(r *http.Request) {
	re.Links = append(re.Links,
		Link{
			Rel:    "self",
			Action: http.MethodPut,
			Uri:    r.Host + r.RequestURI},
	)
}

func (re *Links) SelfDelete(r *http.Request) {
	re.Links = append(re.Links,
		Link{
			Rel:    "self",
			Action: http.MethodDelete,
			Uri:    r.Host + r.RequestURI},
	)
}

func (r *Response) Rel(l Link) {
	r.Links = append(r.Links, l)
}
