package logging

import (
  "fmt"
  "io"
  "log"
  "os"
  "path/filepath"
  "runtime"
)

func initializeLog(logFileName string) (logFile *os.File, err error) {
  log.SetFlags(log.LstdFlags | log.Lshortfile)
  logFile, err = os.OpenFile(getLogFilePath(logFileName),
    os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
  if err != nil {
    fmt.Println("os.OpenFile get error: ", err)
    return
  }
  log.SetOutput(logFile)
  mw := io.MultiWriter(os.Stdout, logFile)
  log.SetOutput(mw)
  return
}

func getLogFilePath(logFileName string) (path string) {
  if runtime.GOOS == "windows" {
    appPath, err := getExecPath()
    if err != nil {
      return logFileName
    } else {
      return filepath.Join(appPath, logFileName)
    }
  }
  return filepath.Join("/var/log/", logFileName)
}

func getExecPath() (string, error) {
  return filepath.Abs(filepath.Dir(os.Args[0]))
}
