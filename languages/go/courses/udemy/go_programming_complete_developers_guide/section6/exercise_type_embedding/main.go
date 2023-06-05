package main

import (
	"fmt"
)

type Bytes int
type Celsius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celsius
}

type MemoryUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) AverageBandwidth() int {
	sum := 0

	for _, v := range b.amount {
		sum += int(v)
	}

	return sum / len(b.amount)
}

func (c *CpuTemp) AverageTemp() float32 {
	var sum float32 = 0.0

	for _, v := range c.temp {
		sum += float32(v)
	}

	return sum / float32(len(c.temp))
}

func (m *MemoryUsage) AverageMemory() int {
	sum := 0

	for _, v := range m.amount {
		sum += int(v)
	}

	return sum / len(m.amount)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	fmt.Printf("\nType Embedding Exercise\n")

	bandwidth := BandwidthUsage{[]Bytes{50_000, 100_000, 130_000, 80_000, 90_000}}
	temp := CpuTemp{[]Celsius{50.1, 51.1, 53.3, 51.6, 52.8}}
	memory := MemoryUsage{[]Bytes{800_000, 800_000, 810_000, 820_000, 800_000}}

	fmt.Printf("The average bandwidth is %d.\n", bandwidth.AverageBandwidth())
	fmt.Printf("The average temp is %f.\n", temp.AverageTemp())
	fmt.Printf("The average RAM is %d.\n", memory.AverageMemory())

	dash := Dashboard{bandwidth, temp, memory}

	fmt.Println("Dash:", dash)
	fmt.Printf("The average bandwidth is %d.\n", dash.AverageBandwidth())
	fmt.Printf("The average temp is %f.\n", dash.AverageTemp())
	fmt.Printf("The average RAM is %d.\n", dash.AverageMemory())

}
