package main

import (
	"archive/zip"
	"bufio"
	"compress/gzip"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"io/ioutil"
	"net/http"
	"time"
	//	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//	newFile()
	//truncateFile()
	//getFileInfo()
	//	renameAndMoveFile()
	//	delFile()
	//	changeMode()
	//	copyFile()
	//	seekFile()
	//	writeFile()
	//	appendToFile()
	//	quickWriteFile()
	//	bufioWriteFile()
	//	readFileMaxN()
	//	readFileOnlyN()
	//	readFileMinN()
	//	readFileAll()
	//	bufioRead()
	//	scanRead()
	//	zipFile()
	//	uzipFile()
	//	downHttpFile()
	hashFile()
	//write_file()
	//read_file()
	//getFileList()

	//	flag.Parse()
	//	root := flag.Arg(0)
	//	//接收命令行参数  返回这个目录下面的所有的文件
	//	getFileList(root)
}

//创建一个空文件
func newFile() {
	newfile, err := os.Create("newfile.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("文件创建成功")
	}
	newfile.Close()
}

//剪裁文件
func truncateFile() {
	// 裁剪一个文件到10个字节。
	// 如果文件本来就少于10个字节，则文件中原始内容得以保留，剩余的字节以null字节填充。
	// 如果文件本来超过10个字节，则超过的字节会被抛弃。
	// 这样我们总是得到精确的10个字节的文件。
	// 传入0则会清空文件。
	err := os.Truncate("newfile.txt", 10)
	if err != nil {
		log.Fatal(err)
	}
}

//获取文件信息
func getFileInfo() {
	fileinfo, err := os.Stat("newfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file name:", fileinfo.Name())
	fmt.Println("size in bytes:", fileinfo.Size())
	fmt.Println("permissions:", fileinfo.Mode())
	fmt.Println("last modify:", fileinfo.ModTime())
	fmt.Println("is directory:", fileinfo.IsDir())
	fmt.Printf("system interface type: %T\n", fileinfo.Sys())
	fmt.Printf("system info: %+v\n\n", fileinfo.Sys())
}

//重命名和移动
func renameAndMoveFile() {
	originalpath := "newfile.txt"
	newpath := "newfile2.txt"
	err := os.Rename(originalpath, newpath)
	if err != nil {
		log.Fatal(err)
	}
}

//删除文件
func delFile() {
	err := os.Remove("newfile2.txt")
	if err != nil {
		log.Fatal(err)
	}
}

//打开和关闭文件
func openAndCloseFile() {
	//简单的以只读的方式打开
	file, _ := os.Open("test.txt")
	file.Close()

	// OpenFile提供更多的选项。
	// 最后一个参数是权限模式permission mode
	// 第二个是打开时的属性
	file, _ = os.OpenFile("test.txt", os.O_APPEND, 0666)
	file.Close()

	// 下面的属性可以单独使用，也可以组合使用。
	// 组合使用时可以使用 OR 操作设置 OpenFile的第二个参数，例如：
	// os.O_CREATE|os.O_APPEND
	// 或者 os.O_CREATE|os.O_TRUNC|os.O_WRONLY

	// os.O_RDONLY // 只读
	// os.O_WRONLY // 只写
	// os.O_RDWR // 读写
	// os.O_APPEND // 往文件中添建（Append）
	// os.O_CREATE // 如果文件不存在则先创建
	// os.O_TRUNC // 文件打开时裁剪文件
	// os.O_EXCL // 和O_CREATE一起使用，文件不能存在
	// os.O_SYNC // 以同步I/O的方式打开
}

//判断文件是否存在
func isFileExist() {
	_, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("file dose not exist")
		}
	}
	fmt.Println("file exist")
}

//检查读写权限
func fileMod() {
	// 这个例子测试写权限，如果没有写权限则返回error。
	// 注意文件不存在也会返回error，需要检查error的信息来获取到底是哪个错误导致。
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("error:Write permission denied.")
		}
	}
	file.Close()

	//测试读权限
	file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("error:Read permission denied.")
		}
	}
	file.Close()
}

//改变权限、拥有者、时间戳
func changeMode() {
	//使用linux风格改变文件权限
	err := os.Chmod("test.txt", 0777)
	if err != nil {
		log.Println(err)
	}

	//改变文件所有者 windows环境不支持
	//	err = os.Chown("test.txt", os.Getuid(), os.Getgid())
	//	if err != nil {
	//		log.Println(err)
	//	}

	//改变时间戳
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}
}

//复制文件
func copyFile() {
	// 打开原始文件
	originalFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	// 创建新的文件作为目标文件
	newFile, err := os.Create("test_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// 从源中复制字节到目标文件
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	// 将文件内容flush到硬盘中
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}

//跳转到文件指定位置
func seekFile() {
	file, _ := os.Open("test.txt")
	defer file.Close()

	// 偏离位置，可以是正数也可以是负数
	var offset int64 = 7

	// 用来计算offset的初始位置
	// 0 = 文件开始位置
	// 1 = 当前位置
	// 2 = 文件结尾处
	var whence int = 0
	newPosition, err := file.Seek(offset, whence)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just moved to 7:", newPosition)

	// 从当前位置回退两个字节
	newPosition, err = file.Seek(-2, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just moved back two:", newPosition)

	// 使用下面的技巧得到当前的位置
	currentPosition, err := file.Seek(0, 1)
	fmt.Println("Current position:", currentPosition)

	// 转到文件开始处
	newPosition, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Position after seeking 0,0:", newPosition)
}

func writeFile() {
	// 可写方式打开文件
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 写字节到文件中
	byteSlice := []byte("Bytes!\n")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}

//在文件末尾追加内容
func appendToFile() {
	// 以只写的模式，打开文件
	f, err := os.OpenFile("test.txt", os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte("1235jhghjkk23424"), n)
	}
	defer f.Close()
}

//快写文件
func quickWriteFile() {
	//ioutil.WriteFile可以处理创建／打开文件、写字节slice和关闭文件一系列的操作
	err := ioutil.WriteFile("test.txt", []byte("Hi\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}

	// 快读文件文件到byte slice中 内存中
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Data read: %s\n", data)
}

func bufioWriteFile() {
	// 打开文件，只写
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 为这个文件创建buffered writer
	bufferedWriter := bufio.NewWriter(file)

	// 写字节到buffer
	bytesWritten, err := bufferedWriter.Write([]byte{65, 66, 67})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)

	// 写字符串到buffer
	// 也可以使用 WriteRune() 和 WriteByte()
	bytesWritten, err = bufferedWriter.WriteString("Buffered string\n")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)

	// 检查缓存中的字节数
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	// 还有多少字节可用（未使用的缓存大小）
	bytesAvailable := bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// 写内存buffer到硬盘
	bufferedWriter.Flush()

	// 丢弃还没有flush的缓存的内容，清除错误并把它的输出传给参数中的writer
	// 当你想将缓存传给另外一个writer时有用
	bufferedWriter.Reset(bufferedWriter)

	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// 重新设置缓存的大小。
	// 第一个参数是缓存应该输出到哪里，这个例子中我们使用相同的writer。
	// 如果我们设置的新的大小小于第一个参数writer的缓存大小， 比如10，我们不会得到一个10字节大小的缓存，
	// 而是writer的原始大小的缓存，默认是4096。
	// 它的功能主要还是为了扩容。
	bufferedWriter = bufio.NewWriterSize(bufferedWriter, 8000)

	// resize后检查缓存的大小
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)
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

//读取最多n个字节
func readFileMaxN() {
	// 打开文件，只读
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 从文件中读取len(b)字节的文件。
	// 返回0字节意味着读取到文件尾了
	// 读取到文件会返回io.EOF的error
	byteSlice := make([]byte, 16)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", bytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}

//读取正好n个字节
func readFileOnlyN() {
	// Open file for reading
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// file.Read()可以读取一个小文件到大的byte slice中，
	// 但是io.ReadFull()在文件的字节数小于byte slice字节数的时候会返回错误
	byteSlice := make([]byte, 2)
	numBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}

//读取文件最少n个字节
func readFileMinN() {
	// 打开文件，只读
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	byteSlice := make([]byte, 512)
	minBytes := 8
	// io.ReadAtLeast()在不能得到最小的字节的时候会返回错误，但会把已读的文件保留
	numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}

//读取所以的字节
func readFileAll() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// os.File.Read(), io.ReadFull() 和
	// io.ReadAtLeast() 在读取之前都需要一个固定大小的byte slice。
	// 但ioutil.ReadAll()会读取reader(这个例子中是file)的每一个字节，然后把字节slice返回。
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as hex: %x\n", data)
	fmt.Printf("Data as string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))
}

//缓存读取
func bufioRead() {
	// 打开文件，创建buffered reader
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	bufferedReader := bufio.NewReader(file)

	// 得到字节，当前指针不变
	byteSlice := make([]byte, 5)
	byteSlice, err = bufferedReader.Peek(5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)

	// 读取，指针同时移动
	numBytesRead, err := bufferedReader.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSlice)

	// 读取一个字节, 如果读取不成功会返回Error
	myByte, err := bufferedReader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read 1 byte: %c\n", myByte)

	// 读取到分隔符，包含分隔符，返回byte slice
	dataBytes, err := bufferedReader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read bytes: %s\n", dataBytes)

	// 读取到分隔符，包含分隔符，返回字符串
	dataString, err := bufferedReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read string: %s\n", dataString)

	//这个例子读取了很多行，所以test.txt应该包含多行文本才不至于出错
}

//按指定的符号读取数据
func scanRead() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	// 缺省的分隔函数是bufio.ScanLines,我们这里使用ScanWords。
	// 也可以定制一个SplitFunc类型的分隔函数
	scanner.Split(bufio.ScanWords)

	// scan下一个token.
	success := scanner.Scan()
	if success == false {
		// 出现错误或者EOF是返回Error
		err = scanner.Err()
		if err == nil {
			log.Println("Scan completed and reached EOF")
		} else {
			log.Fatal(err)
		}
	}

	// 得到数据，Bytes() 或者 Text()
	fmt.Println("First word found:", scanner.Text())

	// 再次调用scanner.Scan()发现下一个token
}

//压缩文件
func zipFile() {
	// 创建一个打包文件
	outFile, err := os.Create("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// 创建zip writer
	zipWriter := zip.NewWriter(outFile)

	// 往打包文件中写文件。
	// 这里我们使用硬编码的内容，你可以遍历一个文件夹，把文件夹下的文件以及它们的内容写入到这个打包文件中。
	var filesToArchive = []struct {
		Name, Body string
	}{
		{"test.txt", "String contents of file"},
		{"test2.txt", "\x61\x62\x63\n"},
	}

	// 下面将要打包的内容写入到打包文件中，依次写入。
	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = fileWriter.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// 清理
	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
}

//解压缩文件
func uzipFile() {
	// 打开一个gzip文件。
	// 文件是一个reader,但是我们可以使用各种数据源，比如web服务器返回的gzipped内容，
	// 它的内容不是一个文件，而是一个内存流  gzipFile
	gzipFile, err := os.Open("test1.zip")
	if err != nil {
		log.Fatal(err)
	}

	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipReader.Close()

	// 解压缩到一个writer,它是一个file writer
	outfileWriter, err := os.Create("unzipped.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outfileWriter.Close()

	// 复制内容
	_, err = io.Copy(outfileWriter, gzipReader)
	if err != nil {
		log.Fatal(err)
	}
}

//通过http下载文件
func downHttpFile() {
	newFile, err := os.Create("devdungeon.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	url := "http://www.devdungeon.com/archive"
	response, err := http.Get(url)
	defer response.Body.Close()

	// 将HTTP response Body中的内容写入到文件
	// Body满足reader接口，因此我们可以使用ioutil.Copy
	numBytesWritten, err := io.Copy(newFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
}

//hash和摘要
func hashFile() {
	// 得到文件内容
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// 计算Hash
	fmt.Printf("Md5: %x\n\n", md5.Sum(data))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(data))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(data))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(data))
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
