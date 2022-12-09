package user

import "ddrag23/gae-soal/model"

type ResponseUser struct {
	ID       uint       `json:"id"`
	Username string     `json:"username"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Role     model.Role `json:"role"`
}

type UpdateUser struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	RoleId   uint   `json:"role_id"`
}

type ChangePasswordUser struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
