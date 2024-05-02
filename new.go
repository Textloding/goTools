package main

import (
	"errors"
	"fmt"
)

func main() {
	//简单输出
	//Printf函数用于格式化输出，它接受一个格式化字符串作为第一个参数，后面跟着一系列可变参数，这些参数将按照格式化字符串中的占位符被替换并输出。
	//格式化字符串中可以包含占位符（如%d、%s、%f等），用于指定要插入值的类型和格式。
	//Printf不会自动在输出末尾添加换行符，因此如果需要换行，你需要在格式化字符串中显式地添加\n。
	fmt.Printf(study1()) //控制台输出 study1

	//计算相加输出 要么加格式说明符[占位符]%d 要么用strconv.Itoa转换为字符串 不可以直接输出
	fmt.Printf("%d", study2(1, 2)) //控制台输出3

	//引入符号计算输出
	// 如果计算没有错误，则打印result
	if result, err := study3(1, 2, "*"); err == nil {
		// 如果没有错误，则打印result
		fmt.Printf("%d\n", result) //控制台输出2
	} else {
		// 如果有错误，则处理错误（这里只是简单地打印错误）
		////Println函数接受任意数量的参数，并将它们以空格分隔的形式输出到标准输出。
		//		//它会在每个参数之后（以及整个输出之后）自动添加一个换行符。
		//		//Println不接受格式化字符串，因此你不能指定值的输出格式。
		e := 123
		fmt.Println("Error:", err, "study3", e)
		//当调用 study3(1, 2, "**") 控制台输出 Error: 您的计算符号**不正确 study3 123
	}

}

func study1() string {
	return "study1"
}

func study2(a, b int) int {
	return a + b
}

func study3(a, b int, c string) (int, error) {
	switch c {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			//errors.New 简单报错不能用占位符和变量
			return 0, errors.New("除数不能为0")
		}
		return a / b, nil
	default:
		//fmt.Errorf 复杂报错，可以使用占位符和报错
		return 0, fmt.Errorf("您的计算符号%s不正确", c)

	}
}

// 力扣两数之和哈希表解法
func twoSumHash(nums []int, target int) []int {
	//创建键值都为整数的哈希表
	hashTable := map[int]int{}
	//遍历nums数组
	for i, x := range nums {
		//检查(target - x)是否存在哈希表中
		if p, ok := hashTable[target-x]; ok {
			//如果存在，返回两数索引
			return []int{p, i}
		}
		//如果不存在，将当前元素x的值和索引i存进哈希表
		hashTable[x] = i
	}
	return nil
}

// 力扣两数之和暴力解法
func twoSum(nums []int, target int) []int {
	//双重循环相加匹配结果
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
