package main

import (
  "fmt"
  "os"
)

type Animal interface {
  Eat() string
  Move() string
  Speak() string
}

type Cow struct {
	name, food, locomotion, noise string
}

type Bird struct {
  name, food, locomotion, noise string

}

type Snake struct {
  name, food, locomotion, noise string
}

func (c Cow) Eat() string{
  return c.food
}

func (c Cow) Move() string{
  return c.locomotion
}

func (c Cow) Speak() string{
  return c.noise
}

func (b Bird) Eat() string{
  return b.food
}

func (b Bird) Move() string{
  return b.locomotion
}

func (b Bird) Speak() string{
  return b.noise
}

func (s Snake) Eat() string{
  return s.food
}

func (s Snake) Move() string{
  return s.locomotion
}

func (s Snake) Speak() string{
  return s.noise
}

func operation(a Animal, operation string) string{
  if operation == "eat" {
    return a.Eat()
  }else if operation == "move" {
    return a.Move()
  }else if operation == "speak" {
    return a.Speak()
  }else {
    return "No Operation found"
  }
}


func main() {

   for {

   var newanimal, animal_name, animal_type string
   fmt.Println("Enter a new animal and its type in a single line by giving space to create.....(I.E newanimal c cow, newanimal b bird, newanimal s snake).....Please follow this pattern.")
   fmt.Scanf("%s %s %s\n",&newanimal, &animal_name, &animal_type)

   if animal_type == "cow" {
    cow := Cow{
    name: animal_name,
		food:  "grass",
		locomotion: "walk",
    noise: "moo",
    }
    fmt.Printf("Your animal %s is created!!\n", cow.name)


    var query, animal_name1, operation_name string
    fmt.Println("Enter the create animals name and its operation in a single line by giving space to get the information.....(I.E query c eat, query b move, query s speak).....Please follow this pattern.")
    fmt.Scanf("%s %s %s\n",&query, &animal_name1, &operation_name)


  if operation_name == "eat" {
    fmt.Println(operation(cow, "eat"))
  } else if operation_name == "move" {
    fmt.Println(operation(cow, "move"))
  } else if operation_name == "speak" {
    fmt.Println(operation(cow, "speak"))
  } else {
    os.Exit(3)
  }

  } else if animal_type == "bird" {
    bird := Bird{
    name: animal_name,
		food:  "worms",
		locomotion: "fly",
    noise: "peep",
    }
    fmt.Printf("Your animal %s is created!!\n", bird.name)

    var query, animal_name1, operation_name string
    fmt.Println("Enter the create animals name and its operation in a single line by giving space to get the information.....(I.E query c eat, query b move, query s speak).....Please follow this pattern.")
    fmt.Scanf("%s %s %s\n",&query, &animal_name1, &operation_name)


  if operation_name == "eat" {
    fmt.Println(operation(bird, "eat"))
  } else if operation_name == "move" {
    fmt.Println(operation(bird, "move"))
  } else if operation_name == "speak" {
    fmt.Println(operation(bird, "speak"))
  } else {
    os.Exit(3)
  }


  } else if animal_type == "snake" {
    snake := Snake{
    name: animal_name,
		food:  "mice",
		locomotion: "slither",
    noise: "hsss",
  }
   fmt.Printf("Your animal %s is created!!\n", snake.name)

   var query, animal_name1, operation_name string
    fmt.Println("Enter the created animals name and its operation in a single line by giving space to get the information.....(I.E query c eat, query b move, query s speak).....Please follow this pattern.")
    fmt.Scanf("%s %s %s\n",&query, &animal_name1, &operation_name)


  if operation_name == "eat" {
    fmt.Println(operation(snake, "eat"))
  } else if operation_name == "move" {
    fmt.Println(operation(snake, "move"))
  } else if operation_name == "speak" {
    fmt.Println(operation(snake, "speak"))
  } else {
    os.Exit(3)
  }
  } else {
    os.Exit(3)
  }

  fmt.Println("......................")






   }



}
