package entities

type Users []User

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
