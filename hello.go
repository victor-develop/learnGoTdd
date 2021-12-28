package main

import "fmt"

const enlglishHelloPrefix = "Hello, "

func Hello(name string) string {
  return enlglishHelloPrefix + name
}

func main() {
  fmt.Println(Hello("Chris"))
}

