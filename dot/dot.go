package dot

import (
	"fmt"
)

func Dot() {
	vector1 := []int{123, 456}
	vector2 := []int{700, 570}

	dotProd, err := dotProduct2[int](vector1, vector2)
	if err != nil {
		fmt.Println("错误：", err)
		return
	}

	fmt.Printf("内积结果：%.2f\n", dotProd)
}

// []float64{700, 570, 570, 600, 550, 500, 400, 300, 100, 60, 60, 50, 50, 50, 20, 20, 20, 10, 10, 10},
// 	[]float64{0, 0, 0, 0, 0, 0, 500, 350, 400, 60, 60, 50, 50, 20, 20, 10, 0, 0, 0, 0}, //	21	FISHC06 財神

// func dotProduct(a, b []float64) (float64, error) {
// 	if len(a) != len(b) {
// 		return 0, fmt.Errorf("切片长度不相等")
// 	}

// 	var result float64
// 	for i := 0; i < len(a); i++ {
// 		result += a[i] * b[i]
// 	}
// 	return result, nil
// }

func dotProduct2[T int](vector1, vector2 []T) (T, error) {
	if len(vector1) != len(vector2) {
		return 0, fmt.Errorf("切片长度不相等")
	}

	var result T
	for i := 0; i < len(vector1); i++ {
		result += vector1[i] * vector2[i]
	}
	return result, nil
}
