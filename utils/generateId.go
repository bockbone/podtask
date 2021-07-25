package utils

import (
	"fmt"
	"math/rand"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func GenerateId() string {

	timestamp := timestamppb.Now()
	hex := fmt.Sprintf("%x", timestamp.Seconds)

	randomNumber := sliceToInt(randArray())
	hexRandomNum := fmt.Sprintf("%x", randomNumber)

	id := hex + hexRandomNum

	return id

}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}

func randArray() []int {

	rand.Seed(time.Now().UnixNano())
	var len int = 16
	a := make([]int, len)

	for i := 0; i <= len-1; i++ {

		a[i] = rand.Intn(9)
	}

	return a
}
