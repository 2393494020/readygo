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

type Person struct {
	name string
	age int
}

type Student struct {
	Person
	id string
}

func (person *Person) setAge(age int) {
	person.age = age
}

func (student *Student) setId(id string) {
	student.id = id
}

type Human interface {
	greeting() string
}

type Chinese struct {
	word string
}

type English struct {
	word string	
}

func (chinese *Chinese) greeting() string {
	return chinese.word
}

func (english *English) greeting() string {
	return english.word
}

func main() {
	array.SparseArray()
	student := Student{Person{"小明", 32}, "201908010000"}
	student.setAge(35)
	student.setId("201908010001")

	fmt.Println(student)

	chinese := Chinese{"你好！"}
    english := English{"Hello!"}

    fmt.Println(chinese.greeting())
    fmt.Println(english.greeting())
}