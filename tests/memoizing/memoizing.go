package memoizing

import (
	"io/ioutil"
	"net/http"
)

// memoizing 缓存函数
// 缓存函数的返回结果, 这样在对函数进行调用的时候, 我们就只需要一次计算, 之后只要返回计算的结果就可以了.

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
