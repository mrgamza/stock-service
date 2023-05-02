package helper

import "fmt"

func Encode(keyword string) string {
	length := len(keyword)

	var encode = ""
	for i := 0; i < length; i++ {
		encode += "%" + fmt.Sprintf("%02X", keyword[i])
	}

	return encode
}
