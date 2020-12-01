package dep

import (
	"testing"
)

func TestAnalysis(t *testing.T) {
	Analysis("github.com/assmdx/gsitg", "../test/gsitg", "../test_result.png")
}
