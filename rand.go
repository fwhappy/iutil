package iutil

import (
	"bytes"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// rand()
func Rand(min, max int) int {
	if min > max {
		panic("min: min cannot be greater than max")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(math.MaxInt32)
	return n/((math.MaxInt32+1)/(max-min+1)) + min
}

// RandIntn 获取一个 0 ~ n 之间的随机值
func RandIntn(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Intn(n)
}

// GetRandString 生成n位随机数字字符串
func GetRandString(n int) string {
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		buffer.WriteString(strconv.Itoa(RandIntn(10)))
	}

	return buffer.String()
}
