package main

import (
	"fmt"
	_ "strings"
	"time"
	"math/rand"
	"github.com/2393494020/readygo/sort"
)

func main()  {
	seed := rand.NewSource(time.Now().UnixNano())
    randGenerator := rand.New(seed)
	
	arr := [30]int{}
	for i := 0; i < len(arr); i++ {
		arr[i] = randGenerator.Intn(100)
	}
	
	sort.BubbleSort(&arr)

	now := time.Now()
	fmt.Printf("%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Month(), now.Hour(), now.Minute(), now.Second())
}