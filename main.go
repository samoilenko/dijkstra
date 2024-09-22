package main

import (
	"fmt"
)

func main() {
	fmt.Println("run tests with the following command: go test -cpuprofile=assets/recursion_cpu.prof -memprofile=assets/recursion_mem.prof -bench .")
	fmt.Println("run fuzz tests with the following command: go test -v -fuzz=<fuzzTestName>")
	fmt.Println("go tool pprof -png assets/recursion_cpu.prof > assets/recursion_cpu.png")
	fmt.Println("go tool pprof -png assets/recursion_mem.prof > assets/recursion_mem.png")
	fmt.Println("go tool pprof -http=:8084 assets/recursion_mem.prof")
	fmt.Println("go tool pprof -http=:8084 assets/recursion_cpu.prof")
	fmt.Println("To compare reports: go tool pprof -base=assets/cpu_old.prof assets/cpu_new.prof then type 'top' or 'web' to get more info")
}
