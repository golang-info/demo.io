package main

/* 例子中定义了一个buffer为0的channel done, 子协程运行着定时任务。
如果主协程需要在某个时刻发送消息通知子协程中断任务退出，那么就可以让子协程监听这个done channel，
一旦主协程关闭done channel，那么子协程就可以推出了，
这样就实现了主协程通知子协程的需求。这很好，但是这也是有限的。 

作者：Turling_hu
链接：https://juejin.im/post/5e52688c51882549417fc671
来源：掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。*/

import "fmt"
import "time"

func main() {
	messages := make(chan int, 10)
	done := make(chan bool)

	defer close(messages)

	//consumer
	go func(){
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}()

	//producer
	for i := 0; i < 10; i++ {
		messages <- i
	}
	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}