package main

import (
	"fmt"
	"strings"
)

func main() {
	//string类型
	//变量声明,没有初始化，初始值为空字符串，占两个字节
	var str string
	fmt.Println("str = ---", str, "---")
	str = "abc" //字符串用双引号
	fmt.Println("str = ", str)
	/*
		自动推导
	*/
	var str2 = "abc"
	fmt.Println("str2 = ", str2)
	str3 := "abcd"
	fmt.Printf("str3 type is %T str3 = %s\n", str3, str3)
	//输出字符串有几个字符
	fmt.Println("len(str3) = ", len(str3))

	/*
		字符串和字符区别
		字符：
		1.单引号
		2.字符元素往往只有一个字符，转义字符除外'\n'
		字符串：
		1.双引号
		2.字符串有一个或多个字符组成
		3.字符串都是隐藏了一个结束符‘\0’
	*/

	//输出字符串中指定字符
	fmt.Printf("str3[0] = %c , str3[1] = %c \n", str3[0], str3[1])

	//使用反引号“ ` ”可以保持文本格式不变，转义字符不会被转义，会原封不动的打印出来
	str4 := `字符串和字符区别
		字符 \n \t\r
		单引号
		字符元素往往只有一个字符，转义字符除外
		字符串：
		1.双引号
		2.字符串有一个或多个字符组成
		3.字符串都是隐藏了一个结束符‘\0’`
	fmt.Println(str4)

	str5 := "hello go"
	fmt.Println(reverse(str5))
	fmt.Println(reverse1(str5))

	/*判断字符串开头结尾*/
	fmt.Println(urlProcess("www.baidu.com"))
	fmt.Println(pathProcess("D:"))

}

func reverse(str string) string {
	var result string
	strlen := len(str)
	for i := 0; i < strlen; i++ {
		result = result + fmt.Sprintf("%c", str[strlen-i-1])
	}
	return result
}

func reverse1(str string) string {
	var result []byte
	tmp := []byte(str)
	length := len(str)
	for i := 0; i < length; i++ {
		result = append(result, tmp[length-i-1])
	}
	return string(result)
}

//判断是否有http://头
func urlProcess(url string) string {
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)
	}
	return url
}

//判断是否以斜杠/结尾
func pathProcess(path string) string {
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}
	return path
}
