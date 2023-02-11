package ssh

import (
	"bytes"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var config *ssh.ClientConfig

// func init() {
// 	key, err := os.ReadFile("/Users/bytedance/.ssh/id_rsa")
// 	if err != nil {
// 		log.Fatalf("unable to read private key: %v", err)
// 		return
// 	}
// 	// Create the Signer for this private key.
// 	signer, err := ssh.ParsePrivateKey(key)
// 	if err != nil {
// 		log.Fatalf("unable to parse private key: %v", err)
// 		return
// 	}
// 	// ssh config
// 	hostKeyCallback, err := knownhosts.New("/Users/bytedance/.ssh/known_hosts")
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	config = &ssh.ClientConfig{
// 		User: "adam.liu",
// 		Auth: []ssh.AuthMethod{
// 			// Use the PublicKeys method for remote authentication.
// 			ssh.PublicKeys(signer),
// 		},
// 		HostKeyCallback: hostKeyCallback,
// 	}
// }

func CopyFile(path string) error {
	conn, err := ssh.Dial("tcp", "10.37.6.177:22", config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
		return err
	}
	sftp, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("unable to init a sftp client: %v", err)
		return err
	}
	srcPath := path
	srcFile, err := os.Open(srcPath)
	if err != nil {
		log.Fatalf("unable to open the file: %v", err)
		return err
	}
	defer srcFile.Close()
	fileName := filepath.Base(path)
	dstPath := "/data00/home/adam.liu/" + fileName
	dstFile, err := sftp.Create(dstPath)
	if err != nil {
		log.Fatalf("unable to open the remote file: %v", err)
		return err
	}
	defer dstFile.Close()
	if _, err := dstFile.ReadFrom(srcFile); err != nil {
		log.Fatalf("unable to copy the file: %v", err)
		return err
	}
	return nil
}

func RunCmd() {
	conn, err := ssh.Dial("tcp", "10.37.6.177:22", config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
		return
	}
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("unable to init a session: %v", err)
		return
	}
	defer session.Close()
	var b bytes.Buffer  // import "bytes"
	session.Stdout = &b // get output
	cmd := "ls ~"
	err = session.Run(cmd)
	if err != nil {
		log.Fatalf("error to run the command: %v", err)
	}
}
