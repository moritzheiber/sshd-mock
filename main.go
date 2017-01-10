package main

import (
	"github.com/gliderlabs/ssh"
	"io"
	"log"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, "Hello world\n")
	})

	log.Fatal(ssh.ListenAndServe(":22", nil))
}
