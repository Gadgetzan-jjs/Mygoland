package ClassConstructor
import (
	"fmt"
	"math/rand"
	"os"
	"time"
)
func Writefile() {
	urlFile := "/home/mikasa/go/src/awesomeProject/tokyo.txt"
	fout, err := os.Create(urlFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(urlFile, err)
		return
	}
	var i int
	for i = 0; i < 100; i++ {
		if i%10 == 0 {
			if i%3 == 0 {
				string := getRandomString() + "=" + "false\n"
				fout.WriteString(string)
			} else {
				string := getRandomString() + "=" + "true\n"
				fout.WriteString(string)
				//fout.WriteString()
			}
		} else if i%7 == 0 {
			string := getRandomString() + "=" + getRandomInt() + "\n"
			fout.WriteString(string)
		} else {
			string := getRandomString() + "=" + getRandomString() + "\n"
			fout.WriteString(string)
		}
	}
	//fmt.Println(urlFile)
	//readFile(urlFile, typeMap)
}
func getRandomString() string {
	str := "ABCDEFGHIJKLMNOPQRSTUYWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i < 5; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
func getRandomInt() string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 5; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}