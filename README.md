## Golang 代码仓储
记录 Go 的学习历程

TBD

## 标准库
常用的标准库的使用，参考 API：https://golang.google.cn/pkg/

### fmt
TBD

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

### io
TBD

### os
TBD

### net
TBD

### bufio
TBD

### crypto
TBD

### log
TBD