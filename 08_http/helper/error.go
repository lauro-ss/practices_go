package helper

type Error struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
}

func NewError(t string, s int) Error {
	return Error{
		Title:  t,
		Status: s,
	}
}
