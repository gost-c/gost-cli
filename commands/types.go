package commands

// User is struct of login api's post data
type User struct {
	Username string
	Password string
}

// Result is struct of common api's response
type Result struct {
	Code string `decoder:"code"`
	Msg  string `decoder:"msg"`
}
