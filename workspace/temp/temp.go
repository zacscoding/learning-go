package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// checkPaths()
	// checkAddress()
	// checkMath()
	// checkNumbers()
	// checkTimes()
	// checkStrings()
	// checkStrings2()
	writeTemp()

}

func writeTemp() {
	dir := "/home/zaccoding/workspaces/berith/testnetkey/keystore"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Printf("- %s\n", f.Name())
		fmt.Println("```")
		b, _ := ioutil.ReadFile(filepath.Join(dir, f.Name()))
		fmt.Println(string(b))
		fmt.Println("```")
		fmt.Println()
	}
}

func checkStrings2() {
	var buff bytes.Buffer
	buff.WriteString("aa")
	buff.WriteString("bb")

	s1 := buff.String()
	s2 := buff.String()

	fmt.Println("s1 :", s1)
	fmt.Println("s2 :", s2)
}

func checkStrings() {
	tests := []struct {
		Input string
		Len   int
	}{
		{
			`aaaa`,
			4,
		},
		{
			`"aaa`,
			3,
		},
		{
			`aaa"`,
			3,
		},
		{
			`"aa"`,
			2,
		},
	}

	for i, tt := range tests {
		r1 := strings.ReplaceAll(tt.Input, "\"", "")
		r2 := strings.Replace(tt.Input, "\"", "", -1)

		fmt.Println("test : ", tt, "--> r1 :", r1, ", r2 :", r2)

		if r1 != r2 || len(r1) != tt.Len || len(r2) != tt.Len {
			log.Fatalf("find diff result in #%d. input : %s\n", i, tt.Input)
		}
	}
}

func checkTimes() {
	var p uint64
	p = 20
	parentTime := big.NewInt(time.Now().Unix())
	currentTime := new(big.Int).Add(parentTime, new(big.Int).SetUint64(p))
	delay := time.Duration(2) * time.Second

	fmt.Println(parentTime.Uint64())
	fmt.Println(currentTime.Uint64())
	fmt.Println(delay.Seconds())

	fmt.Println("----------------------")
	fmt.Println(parentTime.Int64())
	fmt.Println(int64(p))
	fmt.Println(int64(delay.Seconds()))
}

func checkNumbers() {
	//e := big.NewInt(5 + 1)
	//for i := 0; i < 10; i++ {
	//	number := big.NewInt(int64(i))
	//	d := new(big.Int).Mod(number, e).Uint64()
	//	fmt.Printf("# Check %d -> %d\n", i, d)
	//}
	idx := uint64(18446744073709551615 - 5)
	for i := 0; i < 10; i++ {
		fmt.Println(idx, int(idx))
		idx++
	}
}

func checkMath() {
	numbers := []float64{
		0.1,
		0.2,
		0.3,
		0.4,
		0.5,
		0.6,
		0.7,
		0.8,
		0.9,
	}

	for _, n := range numbers {
		r := math.Round(n)
		result := uint64(r)
		fmt.Printf("%f -> %f -> %d\n", n, r, result)
	}
}

func checkPaths() {
	paths := [] string{
		"~/test.pem",
		"/home/zaccoding/test.mv.db",
		"/home/zaccoding/not_found.go",
	}

	for _, path := range paths {
		path, _ := filepath.Abs(path)
		path = filepath.Clean(path)
		fi, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("Not exist")
			} else {
				fmt.Println(">> error ", err)
			}
		} else {
			fmt.Printf(">> %v\n", fi)
		}
	}
}

type Address [20]byte

func (a Address) MarshalText() ([]byte, error) {
	prefix := "PR"
	b := a[:]
	result := make([]byte, len(b)*2+2)
	copy(result, prefix)
	hex.Encode(result[2:], b)
	return result, nil
}

func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-20:]
	}
	copy(a[20-len(b):], b)
}

func checkAddress() {
	b, _ := hex.DecodeString("d8a25ff31c6174ce7bce74ca4a91c2e816dbf91e")
	var addr Address
	addr.SetBytes(b)
	m, err := addr.MarshalText()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(m))
}
