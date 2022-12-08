package user

type RequestUser struct{
	Username string `json:"username"`
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	RoleId string `json:"role_id"`
}
type ResponseUser struct{
	ID uint `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	Email string `json:"email"`
}