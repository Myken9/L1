package L1

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

func closeChan(in <-chan interface{}) {
	for range in {
	}
	fmt.Println("close chan")
}

func cancelCtx(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("cancelCtx")
}

func chanWaitForRead(in <-chan interface{}) {
	<-in
	fmt.Println("chanWaitForRead")
}

func chanWaitForWriteStruct(in chan<- struct{}) {
	in <- struct{}{}
	fmt.Println("chanWaitForWriteStruct")
}

func Exit() {
	os.Exit(1)
}

func Test_gorutinesStop(t *testing.T) {

	//1 close chan from writer
	ch := make(chan interface{})
	go closeChan(ch)
	time.Sleep(time.Millisecond)
	close(ch)

	//2 cancel context
	ctx, cancel := context.WithCancel(context.Background())
	go cancelCtx(ctx)
	time.Sleep(time.Millisecond)
	cancel()
	time.Sleep(time.Millisecond)

	//3 read from chan
	ch1 := make(chan interface{})
	go chanWaitForRead(ch1)
	time.Sleep(time.Millisecond)
	ch1 <- struct{}{}
	time.Sleep(time.Millisecond)

	//4 write to chan
	ch2 := make(chan struct{})
	go chanWaitForWriteStruct(ch2)
	time.Sleep(time.Millisecond)
	<-ch2
	time.Sleep(time.Millisecond)

}
