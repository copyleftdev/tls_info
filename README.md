# TLS Info Tool

The `tls_info` tool is designed to check TLS versions and cipher suites used by servers. It supports checking a single server as well as a list of servers from a file.

## Prerequisites

Before you can run or build the `tls_info` tool, ensure you have the Go language installed on your system. You can download and install Go from [https://golang.org/dl/](https://golang.org/dl/).

## Installation

To get started with `tls_info`, clone the repository using Git:

```bash
git clone https://github.com/copyleftdev/tls_info.git
cd tls_info
```

## Compilation

Compile the `tls_info` tool using the Go compiler:

```bash
go build -o tls_info
```

This command will create an executable named `tls_info` in the current directory.

## Usage

After compilation, you can run the `tls_info` tool using one of the following methods:

### Single Host

To check the TLS version and cipher suite for a single host, use the `-host` flag:

```bash
./tls_info -host "example.com:443"
```

### File With List of Hosts

To check multiple hosts from a file, each host on a new line, use the `-file` flag:

```bash
./tls_info -file "path/to/your/file.txt"
```

Each line in the file should contain one host, optionally including the port (default is 443 if not specified).

## Output

The tool prints the TLS version and cipher suite for each server to the standard output in the following format:

```
server_name
TLS version: [version]
TLS cipher: [cipher_suite]
```

## Contributing

Contributions to `tls_info` are welcome. Please fork the repository, make your changes, and submit a pull request.
