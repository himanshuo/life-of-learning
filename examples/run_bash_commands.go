package main

import (
    "bytes"
    "io"
    "log"
    "os"
    "os/exec"
    "sync"
    "time"
)

func main() {
    runCatFromStdinWorks(populateStdin("put LICENSE"))
    //runCatFromStdinWorks(populateStdin("bbb\n"))
}

func populateStdin(str string) func(io.WriteCloser) {
    return func(stdin io.WriteCloser) {
        defer stdin.Close()
        io.Copy(stdin, bytes.NewBufferString(str))
    }
}

func runCatFromStdinWorks(populate_stdin_func func(io.WriteCloser)) {
    cmd := exec.Command("tftp", "localhost", "10001")
    stdin, err := cmd.StdinPipe()
    if err != nil {
        log.Panic(err)
    }
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        log.Panic(err)
    }
    err = cmd.Start()
    if err != nil {
        log.Panic(err)
    }
    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        defer wg.Done()
        populate_stdin_func(stdin)
    }()
    go func() {
        defer wg.Done()
        time.Sleep(5 * time.Second)
        io.Copy(os.Stdout, stdout)
    }()
    wg.Wait()
    err = cmd.Wait()
    if err != nil {
        log.Panic(err)
    }
}
