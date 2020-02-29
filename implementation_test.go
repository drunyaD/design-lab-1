package lab1

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func TestPrefixCalculate(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestExpressionCalculation(c *C) {
	expressions := []string{"+ 5 * - 4 2 3", "- * / 15 - 7 + 1 1 3 + 2 + 1 1"}
	expectations := []string{"11", "5"}
	for i := 0; i < len(expressions); i++ {
		res, _ := PrefixCalculate(expressions[i])
		c.Assert(res, Equals, expectations[i])
		//ftm.Println(res)
	}
}

func (s *MySuite) TestExceptions(c *C) {
	expressions := []string{"", "☆☆☆☆☆☆", "7 + 8 + 9"}
	for i := 0; i < len(expressions); i++ {
		_, err := PrefixCalculate(expressions[i])
		c.Assert(err, NotNil)
		//fmt.Println(err)
	}
}

func ExamplePrefixCalculate() {
	res, err := PrefixCalculate("+ 5 * - 4 2 3")
	if err == nil {
		fmt.Println(res)
	}

	// Output:
	// 11
}
