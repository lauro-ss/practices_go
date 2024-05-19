package services

type UserRepository interface {
	New(User) (string, error)
	Authentication(login string, password string) (*struct {
		AccessToken  string
		RefreshToken string
	}, error)
}
