## Golang 代码仓储
记录 Go 的学习历程

TBD

## 标准库
常用的标准库的使用，参考 API：https://golang.google.cn/pkg/

以下库仅作简短的作用说明，具体都使用方法参考具体的代码

### fmt
基础使用的 Print 系列就不多说了，最常用的三个（Print、Println、Printf）

#### Fprint
Fprint 系列函数会将内容输出到一个 `io.Writer` 接口类型的变量w中，我们通常用这个函数往文件中写入内容。

注意，只要满足 `io.Writer` 接口的类型都支持写入。

同样，它也有三个函数，Fprint、Fprintf、Fprintln

#### Sprint
Sprint 系列函数会把传入的数据生成并返回一个字符串。

注意是返回字符串，而不是打印！

同样，它也是包括三个函数，同上。

#### Errorf
Errorf 函数根据 format 参数生成格式化字符串并返回一个包含该字符串的错误（类型）。

通常使用这种方式来自定义错误类型;
Go1.13 版本为 `fmt.Errorf` 函数新加了一个 `%w` 占位符用来生成一个可以包裹 Error 类型的 Wrapping Error。

#### 常用占位符
通用型：

| 占位符 | 说明                               |
| :----: | :--------------------------------- |
|   %v   | 值的默认格式表示                   |
|  %+v   | 类似%v，但输出结构体时会添加字段名 |
|  %#v   | 值的Go语法表示                     |
|   %T   | 打印值的类型                       |
|   %%   | 百分号                             |

基本数据类型：

| 占位符 | 说明                                                         |
| :----: | :----------------------------------------------------------- |
|   %b   | 表示为二进制                                                 |
|   %c   | 该值对应的unicode码值                                        |
|   %d   | 表示为十进制                                                 |
|   %o   | 表示为八进制                                                 |
|   %x   | 表示为十六进制，使用a-f                                      |
|   %X   | 表示为十六进制，使用A-F                                      |
|   %U   | 表示为Unicode格式：U+1234，等价于”U+%04X”                    |
|   %q   | 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示 |
|   %t   | true或false                                                  |
|   %s   | 直接输出字符串或者[]byte                                     |
|   %q   | 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示 |
|   %x   | 每个字节用两字符十六进制数表示（使用a-f）                    |
|   %X   | 每个字节用两字符十六进制数表示（使用A-F）                    |
|   %p   | 指针，表示为十六进制，并加上前导的 0x                        |

宽度标识：

| 占位符 | 说明               |
| :----: | :----------------- |
|   %f   | 默认宽度，默认精度 |
|  %9f   | 宽度9，默认精度    |
|  %.2f  | 默认宽度，精度2    |
| %9.2f  | 宽度9，精度2       |
|  %9.f  | 宽度9，精度0       |

其他：

| 占位符 | 说明                                                         |
| :----: | :----------------------------------------------------------- |
|  ’+’   | 总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）； |
|  ’ ‘   | 对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格 |
|  ’-’   | 在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）； |
|  ’#’   | 八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的go字面值； |
|  ‘0’   | 使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面； |

#### 获取输入
Go 语言 fmt 包下有 `fmt.Scan`、`fmt.Scanf`、`fmt.Scanln` 三个函数，可以在程序运行过程中从标准输入获取用户的输入。

Scan 从标准输入扫描文本，读取**由空白符分隔**的值保存到传递给本函数的参数中（将以空白符分隔的数据分别存入指定的参数），**换行符视为空白符**。
本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因。

示例：

``` go
func main() {
  var (
    name    string
    age     int
    married bool
  )
  fmt.Scan(&name, &age, &married)
  fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
```

一般情况，ln 用的比较多；

#### 其他
Fscan 系列：与 Scan 系列基本一致，只不过它们不是从标准输入中读取数据而是从 `io.Reader` 中读取数据。

Sscan 系列：它们不是从标准输入中读取数据而是从指定字符串中读取数据。

### strconv

它主要实现对字符串和基本数据类型之间的转换，常用 API：

``` go
func Atoi(s string) (int, error)
func CanBackquote(s string) bool
func FormatBool(b bool) string
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
func FormatInt(i int64, base int) string
func FormatUint(i uint64, base int) string
func IsGraphic(r rune) bool
func IsPrint(r rune) bool
func Itoa(i int) string
func ParseBool(str string) (bool, error)
func ParseFloat(s string, bitSize int) (float64, error)
func ParseInt(s string, base int, bitSize int) (i int64, err error)
func ParseUint(s string, base int, bitSize int) (uint64, error)
```

相比直接使用 `string()` 之类，它更安全，例如 Itoa 可以将整型数字转换为字符串的数字，而如果使用 string() 的话，就会被当作 ASCII，然后转成字符。

### sort
TBD

### time
TBD

### io和filepath

简单都文件可以尝试使用 ioutil 来进行读取；它默认读取都是整个文件；

使用 `io.copy` 可以实现文件拷贝；

与之相关联的处理路径的包是 filepath；例如自动的路径分隔符（or 自动替换），获取扩展名，路径都判断（相对、绝对），路径的模式匹配等。

TBD

### os

与 I/O 的联系挺密切，我也主要关注 IO 方面，例如 open（默认只读方式打开）、close、OpenFile 方法；以此获得文件句柄对象；

| 模式          | 含义     |
| :------------ | :------- |
| `os.O_WRONLY` | 只写     |
| `os.O_CREATE` | 创建文件 |
| `os.O_RDONLY` | 只读     |
| `os.O_RDWR`   | 读写     |
| `os.O_TRUNC`  | 清空     |
| `os.O_APPEND` | 追加     |

`perm`：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01，跟 linux 类似。

TBD

### net
TBD

### bufio
有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用 bufio 包来实现；

bufio 是在 file 的基础上封装了一层 API，支持更多的功能；

TBD

### crypto
TBD

### log
TBD

## 常用库

以下库的具体使用的代码不贴了，都在 uselib 中。

### MySQL
基本的驱动： `github.com/go-sql-driver/mysql`

新版本推荐使用 `go mod` 进行安装（download、tidy）

---

对于数据库，Go 也提供了完整的一套接口，就是 `database/sql`，但是没有具体实现；类似 JDBC 的 SPI 机制，具体的驱动实现需要你 get 相应的数据库依赖，虽然没有驱动不妨碍你程序的编写，但是跑不起来。

SQL 包中的 DB 结构体是一个数据库（操作）句柄，代表一个具有零到多个底层连接的**连接池**（可以理解为自带连接池定义）。它可以安全的被多个 go 程同时使用，是操作数据库的主要方式。

sql 包会自动创建和释放连接；它也会维护一个闲置连接的连接池。如果数据库具有单连接状态的概念，该状态只有在事务中被观察时才可信。一旦调用了 BD.Begin，返回的 Tx 会绑定到单个连接。当调用事务 Tx 的 Commit 或 Rollback 后，该事务使用的连接会归还到 DB 的闲置连接池中。连接池的大小可以用 SetMaxIdleConns 方法控制。

> 一般而言，程序需要在退出时通过`sql.DB`的`Close()`方法释放数据库连接资源。如果其生命周期不超过函数的范围，则应当使用`defer db.Close()`
>
> `sql.DB`对象是为了长连接而设计的，不要频繁`Open()`和`Close()`数据库。而应该为每个待访问的数据库创建**一个**`sql.DB`实例，并在用完前一直保留它。需要时可将其作为参数传递，或注册为全局对象。
>
> http://vonng.com/blog/go-database-tutorial/

---

预处理：

1. 把 SQL 语句分成两部分，命令部分与数据部分。
2. 先把命令部分发送给 MySQL 服务端，MySQL 服务端进行 SQL 预处理（编译）。
3. 然后把数据部分发送给 MySQL 服务端，MySQL 服务端对 SQL 语句进行占位符替换。
4. MySQL 服务端执行完整的 SQL 语句并将结果返回给客户端。

好处有两点：优化 MySQL 服务器重复执行 SQL 的效率，提前让服务器编译，一次编译多次执行；还有就是防止 SQL 注入。

Go 中的 `Prepare` 方法会先将 sql 语句发送给 MySQL 服务端，返回一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。

| 数据库     | 占位符语法   |
| :--------- | :----------- |
| MySQL      | `?`          |
| PostgreSQL | `$1`, `$2`等 |
| SQLite     | `?` 和`$1`   |
| Oracle     | `:name`      |

### sqlx

使用 sqlx 能简化 DB 的操作，相当于是 `database/sql` 的一层封装，也需要引入数据库驱动。

> panic 与 recover 是 Go 的两个内置函数，这两个内置函数用于处理 Go 运行时的错误。
>
> panic 用于主动抛出错误, recover 用来捕获 panic 抛出的错误。