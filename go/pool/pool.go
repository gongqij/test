package pool

import (
	"fmt"
	"sync"
)

// outBufPool pools the record-sized scratch buffers used by writeRecordLocked.
var outBufPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("创建一个新的[]byte")
		return new([]byte)
	},
}

func pool(index int) {
	outBufPtr := outBufPool.Get().(*[]byte)
	outBuf := *outBufPtr
	defer func() {
		*outBufPtr = outBuf
		outBufPool.Put(outBufPtr)
	}()
	_, outBuf = sliceForAppend(outBuf[:0], 5)
	for i := 0; i < index; i++ {
		outBuf = append(outBuf, 'a')
	}
}

func sliceForAppend(in []byte, n int) (head, tail []byte) {
	if total := len(in) + n; cap(in) >= total {
		head = in[:total]
	} else {
		head = make([]byte, total)
		copy(head, in)
	}
	tail = head[len(in):]
	return
}
