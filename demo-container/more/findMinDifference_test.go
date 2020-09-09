package more

import (
	"fmt"
	"testing"
)

func TestFindMinDifference(t *testing.T) {
	arr := []string{"08:01", "00:17", "14:12", "12:44", "06:06"}
	ret := findMinDifference(arr)
	fmt.Println(ret)
}
