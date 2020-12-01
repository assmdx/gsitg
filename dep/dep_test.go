package dep

import (
	"testing"
	"fmt"
)

func TestAnalysis(t *testing.T) {
	Analysis("github.com/assmdx/gsitg", "../test/gsitg", "../test_result.png")
	fmt.Println("test finish! Open the test_result.png!")
}
