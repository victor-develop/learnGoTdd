package main

import "fmt"

const enlglishHelloPrefix = "Hello, "
const french = "French"
const spanish = "Spanish"

func Hello(name string, language string) string {
  if name == "" {
    name = "World"
  }

  prefix := enlglishHelloPrefix
  
  switch language {
    case french:
      prefix = "Bonjour, "
    case spanish:
      prefix = "Hola, "
  }
  return prefix + name
}

func main() {
  fmt.Println(Hello("Chris", ""))
}
