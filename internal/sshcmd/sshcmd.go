package sshcmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
	"multi-node-controller/internal/yamlcustom"
)


// getKey to read private key
func getKey() []uint8 {
	conf := yamlcustom.ParseYAML()
	
	key, err := ioutil.ReadFile(conf.Conf[1].PrivateKey)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	return key
}

// RemoteCommand to execute command to multiply nodes
func RemoteCommand(ip string, cmd string) {
	// Create the Signer for this private key.
	conf := yamlcustom.ParseYAML()

	signer, err := ssh.ParsePrivateKey(getKey())
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: conf.Conf[0].UserID,
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// Connect to the remote server and perform the SSH handshake.
	client, err := ssh.Dial("tcp", ip + ":22", config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer client.Close()

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(cmd); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}
