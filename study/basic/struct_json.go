package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// 注意大小写代表权限
type User struct {
	Name string
	Age  int
	// 使用 tag，在序列化的时候使用 tag 名
	// 可以在运行时通过反射读出来，一定要注意格式
	// omitempty 表示为空（零值）时不进行序列化
	Desc string `json:"desc,omitempty"`
}

func main() {
	users := []User{
		{"mps", 12, "desc"},
		{Name: "skye", Age: 23},
	}
	fmt.Println(users)

	// 序列化 json (编组)
	marshal, err := json.Marshal(users)
	if err != nil {
		log.Fatal("json marshal failed: ", err)
		return
	}
	fmt.Printf("%s\n", marshal)

	// 设置前缀和缩进
	data, _ := json.MarshalIndent(users, "", "    ")
	fmt.Println(string(data))

	// 反序列化，多余字段会被忽略
	str := `[{"Name":"mps","Age":12,"desc":"desc"},{"Name":"skye","Age":23,"desc":"empty"}]`
	u1 := &[]User{}
	err = json.Unmarshal([]byte(str), u1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", u1)
}

// web 中用的流式解码器
func webJson(url string) {
	// 为了方便，忽略 err 处理了
	resp, _ := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return
	}
	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		resp.Body.Close()
		return
	}
	resp.Body.Close()
	fmt.Println(user)
}
