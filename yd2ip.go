package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func _d2ip(domain string){
	iprecords, _ := net.LookupIP(domain)
	res:=domain+"\t"
	for _, ip := range iprecords {
		res+=ip.String()+"\t"
	}
	fmt.Println(res)
}


func main() {
	sc:=bufio.NewScanner(os.Stdin)
	thread:=flag.Int("t",50,"threads")
	flag.Parse()
	var wg sync.WaitGroup
	var ch = make(chan struct{}, *thread)
	for sc.Scan(){
		wg.Add(1)
		domain:=strings.TrimSpace(sc.Text())
		ch <- struct{}{} // acquire a token
		go func(domain string) {
			defer wg.Done()
			_d2ip(domain)
			<-ch // release the token
		}(domain)
	}
	wg.Wait()
}
