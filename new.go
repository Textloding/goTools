package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

//2024.5.7 今天业务繁忙没写力扣没学产品，罪过罪过
//2024.5.13 今天业务不繁忙搞了一天黑群晖但是又没写力扣没学产品没学一点GO，罪孽深重啊
//2024.5.14 又颓废了，学了一节产品没学GO但是深有启发，人其实也可以是产品，一个优秀的产品要有消费场景
//场和景。“场”是时间和空间的概念，⼀个场就是时间加空间。⽤户可以在这个空间⾥停留和消费，如果⼀个⼈不能在某个空间去停留、消费，这个场就是不存在的。
//“景”是什么意思？景就是情景和互动。当⽤户停留在这个空间的时间⾥，要有情 景和互动让⽤户的情绪触发，并且裹挟⽤户的意见，这就是场景。
//场景要能触发情绪
//场景的核⼼是在空间加时间的点上触发别⼈的情绪。
//实⼈所有的理智和意识都是把⼈往回拉的，⽽所有的情绪是在推动⼈的⾏动。所以⼈会去做⼀个动作，会往前⾛，是被情绪推动的。
//你要规划⼀个产品，需要问⾃⼰ 七个问题：
//1.我的产品解决了什么问题？是痛点、痒点，还是爽点？
//2.我在为谁解决这个问题？他得到即刻满⾜了吗？这是⽤户画像的问题。
//3.有多少⼈需要解决这个问题呢？这是市场规模的问题。
//4.⽬前⼈们是怎么解决这个问题的呢？这是竞争分析的问题。
//5.我的竞争⽅案为什么能够在市场竞争中胜出？你不要简单地看单点的竞争 ⼒，⽽是要看点线⾯体，谁给你赋能。
//6.⽤户会在什么样的场景触发情绪？需要马上去解决问题，这是场景问题。
//7.当⽤户遇到问题的时候，他会想到哪个名字呢？
//2024.5.13 今日看业务代码，第一次接触php执行cmd命令查看php的pid进程↓
//$cmd = "ps  -e|grep  " . $task->pid . "|wc  -l";
//$exists = trim(shell_exec("$cmd"), PHP_EOL);
//2024.5.16 对接的文档跟屎一样，连分类都没有直接一个页装所有接口，我就想问问你以后这个接口都不维护吗，xxooxoxo的

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
	if result, err := study3(1, 0, "/"); err == nil {
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

	theDay := cutMyhead(false)
	fmt.Println(theDay)

	thisDay()

	//根据进程名称循环打印出进程pid
	cmdSearchPid("php")

}

func study1() string {
	return "study1"
}

func cutMyhead(study bool) string {
	if study == true {
		return "很好很好，值得鼓励"
	} else {
		return "尼玛又没学习？拖出去车裂"
	}
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
			//errors.New 简单报错不能用占位符
			return 0, errors.New("除数不能为0" + strconv.Itoa(b))
		}
		return a / b, nil
	default:
		//fmt.Errorf 复杂报错，可以使用占位符和报错
		return 0, fmt.Errorf("您的计算符号%s不正确", c)

	}
}

func thisDay() {
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("今天是星期一")
	case time.Tuesday:
		fmt.Println("今天是星期二")
	case time.Wednesday:
		fmt.Println("今天是星期三")
	case time.Thursday:
		fmt.Println("今天是星期四")
	case time.Friday:
		fmt.Println("今天是星期五")
	case time.Saturday:
		fmt.Println("今天是星期六")
	}
}

// GetPidByProcessName 根据进程名称获取PID
func GetPidByProcessName(processName string) ([]int, error) {
	// 执行pidof命令
	cmd := exec.Command("pidof", processName)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run pidof: %v", err)
	}

	// 分割输出，转换为整数切片
	pids := strings.Fields(out.String())
	result := make([]int, 0, len(pids))
	for _, pidStr := range pids {
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			return nil, fmt.Errorf("invalid PID '%s': %v", pidStr, err)
		}
		result = append(result, pid)
	}
	return result, nil
}

func cmdSearchPid(processName string) {
	pids, err := GetPidByProcessName(processName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(pids) == 0 {
		fmt.Println("No processes found with name", processName)
	} else {
		fmt.Println("PIDs for", processName, "are:")
		for _, pid := range pids {
			fmt.Println(pid)
		}
	}
}

// 两数之和哈希表解法
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

// 两数之和暴力解法
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

// 20 有效的括号
// 该方法只能按序判断并且必须左括号开头
func isValid(s string) bool {
	stack := make([]rune, 0) // 创建一个空的 rune 切片作为栈，用于存储待匹配的右括号

	for _, ch := range s { // 遍历输入字符串 s 中的每个字符
		switch ch { // 使用 switch 语句来判断当前字符的类型
		case '(': // 如果字符是左小括号 '('
			stack = append(stack, ')') // 将对应的右小括号 ')' 压入栈中
		case '{': // 如果字符是左大括号 '{'
			stack = append(stack, '}') // 将对应的右大括号 '}' 压入栈中
		case '[': // 如果字符是左方括号 '['
			stack = append(stack, ']') // 将对应的右方括号 ']' 压入栈中
		case ')', '}', ']': // 如果字符是右括号（')', '}', ']'）之一
			if len(stack) == 0 || stack[len(stack)-1] != ch { // 检查栈是否为空，或者栈顶元素是否与当前右括号匹配
				return false // 如果不匹配，返回 false
			}
			stack = stack[:len(stack)-1] // 弹出栈顶元素（即已匹配的右括号）
		}
	}

	return len(stack) == 0 // 如果栈为空，说明所有左括号都找到了匹配的右括号，返回 true；否则返回 false
}

// 35 搜索插入的位置
func searchInsert(nums []int, target int) int {
	//循环
	for index, value := range nums {
		//如果当前值大于或等于target直接输出index
		if value >= target {
			return index
		}
	}
	//跳出循环代表数组内所有的数均小于target，那么直接插入最后一个位置
	return len(nums)
}

// todo:: 1652 拆炸弹 未理解
// decrypt 函数用于解密给定的编码数组。
// 参数 code 是需要解密的整数数组，k 是解密的偏移量。
// 返回值是解密后的整数数组。
func decrypt(code []int, k int) []int {
	// 计算数组长度
	n := len(code)
	// 创建一个长度为 n 的整数切片，并初始化为全 0
	ans := make([]int, n)
	// 如果偏移量 k 为 0，则直接返回全 0 的切片
	if k == 0 {
		return ans
	}
	// 将 code 数组复制一份并拼接到原数组后面，以处理循环解密
	code = append(code, code...)
	// 初始化双指针 l 和 r，用于选择每次解密的子数组
	l, r := 1, k
	// 如果偏移量 k 为负数，调整指针 l 和 r 的初始位置以从数组末尾开始解密
	if k < 0 {
		l, r = n+k, n-1
	}
	// 初始化累积和变量，用于计算每次解密子数组的和
	sum := 0
	// 计算初始子数组的和
	for _, v := range code[l : r+1] {
		sum += v
	}
	// 遍历切片，解密每个位置的元素
	for i := range ans {
		// 将初始子数组的和赋值给每个位置
		ans[i] = sum
		// 更新累积和：减去左指针的元素，加上右指针下一个元素
		sum -= code[l]
		sum += code[r+1]
		// 移动指针
		l, r = l+1, r+1
	}
	return ans
}
