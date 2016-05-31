package reddit

import (
	"math"
	"testing"
	"time"
)

func TestCalcuScore(t *testing.T) {
	s := score(100, 90)

	if s != 10 {
		t.Fail()
	}
}

func TestLogMath(t *testing.T) {
	res := math.Log10(100)

	t.Log(res)
	println(res)
}

func TestCalculate() {
	ups, down := 100, 19
	date := time.Unix(1464684188)

	println(Hot(ups, down, date))
}
