package os

import (
  "bytes"
  "errors"
  "fmt"
  "github.com/shirou/gopsutil/process"
  "os"
  "os/exec"
  "strings"
)

func ExecuteCommand(name string, arg ...string) (err error, outStr string, errStr string) {
  cmd := exec.Command(name, arg...)
  var stdoutBuffer, stderrBuffer bytes.Buffer
  cmd.Stdout = &stdoutBuffer
  cmd.Stderr = &stderrBuffer
  err = cmd.Run()
  if err != nil {
    return
  }
  outStr = stdoutBuffer.String()
  errStr = stderrBuffer.String()

  return
}

func KillProcByName(procName string) (err error) {
  err, _, errStr := ExecuteCommand("killall", procName)
  if nil != err {
    return err
  }
  if len(errStr) > 0 {
    return errors.New(errStr)
  }
  return nil
}

func GetPidByName(procName string) (pid int, err error) {
  pid = -1
  var procs []*process.Process
  if procs, err = process.Processes(); err != nil {
    return 0, err
  }
  for _, proc := range procs {
    var name string
    if name, err = proc.Name(); err != nil {
      continue
    }

    if name == procName {
      continue
    }

    pid = int(proc.Pid)
    return pid, nil
  }
  return 0, fmt.Errorf("can not find pid of %v", procName)
}

func GetHostname() (hostname string, err error) {
  if hostname, err = os.Hostname(); err != nil {
    return "", err
  }
  segments := strings.Split(hostname, ".")
  hostname = segments[0]
  return
}
