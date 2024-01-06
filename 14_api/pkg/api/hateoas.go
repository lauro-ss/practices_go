package api

import "net/http"

type Hateoas struct {
	Links []Link `json:"links"`
}

type Link struct {
	Rel    string `json:"rel"`
	Action string `json:"action"`
	Uri    string `json:"uri"`
}

func (re *Hateoas) SelfGet(r *http.Request) {
	re.Links = append(re.Links,
		Link{
			Rel:    "self",
			Action: http.MethodGet,
			Uri:    r.Host + r.RequestURI},
	)
}

func (re *Hateoas) SelfPut(r *http.Request) {
	re.Links = append(re.Links,
		Link{
			Rel:    "self",
			Action: http.MethodPut,
			Uri:    r.Host + r.RequestURI},
	)
}

func (re *Hateoas) SelfDelete(r *http.Request) {
	re.Links = append(re.Links,
		Link{
			Rel:    "self",
			Action: http.MethodDelete,
			Uri:    r.Host + r.RequestURI},
	)
}
