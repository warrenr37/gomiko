package gomiko

import (
    "log"
    "bytes"
    "strconv"

    "golang.org/x/crypto/ssh"
)

func ConnectHandler(ip string, username string, password string, port int) *ssh.Client {
    var host string = ip + ":" + strconv.Itoa(port)
    log.Print(host)
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}

    return client
}

func SendCommand(client *ssh.Client,command string) string {
	// Creates a session to which a single command is passed.
    // Then the session is closed after the output is returned as a string
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// 
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(command); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
    return b.String()
}
