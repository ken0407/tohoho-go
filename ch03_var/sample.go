package main
import "fmt"

func main() {
  var a1 int
  var a2 int = 123
  var a3 = 123
  a4 := 123

  var (
    a5 int = 123
    a6 int = 456
  )

  var name string = "Yamada"
  var age int = 27

  fmt.Print(a1)
  fmt.Print(a2)
  fmt.Print(a3)
  fmt.Print(a4)
  fmt.Print(a5)
  fmt.Print(a6)


  fmt.Printf("name: %s, age: %d", name, age)
}
