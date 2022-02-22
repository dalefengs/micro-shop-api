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
