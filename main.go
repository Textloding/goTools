package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"text/template"
	"time"
	"unicode/utf8"

	"github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()
var db *sql.DB

func liudelong() {
	//这里是龙哥提交的数据
	//2023年5月22日	10:01:28
	//	博涛周一周二来我家住了两天我胖了五斤,之后周三周四我把这五斤减掉了，周六晚上博涛来我家吃饭，今天周一我有胖了两斤。emmmm
	//	上周周五打麻将输了140 周日打麻将 输了100 emmmm
	//2023年5月24日	10:41:45
	//	昨天晚上11点被叫去打台球 玩到2点 三点才睡 我一直赢 平常也就让一让 昨天我躺床上了叫我出来 必须体验一下我的起床气
	//2023年6月14日	14:57:29
	//	最近过得都还可以，就是程程、博涛都回去了，小姜来杭州玩了两天也回去了，打了两天麻将，输了好多哈哈哈，小杨也找到工作了，不知道刘晨工作怎么样了，希望都一切顺利吧。
	//2023年6月19日	09:58:46
	//	周五打麻将赢了100多 昨天打麻将输了300多 周六团建给同事扒小龙虾 把手给烫了 周日吃烤羊腿  把嘴给烫了 今天上班 接下来半个月的主题是防止拖延症
	//2023年6月25日	10:10:42
	// 	端午节放假三天 第一天和同事玩剧本杀，晚上一起吃了烤肉，感觉00后的小姑娘有很大的差别，确实年纪大了不太懂了。这种感觉就是传说中的代沟吧。
	//	第二天，躺了一天很无聊，下午两点多家辉说要去猫咖，小陆提议要去打麻将，最后决定去打麻将，打到10点，输了330多，小陆把台费付了，奶茶钱付了真不错，8月份之前不打麻将了。
        //	第三天，小倪叫我打台球，在上城，约的6.30，躺了一天很累，3点多出门，在钱塘江边逛了一个多小时，江不错，雨下的也很大。9点半去玩了一个密室，三男五女，还不错。
	//2023年7月12日	09:33:19
	//	今天小杨要回东北了，现在应该已经踏上回家的路程了吧，又一个玩的很好的盆友回家了。虽然已经是一个成年人了，也经历了不少的离别，但是伤感的情绪还是在心头弥漫	
        //	杭州今天的天气，一点云彩都没有，特别的蓝，估计也是知道小杨要回家了，所以把前路的云彩都赶走，告诉她前路光明。
	//	写着写着想起来，四天前和小杨打麻将输了162,emmm,小杨回去一定要挣大钱
	//2023年7月17日	17:17:48
	//	上学了DOCKER 这是我至少第五次学习docker简单说就是取得了一点进展，最后还没没整明白。
	//	14号晚上和羽儿打语音打到11点，后来钱程叫我打台球打到1点回家睡觉已经三点了。
	//	15号10点多醒了，中午和媛聊会天，她说中午吃的是韩式拌饭，那我也吃这个吧，2点半准备出门的时候，麻将机到了。需要搬到小陆家里小陆住四楼，我和家辉两个人去的。太沉了啊我草啊。
	//麻将机从一楼搬到四楼，我还一天没吃饭了，我们俩搬一层歇一层，杭州还40度我凑啊，一边搬一边暴汗，累死了我凑。搬完去韩式拌饭，店关门了。准备去下一家，距离还有3.1km，我看着我全身的暴汗。
	//果断选择了附近的猪脚饭。吃饭就回家躺着休息了。晚上5点晨和小鹿还有家辉把麻将机拼好，开始打麻将到11点。输了320，4点钟睡的觉。
	//	17号玩剧本杀有17、家辉、家辉对象。四点半开始，地方离钱塘江很近，上次去附近玩，下雨天去江边逛了一下，不美，这次想带着家辉他俩一起去江边逛一逛，路程一小时，日常提前半个小时，
	//江边逛半个小时，所以两点钟出门，到了江边，下雨了，很大，没走两步就下雨了，很难受。转头去剧本杀店里，到达时间15点40。5点钟开始玩，我叫花枝丸，我杀了毛利小四郎，因为他吃关东煮不吃丸子
	//玩到晚上11点。脑子完全已经停转了。12点半睡的觉。
	//2023年7月19日	17:54:05
	//	这个月借出去2K，随礼1K，冲自助餐厅会员0.6K，打麻将输了0.6K，后面还有还有20天才能发工资。这个月要控制花销了。
	//	朋友结婚，羡慕了，我要是有个人能结婚就好了。
}
func lpc() {
	//20230519 这里是lpc提交的数据
	//20230522 麻将，输了15，晚上顶雨去吃牛肉，回家窜了
	//20230606 昨晚吃了炸串炸鸡夜宵，今天窜了
	//20230614 没人帮我提交，这仇我先记下了
	//20230618 最近装修，好累
	//20230629 总算收拾完了，要猛猛挂咸鱼，落后了都
	//20230717 不知道干啥，好无聊
}

func horJin() {
	//这里是HorJin提交的数据
}

func yinHuiLin () {
	//2023-05-19 叫人三国杀 没人
	//2023-05-20 叫人三国杀 没人
	//2023-05-21 叫人红警 没人
	//2023-05-22 没叫
	//2023-05-23 没叫
	//2023-05-30 你狗叫什么？
	func main() {
	    for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
		    for k := 0; k < 10; k++ {
			for l := 0; l < 10; l++ {
			    for m := 0; m < 10; m++ {
				for n := 0; n < 10; n++ {
				    for o := 0; o < 10; o++ {
					for p := 0; p < 10; p++ {
					    for q := 0; q < 10; q++ {
						for r := 0; r < 10; r++ {
						    fmt.Printf("%d%d%d%d%d%d%d%d%d%d\n", i, j, k, l, m, n, o, p, q, r . sb张昊)
						}
					    }
					}
				    }
				}
			    }
			}
		    }
		}
	    }
	}
	//好几天没提交了今天来说一句 sb张昊
	<!-- <div id="app">
        <h2 v-cloak>总价格：{{totalPrice}}</h2>
    </div>
    <script src="../js/vue.js"></script>
    <script>
        const app = new Vue({
            el: '#app',
            data: {
                books: [
                    {id: 1, name: 'Unix编程艺术', price: 100},
                    {id: 2, name: '代码大全', price: 200},
                    {id: 3, name: '深入理解计算机原理', price: 300},
                    {id: 4, name: '现代操作系统', price: 400},
                ]
            },
            computed: {
                totalPrice() {
                    let result = 0
                    for (let i = 0; i < this.books.length; i++) {
                        result += this.books[i].price
                    }
                    // for (let i in this.books) {
                    //     result += this.books[i].price
                    // }
                    // for (let i of this.books) {
                    //     result += i.price
                    // }
                    return result
                }
            },
            methods: {},
        })
    </script> -->
}

func sbLiuBoTao() {
	//这里是谁提交的数据
}

func ZhangJiaHong() {
	//这里是来自艾欧尼亚的ZX程序员提交的数据
	now := time.Now().Unix()
	fmt.Println("230522提交", now)
}

func DBC529() {
	//  这里是来自DBC的ZX程序员

	//  今天也很无聊
	//  gitlab 真不是个好东西 烦死
	//  今天看到老李的骑兵连没了
	fmt.Println("230522")
}

func shuaiGeZh() {
	//  5.22老板不在摸鱼一天
	//  今天又帅了
	//  今天更帅了
	//  我张昊就是个大傻逼
}

func nobody(){
	//这是一个来自神秘人的神秘代码
	//s神秘代码被启用
	//哈哈哈龙哥日记太有节目了
	//new life
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
	fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog！ liu</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:1102389095@qq.com\">1102389095@qq.com</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}

// Article  对应一条文章数据
type Article struct {
	Title, Body string
	ID          int64
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

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "访问文章列表")
}
func articlesCreateHandler(w http.ResponseWriter, r *http.Request) {
	storeURL, _ := router.Get("articles.store").URL()
	data := ArticlesFormData{
		Title:  "",
		Body:   "",
		URL:    storeURL,
		Errors: nil,
	}
	tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

// ArticlesFormData 创建博文表单数据
type ArticlesFormData struct {
	Title, Body string
	URL         *url.URL
	Errors      map[string]string
}

// 表单验证
func validateArticleFormData(title string, body string) map[string]string {
	errors := make(map[string]string)
	// 验证标题
	if title == "" {
		errors["title"] = "标题不能为空"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "标题长度需介于 3-40"
	}

	// 验证内容
	if body == "" {
		errors["body"] = "内容不能为空"
	} else if utf8.RuneCountInString(body) < 10 {
		errors["body"] = "内容长度需大于或等于 10 个字节"
	}

	return errors
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {

	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	errors := validateArticleFormData(title, body)

	// 检查是否有错误
	// 检查是否有错误
	if len(errors) == 0 {
		lastInsertID, err := saveArticleToDB(title, body)
		if lastInsertID > 0 {
			fmt.Fprint(w, "插入成功，ID 为"+strconv.FormatInt(lastInsertID, 10))
		} else {
			checkError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {

		storeURL, _ := router.Get("articles.store").URL()

		data := ArticlesFormData{
			Title:  title,
			Body:   body,
			URL:    storeURL,
			Errors: errors,
		}
		tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	}
}

func saveArticleToDB(title string, body string) (int64, error) {

	// 变量初始化
	var (
		id   int64
		err  error
		rs   sql.Result
		stmt *sql.Stmt
	)

	// 1. 获取一个 prepare 声明语句
	stmt, err = db.Prepare("INSERT INTO articles (title, body) VALUES(?,?)")
	// 例行的错误检测
	if err != nil {
		return 0, err
	}

	// 2. 在此函数运行结束后关闭此语句，防止占用 SQL 连接
	defer stmt.Close()

	// 3. 执行请求，传参进入绑定的内容
	rs, err = stmt.Exec(title, body)
	if err != nil {
		return 0, err
	}

	// 4. 插入成功的话，会返回自增 ID
	if id, err = rs.LastInsertId(); id > 0 {
		return id, nil
	}

	return 0, err
}

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 设置标头
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 2. 继续处理请求
		next.ServeHTTP(w, r)
	})
}
func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 除首页以外，移除所有请求路径后面的斜杆
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}

		// 2. 将请求传递下去
		next.ServeHTTP(w, r)
	})
}

func createTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
    id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    body longtext COLLATE utf8mb4_unicode_ci
); `

	_, err := db.Exec(createArticlesSQL)
	checkError(err)
}

func getRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}

func getArticleByID(id string) (Article, error) {
	article := Article{}
	query := "SELECT * FROM articles WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Body)
	return article, err
}

func main() {
	initDB()
	createTables()
	router.HandleFunc("/", homeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")

	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/create", articlesCreateHandler).Methods("GET").Name("articles.create")
	router.HandleFunc("/articles/{id:[0-9]+}/edit", articlesEditHandler).Methods("GET").Name("articles.edit")
	router.HandleFunc("/articles/{id:[0-9]+}", articlesUpdateHandler).Methods("POST").Name("articles.update")

	// 自定义 404 页面
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// 中间件：强制内容类型为 HTML
	router.Use(forceHTMLMiddleware)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
