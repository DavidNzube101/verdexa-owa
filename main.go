package main

import (
    "log"
    "os"
    "os/exec"
)

func main() {
    cmd := exec.Command("python3", "app.py")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
        log.Fatalf("Failed to run Python script: %v", err)
    }
}