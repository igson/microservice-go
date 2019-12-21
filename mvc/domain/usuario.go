package domain

type Usuario struct {
	ID        uint64 `json:"id"`
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email     string `json:"email"`
}
