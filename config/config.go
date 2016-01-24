package config

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

// GetStorageDir returns the storage dir from the config file
func GetStorageDir() string {
	var storageDir string
	usr, err := user.Current()
	// read whole the file
	configFile, err := readLines(fmt.Sprintf("%v/.huk", usr.HomeDir))
	errCheck(err)

	for _, line := range configFile {
		if strings.Contains(line, "directory") {
			storageDir = strings.Split(line, "=")[1]
		}
	}
	return storageDir
}

// Init runs the sequence to write the config file
func Init() {
	GetStorageDir()
	username := inputUsername()
	storageDir := inputStorageDir()
	writeConfig(username, storageDir)
}

// readLines of the given filepath
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// standard err check, for brevity
func errCheck(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

// pathExits returns true if file exists, false if it does not
func pathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// inputUsername prompts the user to input a username
func inputUsername() string {
	var username string
	fmt.Printf("Please enter a username: ")
	fmt.Scanf("%s\n", &username)
	return username
}

// inputStorageDirectory prompts the user for a directory
// if blank defaults to ~/huk
func inputStorageDir() string {
	var storageDir string
	fmt.Printf("Enter a directory for your huk folder[default ~/huk]: ")
	fmt.Scanf("%s\n", &storageDir)
	if storageDir == "" {
		storageDir = "~/huk"
	}
	return storageDir
}

// converts all ~ in directories to the current users absolute home dir
// creates the storage directory if it doesnt exist
// updates/creates the config file
func writeConfig(username string, storageDir string) {
	// get user
	usr, err := user.Current()
	errCheck(err)

	storageDir = strings.Replace(storageDir, "~", usr.HomeDir, 1)
	configFilePath := usr.HomeDir + "/.huk"

	// Create huk folder if non existant
	if !pathExists(storageDir) {
		err := os.MkdirAll(storageDir, os.ModePerm)
		errCheck(err)
	}

	// Write config file with username and home directory location.
	var config string
	config = fmt.Sprintf("username=%s\ndirectory=%s\n", username, storageDir)

	// create config file
	f, err := os.Create(configFilePath)
	errCheck(err)

	w := bufio.NewWriter(f)
	_, err = w.WriteString(config)
	w.Flush()
	errCheck(err)

	return
}
