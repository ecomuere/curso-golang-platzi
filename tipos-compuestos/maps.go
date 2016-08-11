package main

import (
  "fmt"
)

func main () {
  departamentos := make(map[string]int)

  departamentos["devs"] = 25
  departamentos["marketing"] = 30
  departamentos["design"] = 15
  departamentos["ejecutivos"] = 3
  departamentos["ventas"] = 40

  for key, value := range departamentos {
    //fmt.Println(key, value)
    fmt.Printf("departamento: %s, empleados %d\n", key, value)
  }
}
