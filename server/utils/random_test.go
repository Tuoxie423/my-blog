package utils

import (
	"fmt"
	"testing"
)

func TestGenerateVerificationCode(t *testing.T) {
	fmt.Println(GenerateVerificationCode(6))
}
