/* hello.go */
package main
import (
  "fmt"
  "math/rand"
  "time"
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

  // And now a slice.
  // I don't have to define the length of a slice.
  foods := []string{"Cornbread","French Bread","Eggs","Toast"}
  for _,i := range foods{
    fmt.Printf("I like to eat: %s\n",i);
  }

  // Looping through and array and slice.
  for i := 0; i < len(gospels); i++ {
    fmt.Printf("%s likes to eat %s\n",gospels[i],foods[i]);
  }

  fmt.Println();

  // Playing with random numbers, which is way harder than it should be.
  for _,i :=range gospels{
    rand.Seed(time.Now().UTC().UnixNano());
    r := randInt(0,len(foods));
    fmt.Printf("%s likes to eat %s, weirdo.\n",i,foods[r]);
  }
}

func randInt(min int, max int) int{
  return min + rand.Intn(max-min);
}
