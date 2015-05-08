package commands

import (
	"io"
	"net"

	"github.com/codegangsta/cli"
	"code.google.com/p/go.crypto/ssh"
	log "github.com/Sirupsen/logrus"
)

type clientPassword string

func cmdStart(c *cli.Context) {
	username := c.String("username")
	password := c.String("password")
	localAddrString := c.String("local")

	// Setup SSH config (type *ssh.ClientConfig)
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}

	// Setup localListener (type net.Listener)
	localListener, err := net.Listen("tcp", localAddrString)
	if err != nil {
		log.Fatalf("net.Listen failed: %v", err)
	}

	for {
		log.Infof("Start camp server")

		// Setup localConn (type net.Conn)
		localConn, err := localListener.Accept()
		if err != nil {
			log.Fatalf("listen.Accept failed: %v", err)
		}

		go forward(localConn, config, c)
	}
}

func (password clientPassword) Password(user string) (string, error) {
	return string(password), nil
}

func forward(localConn net.Conn, config *ssh.ClientConfig, c *cli.Context) {
	serverAddrString := c.String("server")
	remoteAddrString := c.String("remote")

	// Setup sshClientConn (type *ssh.ClientConn)
	sshClientConn, err := ssh.Dial("tcp", serverAddrString, config)
	if err != nil {
		log.Fatalf("ssh.Dial failed: %s", err)
	}

	// Setup sshConn (type net.Conn)
	sshConn, err := sshClientConn.Dial("tcp", remoteAddrString)

	// Copy localConn.Reader to sshConn.Writer
	go func() {
		_, err = io.Copy(sshConn, localConn)
		if err != nil {
			log.Fatalf("io.Copy failed: %v", err)
		}
	}()

	// Copy sshConn.Reader to localConn.Writer
	go func() {
		_, err = io.Copy(localConn, sshConn)
		if err != nil {
			log.Fatalf("io.Copy failed: %v", err)
		}
	}()
}
