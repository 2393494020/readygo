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
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		arr[i] = randGenerator.Intn(100)
	}

	arrTmp := arr

	fmt.Printf("冒泡排序前:%d\n", arr)
	sort.BubbleSort(&arr)
	fmt.Printf("冒泡排序后:%d\n", arr)
	
	fmt.Printf("快速排序前:%d\n", arrTmp)
	sort.QuickSort(&arrTmp, 0, arrLen - 1)
	fmt.Printf("快速排序后:%d\n", arrTmp)
	
	now := time.Now()
	fmt.Printf("%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Month(), now.Hour(), now.Minute(), now.Second())
	
}