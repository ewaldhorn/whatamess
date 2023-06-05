package main

import (
	"fmt"
)

func (d *SyncedData) InsertWithDefer(k string, v int) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.inner[k] = v
}

func (d *SyncedData) GetWithDefer(k string) int {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	data := d.inner[k]
	return data
}

func runMutexExample2() {
	fmt.Println("With defer:")
	data := SyncedData{inner: make(map[string]int)}
	data.InsertWithDefer("sample", 5)
	data.InsertWithDefer("test", 2)
	fmt.Println(data.GetWithDefer("sample"))
	fmt.Println(data.GetWithDefer("test"))
}
