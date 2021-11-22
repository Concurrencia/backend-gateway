package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GetRandomLocalHost() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(3)
	temp := fmt.Sprintf("localhost:90%s2", strconv.Itoa(r))
	fmt.Println(temp)
	return temp
}
