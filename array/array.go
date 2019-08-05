package array

import (
	"fmt"
)

func SparseArray()  {
	chess := [11][11]int{}
	chess[1][2] = 1
	chess[2][3] = 2
	chess[3][4] = 3

	singles := 0
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			fmt.Printf("%d\t", chess[i][j])
			if (chess[i][j] != 0) {
				singles ++
			}
		}
		fmt.Println()
	}

	fmt.Println(singles)

	sparseArray := [4][3]int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d\t", sparseArray[i][j])
		}
		fmt.Println()
	}
}