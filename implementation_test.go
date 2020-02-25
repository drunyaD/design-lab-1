package lab1

import (
	"fmt"
	"testing"
	."gopkg.in/check.v1"
)

func TestPrefixCalculate(t *testing.T) { TestingT(t) }
	
	type MySuite struct{}
    
    var _ = Suite(&MySuite{})

    func (s *MySuite) TestComplexExpression(c *C) {
    res, _ := PrefixCalculate("- * / 15 - 7 + 1 1 3 + 2 + 1 1")
    c.Assert(res, Equals, "5")
    fmt.Println(res)
   }
   
   func (s *MySuite) TestEmptyString(c *C) {
    _ , err := PrefixCalculate("")
    fmt.Println(err)
   }

   func (s *MySuite) TestImproperSymbols(c *C) {
    _ , err := PrefixCalculate("☆☆☆☆☆☆")
    fmt.Println(err)
   }
	
	func (s *MySuite) TestSimpleExpression(c *C) {
    res, _ := PrefixCalculate("+ 5 * - 4 2 3")
    c.Assert(res, Equals, "11")
    fmt.Println(res)
   }

   
    
    

    
   
   