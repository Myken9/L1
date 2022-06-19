package L1

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"testing"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
//которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func Service(workers int) {
	ch := make(chan string)
	randStr := func(length int) string {
		rand.Seed(time.Now().UnixNano())
		b := make([]byte, length)
		rand.Read(b)

		return fmt.Sprintf("%x", b)[:length]
	}

	go func() {
		t := time.NewTicker(time.Second)
		for range t.C {
			ch <- randStr(10)
		}
	}()

	worker := func(ch chan string) {
		for {
			fmt.Println(<-ch)
		}
	}

	for i := 0; i < workers; i++ {
		go worker(ch)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()
	fmt.Println("Done")
}

func Test_worker(t *testing.T) {
	Service(10)
}
