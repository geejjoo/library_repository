package migrate

import (
	"fmt"
	"testing"
)

func TestZipper(t *testing.T) {
	strings := Zipper([]string{"1", "3", "5"}, []string{"2", "4", "6"})

	fmt.Println(strings)
}
