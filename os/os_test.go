package os

import (
  "runtime"
  "testing"
)

func TestExecuteCommand(t *testing.T) {
  var err error
  var outStr, errStr string
  if runtime.GOOS == "windows" {
    err, outStr, errStr = ExecuteCommand("ipconfig.exe", "/all")
  } else {
    err, outStr, errStr = ExecuteCommand("ifconfig", "-a")
  }
  if err != nil || len(errStr) > 0 || len(outStr) == 0 {
    t.Error()
  }

  if runtime.GOOS == "windows" {
    err, outStr, errStr = ExecuteCommand("not_exist.exe", "-a")
  } else {
    err, outStr, errStr = ExecuteCommand("not_exist", "-a")
  }
  if err == nil && len(errStr) == 0 {
    t.Error()
  }
}

func TestGetHostname(t *testing.T) {
  hostName, err := GetHostname()
  if err != nil || len(hostName) == 0 {
    t.Error()
  }
}

func TestGetPidByName(t *testing.T) {
  var err error
  var pid int
  if runtime.GOOS != "windows" {
    pid, err = GetPidByName("bash")
    if err != nil || pid == 0 {
      t.Error()
    }
  }
}
