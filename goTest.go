package goTest

import (
	"fmt"
)

func init() {
	fmt.Println("packTest init1")
}
func init() {
	fmt.Println("packTest init2")
}
func init() {
	fmt.Println("packTest init3")
}
func Once() {
	fmt.Println("packTest Once")
}
