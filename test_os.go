package main
import(
    "fmt"
    "os"
)

func main(){
    cmd:=os.Args[1:]
    fmt.Println(cmd);
}