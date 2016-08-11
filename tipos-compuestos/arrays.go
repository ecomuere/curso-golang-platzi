package main

import (
  "fmt"
)

func main () {
  var nombres[5]string

  amigos := [5]string{"Fran","Joan","Carles","Sergio","Y"}

  nombres = amigos

  //for i, nombres := range nombres {
    fmt.Println(nombres)
  //}
}
