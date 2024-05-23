package utils

import (
    "fmt"
)

//第一天
//第二天
//第三天
//第四天
func main() {
    //简单输出
    //Printf函数用于格式化输出，它接受一个格式化字符串作为第一个参数，后面跟着一系列可变参数，这些参数将按照格式化字符串中的占位符被替换并输出。
    //格式化字符串中可以包含占位符（如%d、%s、%f等），用于指定要插入值的类型和格式。
    //Printf不会自动在输出末尾添加换行符，因此如果需要换行，你需要在格式化字符串中显式地添加\n。
    fmt.Printf(studyFirst()) //控制台输出 study1

    //计算相加输出 要么加格式说明符[占位符]%d 要么用strconv.Itoa转换为字符串 不可以直接输出
    fmt.Printf("%d", studySecond(1, 2)) //控制台输出3

    //引入符号计算输出
    // 如果计算没有错误，则打印result
    if result, err := studyThird(1, 0, "/"); err == nil {
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

//第五天
//今天是个好日子，我还在找工作，但是没有新的公司了，我不知道我多久才可以上班，最近的压力真的很大，如果可以，我想说，lbt我凿死你
//第六天
//依然还是找工作的日子，今天有两个人问我面试，一个九点多早上我还没睡醒，他是之前说如果有面试会优先考虑我，应该是找到了一个心仪的然后想跟我对比吧，应该岗位已经关闭了，还有一个等了一周了
//告诉我初筛通过了，希望能快些上班，穷死了
func cutMyhead(learn bool) string {
    if learn == true {
        return "继续加油"
    } else {
        return "信不信我抽你"
    }
}

//上班第三天
func work(working bool) string {
    if working == true {
        return "干吧孩子"
    } else {
        return "别摸了"
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

// 获取请求中的IP地址，优先从X-Real-IP头中获取，如果没有则从RemoteAddr中获取并解析
func searchIp(r *http.Request) string {
    // 首先检查X-Real-IP头，这通常在代理服务器设置中传递原始客户端IP
    ipStr := r.Header.Get("X-Real-IP")
    if ipStr != "" {
        // 检查IP是否有效
        realIp := net.ParseIP(strings.TrimSpace(ipStr))
        if realIp != nil {
            return realIp.String()
        }
    }

    // 如果X-Real-IP无效或不存在，尝试从RemoteAddr中获取
    host, _, err := net.SplitHostPort(r.RemoteAddr)
    if err == nil {
        remoteIp := net.ParseIP(strings.TrimSpace(host))
        if remoteIp != nil {
            return remoteIp.String()
        }
    }

    // 如果所有尝试都失败，返回空字符串表示无法获取IP
    return ""
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
