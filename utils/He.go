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

func cmdSearchProcess(processNames string) {
    ps, err := GetPidByProcessName(processNames)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if len(ps) == 0 {
        fmt.Println("No processes found with name", processNames)
    } else {
        fmt.Println("PIDs for", processNames, "are:")
        for _, pid := range ps {
            fmt.Println(pid)
        }
    }
}

// 两数之和哈希表解法
func sumHash(nums []int, target int) []int {

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

// todo::14 未理解
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    // 假设第一个字符串是最长公共前缀的候选者
    prefix := strs[0]

    // 遍历字符串数组，更新最长公共前缀
    for _, str := range strs[1:] {
        // 使用 strings.Index 查找当前字符串中 prefix 的索引
        // 如果找不到，或者索引不是 0，说明 prefix 不是公共前缀
        // 更新 prefix 为当前公共前缀和 str 的公共部分
        for strings.Index(str, prefix) != 0 {
            // 如果 prefix 只有一个字符或者为空，那么直接返回
            if len(prefix) == 0 {
                return ""
            }
            // 去掉 prefix 的最后一个字符，继续查找
            prefix = prefix[:len(prefix)-1]
        }
        // 如果 prefix 为空，则直接返回
        if prefix == "" {
            return ""
        }
    }

    // 返回最长公共前缀
    return prefix
}

func ExtractDateFromID(id string) (int, int, int, int, int, int, error) {
    // 提取出生日期字符串
    birthStr := id[6:14]

    // 格式化出生日期
    layout := "20060102" //layout为时间模板
    parsedTime, err := time.Parse(layout, birthStr)
    if err != nil {
        return 0, 0, 0, 0, 0, 0, err
    }

    // 提取出生年、月、日
    birthYear := parsedTime.Year()
    birthMonth := int(parsedTime.Month())
    birthDay := parsedTime.Day()

    // 获取当前日期
    now := time.Now()

    // 提取当前年、月、日
    currentYear := now.Year()
    currentMonth := int(now.Month())
    currentDay := now.Day()

    return birthYear, birthMonth, birthDay, currentYear, currentMonth, currentDay, nil
}

// GetRandomDish 返回 dishes 切片中的一个随机美食
func GetRandomDish() Dish {
    // 初始化随机数种子，确保每次运行的结果不同
    rand.Seed(time.Now().UnixNano())

    // 获取切片的长度
    lenDishes := len(dishes)

    // 生成一个0到lenDishes-1之间的随机索引
    randomIndex := rand.Intn(lenDishes)

    // 返回该索引对应的美食
    return dishes[randomIndex]
}

// 发送post请求
func postJSON(url string, data interface{}) (string, error) {
    // 将数据转换为JSON格式
    jsonData, err := json.Marshal(data)
    if err != nil {
        return "", fmt.Errorf("Error marshalling JSON: %w", err)
    }

    // 创建一个HTTP请求
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return "", fmt.Errorf("Error creating request: %w", err)
    }

    // 设置请求头，指定内容类型为JSON
    req.Header.Set("Content-Type", "application/json")

    // 发送请求
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("Error sending request: %w", err)
    }
    defer resp.Body.Close()

    // 读取响应体
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("Error reading response body: %w", err)
    }

    // 返回响应体字符串
    return string(body), nil
}

func GetTime() time.Time {
    return time.Now()
}

func downloadZip(url string, file string, timeout int) (string, error) {
    // 检查文件是否存在，如果存在则删除
    if _, err := os.Stat(file); err == nil {
        if err := os.Remove(file); err != nil {
            return "", err
        }
    }

    // 如果没有提供文件名，则使用URL中的文件名
    if file == "" {
        file = filepath.Base(url)
    }

    // 确保目录存在
    dir := filepath.Dir(file)
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        if err := os.MkdirAll(dir, 0755); err != nil {
            return "", err
        }
    }

    // 替换URL中的空格为%20
    url = strings.ReplaceAll(url, " ", "%20")

    // 创建HTTP客户端
    client := &http.Client{
        Timeout: timeout,
    }

    // 获取远程文件
    resp, err := client.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    // 创建文件并写入内容
    out, err := os.Create(file)
    if err != nil {
        return "", err
    }
    defer out.Close()

    // 将响应内容写入文件
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return "", err
    }

    return file, nil
}

// GetLocationProvince 获取位置 省
func GetLocationProvince() string {
    return "浙江省"
}

func downloadZip(url string, file string, timeout int) (string, error) {
    // 检查文件是否存在，如果存在则删除
    if _, err := os.Stat(file); err == nil {
        if err := os.Remove(file); err != nil {
            return "", err
        }
    }

    // 如果没有提供文件名，则使用URL中的文件名
    if file == "" {
        file = filepath.Base(url)
    }

    // 确保目录存在
    dir := filepath.Dir(file)
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        if err := os.MkdirAll(dir, 0755); err != nil {
            return "", err
        }
    }

    // 替换URL中的空格为%20
    url = strings.ReplaceAll(url, " ", "%20")

    // 创建HTTP客户端
    client := &http.Client{
        Timeout: timeout,
    }

    // 获取远程文件
    resp, err := client.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    // 创建文件并写入内容
    out, err := os.Create(file)
    if err != nil {
        return "", err
    }
    defer out.Close()

    // 将响应内容写入文件
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return "", err
    }

    return file, nil
}

// 发送post请求
func postJSON(url string, data interface{}) (string, error) {
    // 将数据转换为JSON格式
    jsonData, err := json.Marshal(data)
    if err != nil {
        return "", fmt.Errorf("Error marshalling JSON: %w", err)
    }

    // 创建一个HTTP请求
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return "", fmt.Errorf("Error creating request: %w", err)
    }

    // 设置请求头，指定内容类型为JSON
    req.Header.Set("Content-Type", "application/json")

    // 发送请求
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("Error sending request: %w", err)
    }
    defer resp.Body.Close()

    // 读取响应体
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("Error reading response body: %w", err)
    }

    // 返回响应体字符串
    return string(body), nil
}

func initDB() {
    var err error
    config := mysql.Config{
        User:                 "goblog",
        Passwd:               "123456",
        Addr:                 "127.0.0.1:3306",
        Net:                  "tcp",
        DBName:               "goblog",
        AllowNativePasswords: true,
    }

    // 准备数据库连接池
    db, err = sql.Open("mysql", config.FormatDSN())
    checkError(err)

    // 设置最大连接数
    db.SetMaxOpenConns(100)
    // 设置最大空闲连接数
    db.SetMaxIdleConns(25)
    // 设置每个链接的过期时间
    db.SetConnMaxLifetime(5 * time.Minute)

    // 尝试连接，失败会报错
    err = db.Ping()
    checkError(err)
}

func checkError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog！</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
            "<a href=\"mailto:1102389095@qq.com\">1102389095@qq.com</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {

    w.WriteHeader(http.StatusNotFound)
    fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们修复。</p>")
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request) {

    // 1. 获取 URL 参数
    id := getRouteVariable("id", r)

    // 2. 读取对应的文章数据
    article, err := getArticleByID(id)

    // 3. 如果出现错误
    if err != nil {
        if err == sql.ErrNoRows {
            // 3.1 数据未找到
            w.WriteHeader(http.StatusNotFound)
            fmt.Fprint(w, "404 文章未找到")
        } else {
            // 3.2 数据库错误
            checkError(err)
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprint(w, "500 服务器内部错误")
        }
    } else {
        // 4. 读取成功，显示文章
        tmpl, err := template.ParseFiles("resources/views/articles/show.gohtml")
        checkError(err)

        err = tmpl.Execute(w, article)
        checkError(err)
    }
}

func articlesEditHandler(w http.ResponseWriter, r *http.Request) {

    // 1. 获取 URL 参数
    id := getRouteVariable("id", r)

    // 2. 读取对应的文章数据
    article, err := getArticleByID(id)

    // 3. 如果出现错误
    if err != nil {
        if err == sql.ErrNoRows {
            // 3.1 数据未找到
            w.WriteHeader(http.StatusNotFound)
            fmt.Fprint(w, "404 文章未找到")
        } else {
            // 3.2 数据库错误
            checkError(err)
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprint(w, "500 服务器内部错误")
        }
    } else {
        // 4. 读取成功，显示表单
        updateURL, _ := router.Get("articles.update").URL("id", id)
        data := ArticlesFormData{
            Title:  article.Title,
            Body:   article.Body,
            URL:    updateURL,
            Errors: nil,
        }
        tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
        checkError(err)

        err = tmpl.Execute(w, data)
        checkError(err)
    }
}

func articlesUpdateHandler(w http.ResponseWriter, r *http.Request) {
    // 1. 获取 URL 参数
    id := getRouteVariable("id", r)

    // 2. 读取对应的文章数据
    _, err := getArticleByID(id)

    // 3. 如果出现错误
    if err != nil {
        if err == sql.ErrNoRows {
            // 3.1 数据未找到
            w.WriteHeader(http.StatusNotFound)
            fmt.Fprint(w, "404 文章未找到")
        } else {
            // 3.2 数据库错误
            checkError(err)
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprint(w, "500 服务器内部错误")
        }
    } else {
        // 4. 未出现错误

        // 4.1 表单验证
        title := r.PostFormValue("title")
        body := r.PostFormValue("body")

        errors := validateArticleFormData(title, body)

        if len(errors) == 0 {

            // 4.2 表单验证通过，更新数据

            query := "UPDATE articles SET title = ?, body = ? WHERE id = ?"
            rs, err := db.Exec(query, title, body, id)

            if err != nil {
                checkError(err)
                w.WriteHeader(http.StatusInternalServerError)
                fmt.Fprint(w, "500 服务器内部错误")
            }

            // √ 更新成功，跳转到文章详情页
            if n, _ := rs.RowsAffected(); n > 0 {
                showURL, _ := router.Get("articles.show").URL("id", id)
                http.Redirect(w, r, showURL.String(), http.StatusFound)
            } else {
                fmt.Fprint(w, "您没有做任何更改！")
            }
        } else {

            // 4.3 表单验证不通过，显示理由

            updateURL, _ := router.Get("articles.update").URL("id", id)
            data := ArticlesFormData{
                Title:  title,
                Body:   body,
                URL:    updateURL,
                Errors: errors,
            }
            tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
            checkError(err)

            err = tmpl.Execute(w, data)
            checkError(err)
        }
    }
}
