package domain

//User struct
type User struct {
	Name     string `json:"name" firestore:"name"`
	Email    string `json:"email" firestore:"email"`
	Password string `json:"password" firestore:"password"`
}
