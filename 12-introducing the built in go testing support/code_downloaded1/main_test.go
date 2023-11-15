package main
 
import (
    "errors"
    "testing"
)
 
var errorMessage = "width and height must be positive"
 
// Calculate the area of a rectangle
func CalcArea(w, h int) (int, error) {
    if w < 1 || h < 1 {
      return 0, errors.New(errorMessage)
    }
    return w * h, nil
}

func TestCalcAreaViaTable(t *testing.T) {
  var tests = []struct {
      width    int
      height   int
      expected   int
  }{
    {1, 1, 1},
    {5, 6, 30},
    {1, 99, 99},
    {7, 6, 42},
  }

  for _, test := range tests {
    w := test.width
    h := test.height
    r, err := CalcArea(w, h)
    if err != nil {
      t.Errorf("CalcArea(%d, %d) returned an error", w, h)
    } else if r != test.expected {
      t.Errorf("CalcArea(%d, %d) returned %d. Expected %d", 
                w, h, r, test.expected)
    }
  }  
}