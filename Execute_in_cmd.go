package main

import (
		"fmt"
  	 	"os/exec"
    	"sync"
   		"strings"
   		"time"
)

func exe_cmd(cmd string, wg *sync.WaitGroup) {
    fmt.Println(cmd)
    parts := strings.Fields(cmd)
    out, err := exec.Command(parts[0],parts[1]).Output()
    if err != nil {
        fmt.Println("error occured")
        fmt.Printf("%s\n", err)
    }
    fmt.Printf("%s", out)
    wg.Done()
}

func main() {
	duration := time.Second
    wg := new(sync.WaitGroup)
    commands := []string{"Spotify start", "test.go go run", "test.go go run"}
    for _, str := range commands {
        wg.Add(1)
        go exe_cmd(str, wg)
        time.Sleep(duration)
    }
    wg.Wait()

}