package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	//signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGINT, syscall.SIGHUP)
	//signal.Ignore(syscall.SIGHUP)
	fmt.Println(os.Getpid())
	for s := range c {
		fmt.Println("Got signal:", s)
	}
}

// kill -s SIGHUP 1234
