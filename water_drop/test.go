package main

import {
	"io/ioutil"
	"net"
	"os/exec"
	"strings"
	"sync"
	"time"
}


func main() {

	getHostName()

	//	a = "G"
	//	fmt.Println(a)
	//	f1()
}

// func f1() {
// 	a := "O"
// 	fmt.Println(a)
// 	f2()
// }

// func f2() {
// 	fmt.Println(a) //输出G O G
// }


func getHostName() string {

	cmd := exec.Command("/bin/bash", "-c", `uname -a`)
	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		return ""
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return ""
	}

	if err := cmd.Wait(); err != nil {
		return ""
	}

	output := string(bytes[:])
	hostname := strings.Split(output, " ")[1]
	return hostname
}

