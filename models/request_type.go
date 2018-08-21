package models

type LoginFormData struct {
	Email string `json: "email"`
	Pwd string `json: "pwd"`
}
type CreateLoginFormDate struct {
	LoginForm LoginFormData `json: "loginForm"`
	Name string `json: "name"`
}
type LoginResult struct {
	Result bool `json: "result"`
	Token string `json: "token"`
}