package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "top")
	//cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.Stdout = os.Stdout
	cmd.Start()

	time.Sleep(1 * time.Second)
	fmt.Println("退出程序中...", cmd.Process.Pid)
	cancel()

	cmd.Wait()
	fmt.Println("ok")
	fmt.Println("OK")
	fmt.Println("ok")

}
