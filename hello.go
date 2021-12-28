package main

import "fmt"

const enlglishHelloPrefix = "Hello, "

func Hello(name string) string {
  if name == "" {
    name = "World"
  }
  return enlglishHelloPrefix + name
}

func main() {
  fmt.Println(Hello("Chris"))
}
