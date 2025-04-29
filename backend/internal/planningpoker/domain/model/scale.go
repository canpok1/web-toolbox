package model

import "github.com/canpok1/web-toolbox/backend/internal/api"

var ScaleListMap = map[api.ScaleType][]string{
	api.Fibonacci:  {"0", "1", "2", "3", "5", "8", "13", "21", "34", "55", "89", "?"},
	api.TShirt:     {"XS", "S", "M", "L", "XL", "XXL", "?"},
	api.PowerOfTwo: {"1", "2", "4", "8", "16", "32", "64", "128", "256", "512", "1024", "?"},
	api.Custom:     {},
}

var ScaleOrder = []string{
	"0", "1", "2", "3", "4", "5", "8", "13", "16", "21",
	"32", "34", "55", "64", "89", "128", "256", "512", "1024",
	"XS", "S", "M", "L", "XL", "XXL", "?",
}
