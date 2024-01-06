package data

import "github.com/lauro-ss/practices_go/14_api/pkg/api"

type Animal struct {
	Id    string
	Name  string
	Emoji string
	Foods []Food //Many to Many
	api.Hateoas
}

type Food struct {
	Id      string
	Name    string
	Emoji   string
	animals []Animal //Many to Many
}
