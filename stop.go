package stop

import (
	"fmt"
	"os"
)

//func Error(m string) {
//	return error.Error(m)
//}

// User error, tell them and exit.
func Mistake(m string) {
	fmt.Println(m);
	os.Exit(1)
}

func Fatal(m string) {
	panic(m)
}
