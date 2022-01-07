package secrets

const DbUser = "USER"
const DbPassword = "PASSWORD"

func GetUser() string {
	return DbUser
}

func GetPassword() string {
	return DbPassword
}
