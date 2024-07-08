package outputs

type Login struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Token *Token `json:"token"`
}

type Token struct {
	Value   string `json:"value"`
	Expires int64  `json:"expires"`
}

type SignUp struct {
	ID    string `json:"id"`
	Token *Token `json:"token"`
}
