//Implemention of reddit order
//http://www.ruanyifeng.com/blog/2012/03/ranking_algorithm_reddit.html
package reddit

import (
	"math"
	"time"
)

//Calculate the score, according to ups and down
func score(ups int, down int) int {
	return ups - down
}

func Hot(ups int, down int, date time.Time) int {
	//赞成票和反对票的差
	s := score(ups, down)

	order := int(math.Log(math.Max(math.Abs(float64(s)), 1)))

	//投票方向
	var sign int = 1

	if s == 0 {
		sign = 0
	} else if s < 0 {
		sign = -1
	}

	duration := time.Now().Sub(date)

	return math.Floor(order + sign*duration.Seconds()/45000)
}
