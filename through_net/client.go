package main
 
import (
    "net"
    "fmt"
    "flag"
    "os"
)
 
func handler(r net.Conn, localPort int) {
    buf := make([]byte, 1024);
    for {
        //先从远程读数据
        n, err := r.Read(buf);
        if err != nil {
            continue;
        }
        data := buf[:n];
        //建立与本地80服务的连接
        local, err := net.Dial("tcp", fmt.Sprintf(":%d", localPort));
        if err != nil {
            continue;
        }
        //向80服务写数据
        n, err = local.Write(data);
        if err != nil {
            continue;
        }
        //读取80服务返回的数据
        n, err = local.Read(buf);
        //关闭80服务，因为本地80服务是http服务，不是持久连接
        //一个请求结束，就会自动断开。所以在for循环里我们要不断Dial，然后关闭。
        local.Close();
        if err != nil {
            continue;
        }
        data = buf[:n];
        //向远程写数据
        n, err = r.Write(data);
        if err != nil {
            continue;
        }
    }
}
 
func main() {
    //参数解析
    host := flag.String("host", "127.0.0.1", "服务器地址");
    remotePort := flag.Int("remotePort", 8888, "服务器端口");
    localPort := flag.Int("localPort", 80, "本地端口");
    flag.Parse();
    if flag.NFlag() != 3 {
        flag.PrintDefaults();
        os.Exit(1);
    }
    //建立与服务器的连接
    remote, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *remotePort));
    if err != nil {
        fmt.Println(err);
    }
    go handler(remote, *localPort);
 
    select {};
}