package main

import (
	"fmt"
	"sync"
)

type SyncedData struct {
	inner map[string]int
	mutex sync.Mutex
}

func (d *SyncedData) Insert(k string, v int) {
	d.mutex.Lock()
	d.inner[k] = v
	d.mutex.Unlock()
}

func (d *SyncedData) Get(k string) int {
	d.mutex.Lock()
	data := d.inner[k]
	d.mutex.Unlock()
	return data
}

func runMutexExample1() {
	fmt.Println("Without defer:")
	data := SyncedData{inner: make(map[string]int)}
	data.Insert("sample", 5)
	data.Insert("test", 2)
	fmt.Println(data.Get("sample"))
	fmt.Println(data.Get("test"))
}
