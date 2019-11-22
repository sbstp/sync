package sync

import (
	"io"
	"sync"
)

type pipe struct {
	wLock *sync.Mutex
	bufCh chan []byte
	lenCh chan int
}

func newPipe() *pipe {
	return &pipe{
		wLock: new(sync.Mutex),
		bufCh: make(chan []byte),
		lenCh: make(chan int),
	}
}

func (f *pipe) Write(buf []byte) (int, error) {
	f.wLock.Lock()
	defer f.wLock.Unlock()

	total := 0
	for len(buf) > 0 {
		f.bufCh <- buf
		n := <-f.lenCh
		buf = buf[n:]
		total += n
	}
	return total, nil
}

func (f *pipe) Close() error {
	close(f.bufCh)
	return nil
}

func (f *pipe) Read(buf []byte) (int, error) {
	b, ok := <-f.bufCh

	if !ok {
		return 0, io.EOF
	}

	n := copy(buf, b)
	f.lenCh <- n

	return n, nil
}

func TwistedPair() (io.WriteCloser, io.Reader) {
	f := newPipe()
	return f, f
}
