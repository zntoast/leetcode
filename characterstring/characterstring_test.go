package characterstring

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(T *testing.T) {
	s := "race a car"
	fmt.Println(IsPalindrome(s))
}
