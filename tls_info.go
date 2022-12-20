package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	// Define a flag for the file path
	file := flag.String("file", "", "file path")

	// Parse the command-line flags
	flag.Parse()

	// Check if the file flag was provided
	if *file == "" {
		fmt.Println("Error: file path not provided")
		return
	}

	// Open the file
	f, err := os.Open(*file)
	if err != nil {
		// Handle the error
		return
	}
	defer f.Close()

	// Use a wait group to wait for all Go routines to finish
	var wg sync.WaitGroup

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(f)

	// Read the file line by line
	for scanner.Scan() {
		// Split the line by tabs
		fields := strings.Split(scanner.Text(), "\t")

		// Check if there are enough fields
		if len(fields) < 1 {
			continue
		}

		// Get the server address from the first field
		server := fields[0]

		// Start a new Go routine to check the TLS version and cipher of the server
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Create a TLS connection to the server
			conn, err := tls.Dial("tcp", server, &tls.Config{
				InsecureSkipVerify: true,
			})
			if err != nil {
				// Handle the error
				return
			}
			defer conn.Close()

			// Get the TLS version and cipher used by the connection
			tlsVersion := conn.ConnectionState().Version
			tlsCipher := tls.CipherSuiteName(conn.ConnectionState().CipherSuite)

			// Print the TLS version and cipher
			fmt.Printf("%s\nTLS version: %d\nTLS cipher: %s\n", server, tlsVersion, tlsCipher)
		}()
	}

	// Wait for all Go routines to finish
	wg.Wait()
}
