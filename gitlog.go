package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//输出git日志修改过的文件列表并去重
//例:go run .\gitlog.go tyy_api --before=2022-9-1 --after=2022-7-25 --author=nielixin
func main() {
	args := os.Args
	dir := fmt.Sprintf("D:\\nlx\\phpproject\\%s", args[1])
	param := []string{"log", "--name-only", "--oneline", "--format="}
	param = append(param, args[2:]...)
	cmd := exec.Command("git", param...)
	cmd.Dir = dir
	result, err := cmd.Output()
	if err != nil {
		fmt.Printf("命令执行失败%s\n", err.Error())
		os.Exit(1)
	}
	strSlice := strings.Split(string(result), "\n")
	temp := map[string]int{}
	for _,item := range strSlice {
		if _,ok := temp[item]; !ok {
			temp[item] = 1
			fmt.Println(item)
		}
	}
	fmt.Println(temp)
}