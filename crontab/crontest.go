package main
  
 import (
      "fmt"
      "github.com/robfig/cron"
  )
  
  func main() {
	fmt.Println("callYourFunc come here.")
    spec := "c" // 每天凌晨2：30
     c := cron.New()
     c.AddFunc(spec, callYourFunc)
     c.Start()
     select {}
 }
 
 func callYourFunc() {
     fmt.Println("callYourFunc come here.")
 }
 
 // 每天凌晨2：30就会调用一次callYourFunc函数