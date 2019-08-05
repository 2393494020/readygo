package main

import (
	"fmt"
	_ "strings"
	"time"
	"math/rand"
	"github.com/2393494020/readygo/sort"
	"github.com/2393494020/readygo/array"
)

func main00() {
	seed := rand.NewSource(time.Now().UnixNano())
    randGenerator := rand.New(seed)
	
	size := 10000
	arr := make([]int, size, size)	
	intSlice := make([]int, size, size)

	for i := 0; i < size; i++ {
		arr[i] = randGenerator.Intn(size * 10)
		intSlice[i] = arr[i]
	}

	fmt.Printf("冒泡排序前:%d\n", arr[10:20])
	sort.BubbleSort(&arr)
	fmt.Printf("冒泡排序后:%d\n", arr[10:20])
	
	fmt.Printf("快速排序前:%d\n", intSlice[10:20])
	sort.QuickSort(&intSlice, 0, size - 1)
	fmt.Printf("快速排序后:%d\n", intSlice[10:20])
	
	now := time.Now()
	fmt.Printf("%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Month(), now.Hour(), now.Minute(), now.Second())
}

func main() {
	array.SparseArray()
}