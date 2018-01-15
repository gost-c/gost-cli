package commands

// User is struct of login api's post data
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Result is struct of common api's response
type Result struct {
	Success bool
	Message string
	Data    interface{}
}
