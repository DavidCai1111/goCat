package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("需要一个路径参数")
	} else {
		path := os.Args[1]
		buf := make([]byte, 1024)
		f, _ := os.Open(path)
		defer f.Close()
		r := bufio.NewReader(f)
		w := bufio.NewWriter(os.Stdout)
		defer w.Flush()
		for {
			n, _ := r.Read(buf)
			if n == 0 {
				break
			}
			w.Write(buf[:n])
		}
	}
}
