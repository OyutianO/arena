package util

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)


func Get_linux_command(command string) string {
	cmd := exec.Command("/bin/bash", "-c", command)

	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
	}

	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
	}

	return string(bytes)
}