package main

import (
	"fmt"
)

func main() {
	fmt.Println("run tests with the following command: go test -cpuprofile=assets/min_heap_cpu.prof -memprofile=assets/min_heap_mem.prof -bench .")
	fmt.Println("run fuzz tests with the following command: go test -v -fuzz=<fuzzTestName>")
	fmt.Println("go tool pprof -png assets/min_heap_cpu.prof > assets/min_heap_cpu.png")
	fmt.Println("go tool pprof -png assets/min_heap_mem.prof > assets/min_heap_mem.png")
	fmt.Println("go tool pprof -http=:8084 assets/min_heap_mem.prof")
	fmt.Println("go tool pprof -http=:8084 assets/min_heap_cpu.prof")
	fmt.Println("To compare reports: go tool pprof -base=assets/cpu_old.prof assets/cpu_new.prof then type 'top' or 'web' to get more info")
}
