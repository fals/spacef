package main

import (
  "fmt"
  "os"
  "spacef/telemetry/telemetry_server"

)

func main() {
  fmt.Println("Start server...")

  server, err := telemetry_server.Connect(15002)
  checkError(err)

  fmt.Println("Server started...")

  handler := telemetry_server.CreateHandler();
  
  fmt.Println("Handling connections...")

  server.HandleConnections(handler);
}

func checkError(err error) {
  if err != nil {
      fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
      os.Exit(1)
  }
}