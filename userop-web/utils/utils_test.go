package utils

import (
	"fmt"
	"testing"
)

func TestGenValidateCode(t *testing.T) {
	fmt.Println(GenValidateCode(3))
	fmt.Println(GenValidateCode(2))
	fmt.Println(GenValidateCode(6))
}
func TestGetFreePort(t *testing.T) {
	port, err := GetFreePort()
	fmt.Println(port, err)
}
