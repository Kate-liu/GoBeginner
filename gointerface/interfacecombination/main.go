package main

import (
	"fmt"
	"net/http"
)

// func Add(a int64, b int64) int64 {
// 	return a + b
// }
//
// type Adder interface {
// 	Add(int64, int64) int64
// }
//
// func Add(adder Adder, a int64, b int64) int64 {
// 	return adder.Add(a, b)
// }
//
// type MyReader struct {
// 	io.Reader       // underlying reader
// 	N         int64 // max bytes remaining
// }
//
// // func Save(f *os.File, data []byte) error
//
// func (f *File) Chdir() error
// func (f *File) Chmod(mode FileMode) error
// func (f *File) Chown(uid, gid int) error
// ... ...
//
// func Save(w io.Writer, data []byte) error
//
// func TestSave(t *testing.T) {
// 	b := make([]byte, 0, 128)
// 	buf := bytes.NewBuffer(b)
// 	data := []byte("hello, golang")
// 	err := Save(buf, data)
// 	if err != nil {
// 		t.Errorf("want nil, actual %s", err.Error())
// 	}
//
// 	saved := buf.Bytes()
// 	if !reflect.DeepEqual(saved, data) {
// 		t.Errorf("want %s, actual %s", string(data), string(saved))
// 	}
// }
//
// func YourFuncName(param YourInterfaceType)
//
// func YourWrapperFunc(param YourInterfaceType) YourInterfaceType

// YourWrapperFunc1(YourWrapperFunc2(YourWrapperFunc3(...)))

// func main() {
// 	r := strings.NewReader("hello, gopher!\n")
// 	lr := io.LimitReader(r, 4)
// 	if _, err := io.Copy(os.Stdout, lr); err != nil {
// 		log.Fatal(err)
// 	}
// }

func greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(greetings))
}
