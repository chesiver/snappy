package ssh

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

var config *ssh.ClientConfig
var sftpClient *sftp.Client

func InitConfigForUsernamePassword(username, password string) {
	fmt.Printf("!!!!!!: %v %v\n", username, password)
	config = &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

func InitConfigForPublicKey(keyFilePath string) {
	key, err := os.ReadFile(keyFilePath)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}
	// ssh config
	hostKeyCallback, err := knownhosts.New("/Users/bytedance/.ssh/known_hosts")
	if err != nil {
		log.Fatal(err)
	}
	config = &ssh.ClientConfig{
		User: "adam.liu",
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: hostKeyCallback,
	}
}

func Connect(host string) {
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%v:22", host), config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	sftpClient, err = sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("unable to init a sftp client: %v", err)
	}
}

type FileInfo struct {
	Name    string `json:"name"`
	ModTime string `json:"modTime"`
	IsDir   bool   `json:"isDir"`
}

func ListEntries() (string, []FileInfo) {
	curDir, err := sftpClient.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return ListEntriesForPath(curDir)
}

func ListEntriesForPath(path string) (string, []FileInfo) {
	entries, err := sftpClient.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	arr := make([]FileInfo, 0)
	for _, e := range entries {
		arr = append(arr, FileInfo{
			Name:    e.Name(),
			ModTime: e.ModTime().String(),
			IsDir:   e.IsDir(),
		})
	}
	return path, arr
}

func CopyFile(path string) error {
	srcPath := path
	srcFile, err := os.Open(srcPath)
	if err != nil {
		log.Fatalf("unable to open the file: %v", err)
		return err
	}
	defer srcFile.Close()
	fileName := filepath.Base(path)
	dstPath := "/data00/home/adam.liu/" + fileName
	dstFile, err := sftpClient.Create(dstPath)
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
