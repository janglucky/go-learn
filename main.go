package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/**
  from: http://www.isharey.com/?p=143
*/

func main() {
	var filename = "./output1.txt"
	var f *os.File
	var err1 error
	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	check(err1)

	i := 0
	for	i <= 10000 {
		ip := fmt.Sprintf("%d.%d.%d.%d\n",rand.Intn(255),rand.Intn(255),rand.Intn(255),rand.Intn(255))
		io.WriteString(f, ip) //写入文件(字符串)
		i++
	}
	//
	//fmt.Println(i)
	check(err1)



	f.Close()
}