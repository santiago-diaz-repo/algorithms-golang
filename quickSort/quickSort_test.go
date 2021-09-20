package quickSort

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"testing"
)

func Test_Sort(t *testing.T)  {
	arr := []int{4,78,-6,4,1,0,43,21,14}
	want := "[-6 0 1 4 4 14 21 43 78]"
	got := captureOutput(func() {
		QuickSort(&arr,0,len(arr)-1)
		fmt.Print(arr)
	})
	if got != want{
		t.Errorf("wanted: %s, got %s", want, got)
	}

}

// captureOutput is explained in https://medium.com/@hau12a1/golang-capturing-log-println-and-fmt-println-output-770209c791b4
func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}
