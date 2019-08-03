package sort

import (
	// "fmt"
)

func BubbleSort(arr *[30]int) {
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
}

func QuickSort(arr *[30]int, l int, h int) {
	if l < h {
		x := (*arr)[l]
		i := l
		j := h
		for i < j {
			for i < j {
				if (*arr)[j] < x {
					(*arr)[i] = (*arr)[j]
					break
				}

				j--
			}
			// if (*arr)[j] < x {
			// 	(*arr)[i] = (*arr)[j]
			// }

			for i < j {
				if (*arr)[i] > x {
					(*arr)[j] = (*arr)[i]
					break
				}

				i++
			}
			// if (*arr)[i] > x {
			// 	(*arr)[j] = (*arr)[i]
			// }

			(*arr)[i] = x
		}	
		
		QuickSort(arr, l, i - 1)
		QuickSort(arr, i + 1, h)
	}
}