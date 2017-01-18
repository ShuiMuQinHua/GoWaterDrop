package main

import(
    "fmt"
)
//素数:除了一和它自身以外，没有其它的因数
func generate(ch chan int) {
    for i := 2;i<=1000 ; i++ {
        ch <- i // Send 'i' to channel 'ch'.
    }
    defer close(ch)
}

func filter(in, out chan int, prime int) {
    for {
        i := <-in // Receive value of new variable 'i' from 'in'.
        if i%prime != 0 {
            out <- i // Send 'i' to channel 'out'.
        }
    }
}

func main() {
    ch := make(chan int) // Create a new channel.
    go generate(ch) // Start generate() as a goroutine.
    for {
        prime,open := <-ch
        if !open{
            break
        }
        fmt.Print(prime, " ")
        ch1 := make(chan int)
        go filter(ch, ch1, prime)
        ch = ch1
    }
}