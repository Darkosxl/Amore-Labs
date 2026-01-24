package auth


type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Id string `json:"id"`
	Role string `json:"role"`
	Organization string `json:"organization"`
}

func createUser(username string, password string, email string, id string, role string, organization string) *User {
	user := User{
		Username: username,
		Password: password,
		Email: email,
		Id: id,
		Role: role,
		Organization: organization,
	}
	return &user
}

func modifyPassword(user *User, newPassword string) {
	user.Password = newPassword
}

func modifyRole(user *User, newRole string) {
	user.Role = newRole
}

func modifyOrganization(user *User, newOrganization string) {
	user.Organization = newOrganization
}
