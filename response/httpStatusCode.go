package response

const (
	ErrCodeSuccess       = 2001 // Success
	ErrCodeParamInvalid  = 2003 // Email is invalid
	ErrCodeInvalidToken  = 3001 // token is invalid
	ErrInvalidOtp        = 3002 // otp is invalid
	ErrSendEmailOtp      = 3003
	ErrCodeUserHasExists = 5001 // user has already registered
)

// message
var msg = map[int]string{
	ErrCodeSuccess:       "Success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrCodeInvalidToken:  "Token is invalid",
	ErrCodeUserHasExists: "User has already registered",
	ErrInvalidOtp:        "OTP is invalid",
	ErrSendEmailOtp:      "Send email OTP error",
}
