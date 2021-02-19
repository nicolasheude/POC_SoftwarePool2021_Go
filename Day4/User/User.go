package User

type User struct {
	Email     string
	Password  string
}

var UserDB = []User{}

type UserJWT struct {
	Email     string
	Password  string
}

var UsersJWT = []UserJWT{}

var APISECRET = 0
