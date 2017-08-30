package ziptest

import (
	"os"
	"testing"
)

//测试压缩文件
func TestCompress(t *testing.T) {
	f1, err := os.Open("test1.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f1.Close()
	f2, err := os.Open("test2.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f2.Close()
	f3, err := os.Open("dirtest")
	if err != nil {
		t.Fatal(err)
	}
	defer f3.Close()
	var files = []*os.File{f1, f2, f3}
	dest := "test.zip"
	err = Compress(files, dest)
	if err != nil {
		t.Fatal(err)
	}
}

//测试解压文件
func TestDeCompress(t *testing.T) {
	err := DeCompress("test.zip", "dezip")
	if err != nil {
		t.Fatal(err)
	}
}
