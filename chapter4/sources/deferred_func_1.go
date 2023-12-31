package chapter4

import (
	"os"
	"sync"
)

func writeFile(fname string, data []byte, mu *sync.Mutex) error {
	mu.Lock()
	f, err := os.OpenFile(fname, os.O_RDWR, 0600)
	if err != nil {
		mu.Unlock()
		return err
	}

	_, err = f.Seek(0, 2)
	if err != nil {
		f.Close()
		mu.Unlock()
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		f.Close()
		mu.Unlock()
		return err
	}
	err = f.Sync()
	if err != nil {
		f.Close()
		mu.Unlock()
		return err
	}
	err = f.Close()
	if err != nil {
		mu.Unlock()
		return err
	}
	mu.Unlock()
	return nil
}
