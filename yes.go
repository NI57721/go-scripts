package main

import (
  "fmt"
  "io"
  "os"
  "strings"
)

type yesReader struct{
  yes string
  reader io.Reader
}

func (r *yesReader) Read(b []byte) (int, error) {
  n, err := r.reader.Read(b)
  if err != nil {
    return n, err
  }
  for i := range b {
    b[i] = byte(r.yes[i])
  }
  return len(b), nil
}

func isNoCommandlineArgs() bool {
  return len(os.Args) == 1
}

func includeHelpCommandlineArg() bool {
  return os.Args[1] == "-h" || os.Args[1] == "--help"
}

func execHelp() {
  fmt.Println(`Usage: yes String or Option
	-h --help: display this message`)
}

func yesOnece(yes string) {
  s := strings.NewReader(yes)
  r := yesReader{yes, s}
  b := make([]byte, len(yes))
  for {
    _, err := r.Read(b)
    fmt.Println(string(b))
    if err == io.EOF {
      break
    }
  }
}

func main() {
  var yes string
  switch {
  case isNoCommandlineArgs():
    yes = "y"
  case includeHelpCommandlineArg():
    execHelp()
    return
  default:
    yes = strings.Join(os.Args[1:], " ")
  }

  for { yesOnece(yes) }
}

