package stdtypes

import (
	"io/ioutil"
	"sync"
	"testing"

	"github.com/waynz0r/protobuf/proto"
)

func TestConcurrentTextMarshal(t *testing.T) {
	// Verify that there are no race conditions when calling
	// TextMarshaler.Marshal on a protobuf message that contains a StdDuration

	std := StdTypes{}
	var wg sync.WaitGroup

	tm := proto.TextMarshaler{}

	// use a channel to hand off the error
	errs := make(chan error, 2)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := tm.Marshal(ioutil.Discard, &std)
			if err != nil {
				errs <- err
			}
		}()
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	for err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}
}
