package errors

import "fmt"

var (
	TokenTypeError        = fmt.Errorf("token type unknown")
	TokenExtractUserError = fmt.Errorf("type assertion to user err")
)
