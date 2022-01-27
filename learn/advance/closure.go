package advance

import "fmt"

type addOneFunc func() int

func getSequence() addOneFunc {
	i := 0
	return func() int {
		i++
		return i
	}
}

func ClosureTest() {
	/* nextNumber 为一个函数，i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}
