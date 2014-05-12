/* hello.go */
package main
import (
  "fmt"
)



func main(){
  person := map[string]string{
    "Name": "Luke",
    "Phone": "774-573-8963",
  }
  name := person["Name"];
  phone := person["Phone"];
  fmt.Printf("Hello %s, your phone number is %s",name,phone);
}
