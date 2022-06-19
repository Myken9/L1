package L1

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func ReadService(N time.Duration) {
	randStr := func(length int) string {
		rand.Seed(time.Now().UnixNano())
		b := make([]byte, length)
		rand.Read(b)

		return fmt.Sprintf("%x", b)[:length]
	}

	ch := make(chan string)

	go func() { // read
		for {
			fmt.Println(<-ch)
		}
	}()

	ctx, stop := context.WithCancel(context.Background())

	go func() { // write
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Done")
			default:
				ch <- randStr(10)
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second * N)
	stop()
}

func Test_ReadService(t *testing.T) {
	ReadService(10)
}
