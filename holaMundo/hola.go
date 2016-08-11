/**
 * Hola mundo
 */

package main

import (
  "fmt"
  "math"
)


func main() {

  // uniendo cadenas
	fmt.Println("Hello " + "playground");

  // cadenas y sumas
  fmt.Println("5+5=", 5+5)

  // usando flotantes
  fmt.Println("7.5/3.2=", 7.5/3.2);
  fmt.Println("7/3=", 7/3);
  fmt.Println("7.0/3.0=", 7.0/3.0);

  // usando booleanos
  fmt.Println("resultado=", true && false)
  fmt.Println("resultado=", true || false)
  fmt.Println("resultado=", !false)

  // declarando la variable a
  var a string = "Initial"
  fmt.Println(a)

  // declarando otra variable
  var b,c int = 4, 5
  fmt.Println(b,c)

  // declarando booleanos
  var d bool = true
  fmt.Println(d)

  // declarando enteros
  var e int = 4
  fmt.Println(e)

  // declaración
  f := "short"
  fmt.Println(f)

  // Constantes
  const s string = "Soy una constante"
  const n =  500
  const m = 3e20 / n

  fmt.Println(int64(m))

  fmt.Println(math.Sin(n));

}
