package main

import (
  "bufio"
  "fmt"
  "github.com/nchudleigh/huk/key"
  "log"
  "os"
  "os/user"
  // "server"
)

func check(e error) {
  if e != nil {
    fmt.Println(e)
    panic(e)
  }
}

func exists(path string) bool {
    _, err := os.Stat(path)
    if os.IsNotExist(err) { return false }
    return true
}

// Get username.
func getInputUsername() string {
  var username string
  fmt.Printf("Please enter a username: ")
  fmt.Scanf("%s\n", &username)
  return username
}

// Select a directory.
func getInputDirectory() string {
  var hukDir string
  fmt.Println("Enter a directory for your huk folder. Leave blank for default ~/huk.")
  fmt.Scanf("%s\n", &hukDir)
  return hukDir
}

// Create default huk .config filepath
func getConfigFilepath() string {
  var hukConfigFilepath string
  usr, err := user.Current()
  check(err)
  hukConfigFilepath = usr.HomeDir + "/huk/.config"
  return hukConfigFilepath
}

func writeConfig(username string, hukDir string, hukConfigFilepath string) {
  // Create huk folder if defaut option not chosen.
  if hukDir != "" && !exists(hukDir) {
    err := os.Mkdir(hukDir, 0777)
    check(err)
  }

  // Write config file with username and home directory location.
  var config string
  if hukDir != "" {
    config = fmt.Sprintf("%s\n%s", username, hukDir)
  } else {
    config = fmt.Sprintf("%s\n%s", username, hukConfigFilepath)
  }
  f, err := os.Create(hukConfigFilepath)
  check(err)
  w := bufio.NewWriter(f)
  n, err := w.WriteString(config)
  w.Flush()
  check(err)
  if n <= 1 {
    panic("Failed to write config")
  }
}

func main() {

  var filename string
  var myKey string

  args := os.Args[1:]
  action := args[0]

  switch action {
  case "init":
    // user.Init()
    hukConfigFilepath := getConfigFilepath()
    username := getInputUsername()
    hukDir := getInputDirectory()
    writeConfig(username, hukDir, hukConfigFilepath)
  case "send":
    // server
    filename = args[1]
    myKey = key.AddrToKey(key.MyAddress())
    fmt.Printf(
      "The key for your file (%v) is %v.\n"+
      "Tell your friend to run '$ huk %v'\n"+
      "Waiting for connection...\n",
      filename,
      myKey,
      myKey,
    )
    // create server on port_x
    // listen for connections
    // validate incoming request with given key
    // connection established
    // recieves clients public key
    // encrypt file using client's pub key
    // send encrypted file over stream to client
  case "get":
    fmt.Printf(
      "Searching for '%v' on your local network..\n",
      myKey,
    )
    // Client Case
    myKey = args[0]
    // make sure key doesnt have anything but alphabet
    if !isAlpha(myKey) {
      log.Fatal("Key may only contain Lowercase Alphabetic characters")
    }
    // Find server IP by going through list (192.168.0.[1..255]:port_x)
    // connection established
    // generate pgp (private and public keys)
    // send public key to server
    // save encrypted file from stream
    // decrypt using private key
  default:
    // Invalid Args
    log.Fatal("I need either a filename or a key ex: '$ huk -f filename.txt' or '$ huk key'")
  }
}

func isAlpha(input string) bool {
  for _, c := range input {
    if 'a' > c || c > 'z' {
      return false
    }
  }
  return true
}
