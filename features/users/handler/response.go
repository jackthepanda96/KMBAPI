package handler

type RegisterResponse struct {
	Nama string `json:"nama"`
	HP   string `json:"hp"`
}

type LoginResponse struct {
	Nama  string `json:"nama"`
	Token any    `json:"token"`
}
