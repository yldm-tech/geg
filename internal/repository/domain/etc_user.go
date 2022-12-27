package domain

type ETCUser struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
