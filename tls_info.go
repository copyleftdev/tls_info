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
	host, filePath := parseFlags()

	// Error handling for both flags missing
	if host == "" && filePath == "" {
		fmt.Println("Error: neither a single host nor a file path was provided")
		return
	}

	var wg sync.WaitGroup

	// Check single host first
	if host != "" {
		wg.Add(1)
		go checkTLS(host, &wg)
	}

	// Process file if it's provided
	if filePath != "" {
		if err := processFile(filePath, &wg); err != nil {
			fmt.Println("Error:", err)
		}
	}

	// Wait for all goroutines to complete
	wg.Wait()
}

func parseFlags() (string, string) {
	host := flag.String("host", "", "single host to check")
	file := flag.String("file", "", "file path to be processed")
	flag.Parse()
	return *host, *file
}

func processFile(filePath string, wg *sync.WaitGroup) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go checkTLS(line, wg)
	}

	return scanner.Err()
}

func checkTLS(server string, wg *sync.WaitGroup) {
	defer wg.Done()

	server = strings.TrimSpace(server) // Clean up server name from potential extra spaces
	if server == "" {
		return // Skip empty lines or spaces
	}

	tlsVersion, tlsCipher, err := getTLSInfo(server)
	if err != nil {
		fmt.Println("Error connecting to server:", server, err)
		return
	}

	fmt.Printf("%s\nTLS version: %x\nTLS cipher: %s\n", server, tlsVersion, tlsCipher)
}

func getTLSInfo(server string) (uint16, string, error) {
	conn, err := tls.Dial("tcp", server, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return 0, "", err
	}
	defer conn.Close()

	tlsVersion := conn.ConnectionState().Version
	tlsCipher := tls.CipherSuiteName(conn.ConnectionState().CipherSuite)
	return tlsVersion, tlsCipher, nil
}
