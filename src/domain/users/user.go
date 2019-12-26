package users

type Users struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	NickName  string `json:"nick_name"`
	Age       int    `json:"age"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
