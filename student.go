package main
type Student struct {
 id uint `json:"id"`
 FirstName string `json:"first_name"`
 LastName string `json:"last_name"`
 Age int `json:"age"`
 Email string `json:"email"`
}