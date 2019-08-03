package sort

import (
	"fmt"
)

func BubbleSort(arr *[30]int)  {
	fmt.Printf("排序前:%d\n", (*arr))

	arrLen := len(*arr)

	for i := 0; i < arrLen - 1; i++ {
		for j := i + 1; j < arrLen; j++ {
			if (*arr)[i] > (*arr)[j] {
				tmp := (*arr)[i]
				(*arr)[i] = (*arr)[j]
				(*arr)[j] = tmp
			}
		}
	}

	fmt.Printf("排序后:%d\n", (*arr))
}