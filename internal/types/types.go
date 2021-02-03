package types

type Users struct {
	Users []User `json:"users"`
}
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
