package main

import "fmt"

func main () {
  alumnos := 1
  for alumnos <= 3 {
    fmt.Println(alumnos)
    alumnos = alumnos +1
  }

  for calificaciones := 7; calificaciones <= 9; calificaciones++ {
    fmt.Println(calificaciones)
  }
}
