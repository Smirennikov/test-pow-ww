package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	challengePrefix = "Challenge:"
	difficulty      = 5
	serverPort      = ":8080"
)

func main() {

	conn, err := net.Dial("tcp", serverPort)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Printf("Client dial on port %s\n", conn.RemoteAddr().String())

	scanner := bufio.NewScanner(conn)
	scanner.Scan()

	challenge := strings.TrimPrefix(scanner.Text(), challengePrefix)
	solution := solveProofOfWork(challenge)

	if _, err := conn.Write([]byte(solution + "\n")); err != nil {
		log.Println("Error sending solution:", err)
		return
	}
	for scanner.Scan() {
		log.Println("Received quote:", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func solveProofOfWork(challenge string) string {
	var solution string
	for i := 0; ; i++ {
		solution = fmt.Sprintf("%s%d", challenge, i)
		hash := sha256.Sum256([]byte(solution))
		hashString := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hashString, strings.Repeat("0", difficulty)) {
			break
		}
	}
	return solution
}
