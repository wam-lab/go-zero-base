package logic

import (
	"fmt"
	"testing"
)

func TestGenerateVerifyCode(t *testing.T) {
	fmt.Println(GenerateVerifyCode(6))
}
