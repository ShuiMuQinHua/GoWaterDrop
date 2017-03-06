package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()
	zone, _ := now.UTC().Zone()
	fmt.Printf("UTC time is %d-%d-%d %02d:%02d:%02d %s\n", year, mon, day, hour, min, sec, zone)
	//	UTC time is 2017-2-10 07:29:36 UTC

	year, mon, day = now.Date()
	hour, min, sec = now.Clock()
	zone, _ = now.Zone()
	fmt.Printf("local time is %d-%d-%d %02d:%02d:%02d %s\n", year, mon, day, hour, min, sec, zone)
	//	local time is 2017-2-10 15:31:07 CST

	now = time.Now()
	time.Sleep(3 * time.Second)
	end_time := time.Now()

	var dur_time time.Duration = end_time.Sub(now)
	var elapsed_min float64 = dur_time.Minutes()
	var elapsed_sec float64 = dur_time.Seconds()
	var elapsed_nano int64 = dur_time.Nanoseconds()
	fmt.Printf("elasped %f minutes or \nelapsed %f seconds or \nelapsed %d nanoseconds\n", elapsed_min, elapsed_sec, elapsed_nano)

}
