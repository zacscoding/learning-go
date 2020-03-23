package bytestrings

import "fmt"

// WorkWithBuffer는 Buffer 함수에서 생성한 버퍼를 사용 할 것
func WorkWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array"

	b := Buffer(rawString)

	fmt.Println(b.String())

	s, err := toString(b)
	if err != nil {
		return err
	}

}
