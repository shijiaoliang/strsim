package strsim

import (
	"github.com/shijiaoliang/strsim/similarity"
)

func Hamming() OptionFunc {
	return OptionFunc(func(o *option) {

		h := &similarity.Hamming{}
		o.cmp = h.CompareUtf8
		if o.ascii {
			o.cmp = h.CompareAscii
		}
	})
}
