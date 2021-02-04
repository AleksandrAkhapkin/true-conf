package types

type Users struct {
	Users []User `json:"users"`
}
type User struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}
