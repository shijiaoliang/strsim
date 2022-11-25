package strsim

import (
	"github.com/shijiaoliang/similarity"
)

// DiceCoefficient ngram 是筛子系数需要用的一个值
func DiceCoefficient(ngram ...int) OptionFunc {
	return OptionFunc(func(o *option) {
		ngram2 := 2
		if len(ngram) > 0 {
			ngram2 = ngram[0]
		}

		d := &similarity.DiceCoefficient{Ngram: ngram2}
		o.cmp = d.CompareUtf8
	})
}
