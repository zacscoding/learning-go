package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Buffer는 바이트 버퍼를 초기화하는 기법을 몇가지 보여줌
// 이 버퍼들은 io.Reader 인터페이스를 구현
func Buffer(rawString string) *bytes.Buffer {
	rawBytes := []byte(rawString)

	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	b = bytes.NewBuffer(rawBytes)

	b = bytes.NewBufferString(rawString)

	return b
}

// toString은 io.Reader로부터 데이터를 가져와서 모두 사용하고 문자열을 반환
func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
