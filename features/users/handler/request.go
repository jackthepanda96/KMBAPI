package handler

type RegisterInput struct {
	Nama     string `json:"nama"`
	Password string `json:"password"`
	HP       string `json:"hp"`
}

type LoginInput struct {
	HP       string `json:"hp"`
	Password string `json:"password"`
}
