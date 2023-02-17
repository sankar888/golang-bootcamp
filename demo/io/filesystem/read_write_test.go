package filesystem

import (
	"testing"
	"os"
	"os/signal"
	"path/filepath"
	"context"
	"sync"
	"time"
)

//can a same file be be written to and read at the same time
func TestReadWrite(t *testing.T) {
	path := "C:/Users/sankaraa/work/tmp/buffer.log"
	if err := os.MkdirAll(filepath.Dir(path), os.ModeDir); err != nil {
		t.Fatal("could not create directory", err)
	}

	if file, err := os.Create(path); err != nil {
		t.Fatal("could not create file", err)
	} else {
		defer file.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		var wg *sync.WaitGroup = &sync.WaitGroup{}
		wg.Add(2)

		go writetoFile(ctx, file, wg,  t)
		go readFromFile(ctx, file, wg, t)
		waitAndShutdown(cancel, wg, t)
	}
}

func writetoFile(ctx context.Context, file *os.File, wg *sync.WaitGroup, t *testing.T) {
	loop:
	for {
		select{
		case <-time.After(time.Second * 2):
			if _, err := file.WriteString("hai, How are u ? \n"); err != nil {
				wg.Done()
				t.Fatal(err)
			} else {
				file.Sync()
			}
		case <-ctx.Done():
			t.Log("write function shuting down")
			file.Sync()
			wg.Done()
			break loop
		}
	}
}

func readFromFile(ctx context.Context, file *os.File, wg *sync.WaitGroup, t *testing.T) {
	var buffer []byte = make([]byte, 1024)
	loop:
	for {
		select{
		case <-time.After(time.Second * 5):
			if n, err := file.Read(buffer); err != nil {
				wg.Done()
				t.Fatal(err)
			} else {
				t.Logf("read from file: %s", string(buffer[0:n]))	
			}
			
		case <-ctx.Done():
			t.Log("read function shuting down")
			wg.Done()
			break loop
		}
	}
}

func waitAndShutdown(cancel context.CancelFunc, wg *sync.WaitGroup,  t *testing.T) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	t.Log("listening for interrupt")
	<-c
	t.Log("main process interrupted. Cancelling the context to signal the routines to stop")
	cancel()
	t.Log("issued cancel signal to go routines. waiting for all routines to complete")
	wg.Wait()
	t.Log("All goroutines completed. Exiting")
}