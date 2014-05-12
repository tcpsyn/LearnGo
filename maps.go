/* hello.go */
package main
import (
  "fmt"
)



func main(){
  // This is a map, it's like a hash table. You have to define
  // the type of both the key and the value.
  person := map[string]string{
    "Name": "Luke",
    "Phone": "774-573-8963",
  }
  name := person["Name"];
  phone := person["Phone"];
  fmt.Printf("Hello %s, your phone number is %s\n",name,phone);

  // This is an array. It's like an array. ... lets golang
  // count the values, otherwise you would have to define
  // how many values are in the array.
  gospels := [...]string{"Matthew","Mark","Luke","John"}

  // This is kind of like a foreach. 
  for _,element := range gospels{
    fmt.Printf("The gospel according to %s\n",element);
  }

}
