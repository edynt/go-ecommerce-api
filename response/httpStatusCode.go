package response

const (
	ErrCodeSuccess      = 2001 // Success
	ErrCodeParamInvalid = 2003 // Email is invalid
	ErrCodeInvalidToken = 3001 // token is invalid
)

// message
var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrCodeInvalidToken: "Token is invalid",
}
