package main

const (
	windowTitle    string  = "Wordle-ish"
	windowWidth    int     = 435
	windowHeight   int     = 600
	blockSize      int     = 75
	screenDpi      float64 = 72.0
	rows           int     = 6
	columns        int     = 5
	blockCount     int     = rows * columns
	alphabet       string  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	dictionaryFile string  = "dictionary.txt"
)
