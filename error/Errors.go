package error

import "fmt"

type NaverError struct {
	Message string
}

func (e *NaverError) Error() string {
	return fmt.Sprintf("error cause : %v", e.Message)
}
