// Go 언어에서 제공하는 문서화 기능을 보여주기 위한 패키지이다
// 기능은 굉장히 단순하다
package documentme

// Pie는 전역변수이다
// 굳이 주석으로 달 필요가 없는 문장이다
const Pie = 3.141592

// S1() 함수는 스트링의 길이를 구한다.
// 이 함수는 range를 이용하여 스트링에 대해 루프를 돈다.
func S1(s string) int {
	if s == "" {
		return 0
	}
	n := 0
	for range s {
		n++
	}
	return n
}

// F1() 함수는 입력된 정수의 2배수를 리턴한다
// 함수의 이름을 Double()로 지었다면 더 좋았을 것이다
func F1(n int) int {
	return 2 * n
}

/*
$ mkdir ~/go/src/documentme
$ cd /home/zaccoding/go/src/github.com/zacscoding/learning-go/matering-go/ch11/documentme
$ cp document_me* ~/go/src/documentme/
$ ls -l ~/go/src/documentme/
합계 8
-rw-r--r-- 1 zaccoding zaccoding 657  9월  4 14:03 document_me.go
-rw-r--r-- 1 zaccoding zaccoding 218  9월  4 14:03 document_me_test.go
$ go install documentme
$ godoc -http=":8080"
*/
