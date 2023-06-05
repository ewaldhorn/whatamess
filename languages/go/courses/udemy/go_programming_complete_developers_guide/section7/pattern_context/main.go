package main

import (
	"context"
	"fmt"
	"time"
)

func sampleOperation(ctx context.Context, msg string, msDelay time.Duration) <-chan string {
	out := make(chan string)

	go func() {
		for {
			select {
			case <-time.After(msDelay * time.Millisecond):
				out <- fmt.Sprintf("%v operation completed.", msg)
				return
			case <-ctx.Done():
				out <- fmt.Sprintf("%v aborted.", msg)
				return
			}
		}
	}()

	return out
}

func main() {
	fmt.Printf("\nPattern: Context.\n")

	ctx := context.Background()
	ctx, cancelCtx := context.WithCancel(ctx)

	webServer := sampleOperation(ctx, "Load Server", 5000)
	microservice := sampleOperation(ctx, "Load Microservice", 2000)
	database := sampleOperation(ctx, "Load DB", 15000)

MainLoop:
	for {
		select {
		case val := <-webServer:
			fmt.Println(val)
		case val := <-microservice:
			fmt.Println(val)
			fmt.Println("Cancel context")
			cancelCtx()
			break MainLoop
		case val := <-database:
			fmt.Println(val)
		}
	}

	fmt.Println(<-database)
	fmt.Println()
}
