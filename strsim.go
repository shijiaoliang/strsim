package strsim

import (
	"github.com/shijiaoliang/similarity"
)

// Compare 比较两个字符串相似度
func Compare(s1, s2 string, opts ...Option) float64 {
	var o option

	o.fillOption(opts...)

	return compare(s1, s2, &o)
}

// FindBestMatchOne 返回相似度最高的那个字符串
func FindBestMatchOne(s string, targets []string, opts ...Option) *similarity.Match {
	r := findBestMatch(s, targets, opts...)
	return r.Match
}

// FindBestMatch 返回相似度最高的那个字符串, 以及索引位置
func FindBestMatch(s string, targets []string, opts ...Option) *similarity.MatchResult {
	return findBestMatch(s, targets, opts...)
}

// IsMatch 是否存在相似度大于等于score的文本
func IsMatch(s string, targets []string, score float64, opts ...Option) bool {
	return isMatch(s, targets, score, opts...)
}
