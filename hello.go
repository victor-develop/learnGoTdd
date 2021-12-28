package main

import "fmt"

const enlglishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
  if name == "" {
    name = "World"
  }

  if language == "Spanish" {
    return "Hola, " + name
  }

  return enlglishHelloPrefix + name
}

func main() {
  fmt.Println(Hello("Chris", ""))
}
