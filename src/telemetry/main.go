package main

import (
  "fmt"
  "os"
  "spacef/telemetry/telemetry_server"
)

func main() {
  fmt.Println("Start server...")

  err := telemetry_server.Connect(15002)
  checkError(err)
}

func checkError(err error) {
  if err != nil {
      fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
      os.Exit(1)
  }
}