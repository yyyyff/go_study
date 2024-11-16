package unittest_test

import (
	unittest "github.com/yyyyff/go_study/tree/main/skills/unit_test"
	"testing"
)

func TestSum(t *testing.T) {
	t.Log(unittest.Sum(1, 2))

}
