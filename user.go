package dagestan

type User struct {
	id       int    `json:"id"`
	name     string `json:"name"`
	email    string `json:"email"`
	password string `json:"password"`
}
