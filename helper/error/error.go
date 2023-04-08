package cerror

import (
	"fmt"
	"strconv"
	"strings"
)

type customError struct {
	code         int
	errorMessage string
	actualError  string
}

func New(code int, errorMessage, actualError string) customError {
	return customError{
		code:         code,
		errorMessage: errorMessage,
		actualError:  actualError,
	}
}
func (c customError) Error() string {
	return fmt.Sprintf("%d|%s|%s", c.code, c.errorMessage, c.actualError)
}

func ExtractCustomError(err error) (response customError, ok bool) {
	splitError := strings.Split(err.Error(), "|")
	if ok = len(splitError) == 3; ok {
		response.code, _ = strconv.Atoi(splitError[0])
		response.errorMessage = splitError[1]
		response.actualError = splitError[2]
	}
	return
}

func (c customError) GetCode() int {
	return c.code
}

func (c customError) GetErrorMessage() string {
	return c.errorMessage
}

func (c customError) GetActualError() string {
	return c.actualError
}
