package main

import (
	"fmt"
	"time"
)

func findAddress(addresses []string) {
	var tx int64 = 0 // inisialisasi tx dengan nilai 0
	ty := time.Now().UnixNano()
	fmt.Println("The Default Address is", addresses[0])
	fmt.Println("The Performance is", (ty-tx)/1000000, "ms")
}

func main() {
	address := "DKI Jakarta"
	addresses := make([]string, 10)
	for i := range addresses {
		addresses[i] = address
	}
	tx := time.Now().UnixNano()
	findAddress(addresses)
	ty := time.Now().UnixNano()
	fmt.Println("The Performance is", (ty-tx)/1000000, "ms") // menghitung waktu akhir di dalam fungsi main()
}