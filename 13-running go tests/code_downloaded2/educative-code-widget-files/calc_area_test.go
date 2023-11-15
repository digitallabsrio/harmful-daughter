package main

import (
  "time"
  "testing"
)


func TestCalcAreaInParallel(t *testing.T) {
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
    t.Run("", func(tt *testing.T) {
      tt.Parallel()
      time.Sleep(time.Second)
      w := test.width
      h := test.height
      r, err := CalcArea(w, h)
      if err != nil {
        tt.Errorf("CalcArea(%d, %d) returned an error", w, h)
      } else if r != test.expected {
        tt.Errorf("CalcArea(%d, %d) returned %d. Expected %d", 
                  w, h, r, test.expected)
      }
    })
  }  
}