package utils

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

type ErrorRequest struct {
	Result string `json "result"`
	Code int `json: "code"`
	Message string `json: "message"`
}
