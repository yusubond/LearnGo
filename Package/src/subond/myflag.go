/*
* flag 实现命令行解析
* Code: https://golang.org/pkg/flag/
 */
package subond

import (
	"flag"
	"fmt"
)

func Hello() {
	// flag.String() 指定参数名称，默认值，用例
	username := flag.String("name", "", "-name subond")

	b := flag.Bool("key", true, "-key true")

	// flag.Parse()  解析参数
	flag.Parse()
	fmt.Println("Hello, ", *username)

	// flag.Args() 返回非命令行参数
	args := flag.Args()
	fmt.Println(args)
	fmt.Println(len(args))

	// flag.Bool() 返回默认的布尔值参数
	fmt.Println("key is ", *b)
}
