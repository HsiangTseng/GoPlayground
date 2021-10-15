package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	writePHPFile("Member", "WEWQEQ")
}

func writePHPFile(file_name, file_path string) {
	// 讀模板檔，存在content
	content, read_err := ioutil.ReadFile("controllerTemplate.txt")
	if read_err != nil {
		log.Fatal(read_err)
	}

	// 在模板中找出要放入檔名的index，將模板改好再寫檔
	file_name_byte := []byte(file_name)
	file_name_index := strings.Index(string(content), " Controller") + 1                                                  // 取得要插入的位子
	content_after_named := string(content)[:file_name_index] + string(file_name_byte) + string(content)[file_name_index:] //在要插入檔名的位子插入檔名

	fmt.Println(content_after_named)

	os.WriteFile(string(file_name_byte)+"Controller.php", []byte(content_after_named), 0755) //寫檔
}
