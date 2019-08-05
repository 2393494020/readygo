package array

import (
	"fmt"
)

func SparseArray()  {
	chess := [11][11]int{}
	chess[1][2] = 1
	chess[2][3] = 2
	chess[3][4] = 3
	chess[3][5] = 3

	singles := 0
	fmt.Println("原始数组")
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			fmt.Printf("%d\t", chess[i][j])
			if (chess[i][j] != 0) {
				singles ++
			}
		}
		fmt.Println()
	}

	sparseArray := [5][3]int{}
	sparseArray[0][0] = 11
	sparseArray[0][1] = 11
	sparseArray[0][2] = singles

	row := 1
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if (chess[i][j] != 0) {
				sparseArray[row][0] = i
				sparseArray[row][1] = j
				sparseArray[row][2] = chess[i][j]
				row ++
			}
		}
	}

	fmt.Println("\n稀疏数组")
	for i := 0; i < singles + 1; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d\t", sparseArray[i][j])
		}
		fmt.Println()
	}

	chessRecovery := [ 11 ][ 11 ]int{}
	for i := 1; i < singles + 1; i++ {
		chessRecovery[ sparseArray[i][0] ][ sparseArray[i][1] ] = sparseArray[i][2]
	}

	fmt.Println("\n恢复后的数组")
	for i := 0; i < sparseArray[0][0]; i++ {
		for j := 0; j < sparseArray[0][1]; j++ {
			fmt.Printf("%d\t", chessRecovery[i][j])
		}
		fmt.Println()
	}
}