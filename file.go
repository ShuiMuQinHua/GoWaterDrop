package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	//write_file()
	//read_file()
	//getFileList()
	flag.Parse()
	root := flag.Arg(0)
	getFileList(root)
}

//写文件
func write_file() {
	//userfile := "test.txt"
	userfile := "test1.txt"

	//如果文件不存在  会自己创建文件
	fout, err := os.Create(userfile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userfile, err)
		return
	}
	for i := 0; i < 10; i++ {
		fout.WriteString("just a test!\r\n")
		fout.Write([]byte("just a test!\r\n"))
	}
}

//读文件
func read_file() {
	userfile := "test.txt"
	fin, err := os.Open(userfile)
	defer fin.Close()
	if err != nil {
		fmt.Println(userfile, err)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, _ := fin.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

//从a中读出数据，写入ouput中
func readfile_a_write_to_fileb() {
	fi, err := os.Open("a.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fo, err := os.Create("ouput.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		if n2, err := fo.Write(buf[:n]); err != nil {
			panic(err)
		} else if n2 != n {
			panic("error in writing")
		}
	}

}

//删除文件
func del_file() {
	os.Remove("a.txt")
}

//遍历某一个路径下面的所有的文件
func getFileList(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

}
