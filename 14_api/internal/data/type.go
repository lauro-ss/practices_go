package data

type Animal struct {
	Id    string
	Name  string
	Emoji string
	Foods []Food //Many to Many
}

type Food struct {
	Id      string
	Name    string
	Emoji   string
	Animals []Animal //Many to Many
}
