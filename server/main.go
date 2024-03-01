package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
)

var (
	challengePrefix = "Challenge:"
	difficulty      = 5
	port = ":8080"
)

var quotes = []string{
	"When the going gets rough - turn to wonder - Parker Palmer",
	"If you have knowledge, let others light their candles in it - Margaret Fuller",
	"A bird doesn't sing because it has an answer, it sings because it has a song - Maya Angelou",
	"We are not what we know but what we are willing to learn - Mary Catherine Bateson",
}

func main() {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Printf("Server listening on port %s\n", ln.Addr().String())

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	challenge := generateChallenge()

	_, err := conn.Write([]byte(challengePrefix + challenge + "\n"))
	if err != nil {
		log.Println("Error sending quote:", err)
		return
	}
	log.Println("Sent challenge:", challenge)

	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	solution := scanner.Text()

	log.Println("Received solution:", solution)

	if !verifyProofOfWork(solution) {
		conn.Write([]byte("Proof of Work verification failed!\n"))
		return
	}

	conn.Write([]byte(getRandomQuote() + "\n"))
}

func verifyProofOfWork(solution string) bool {
	hash := sha256.Sum256([]byte(solution))
	hashString := hex.EncodeToString(hash[:])
	return strings.HasPrefix(hashString, strings.Repeat("0", difficulty))
}

func getRandomQuote() string {
	return quotes[rand.Intn(len(quotes))]
}

func generateChallenge() string {
	return strconv.Itoa(rand.Int())
}
