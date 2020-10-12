package errno

import "fmt"

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("Error(%s):%s", err.Code, err.Message)
}
