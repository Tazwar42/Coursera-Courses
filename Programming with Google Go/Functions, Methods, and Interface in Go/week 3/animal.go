package main

import (
  "fmt"
  "os"
)

type Animal struct {
  name string
	food string
	locomotion string
  noise string
}

func (a Animal) Eat() string{
   return a.food
}

func (a Animal) Move() string{
   return a.locomotion
}

func (a Animal) Speak() string{
   return a.noise
}


func main() {

  cow := Animal{
    name: "cow",
		food:  "grass",
		locomotion: "walk",
    noise: "moo",
  }

  bird := Animal{
    name: "bird",
		food:  "worms",
		locomotion: "fly",
    noise: "peep",
  }

  snake := Animal{
    name: "snake",
		food:  "mice",
		locomotion: "slither",
    noise: "hsss",
  }

  for {

  var animal_name,operation string
  fmt.Println("Enter the animal name(cow or bird or snake) and Enter the information of the animal(eat or move or speak)")
  fmt.Scanf("%s %s\n", &animal_name, &operation)

  if animal_name == "cow" {
    if operation == "eat"{
      fmt.Println("Food eaten: ",cow.Eat())
    } else if operation == "move" {
      fmt.Println("Locomotion method: ",cow.Move())
    } else if operation == "speak" {
      fmt.Println("Spoken sound: ",cow.Speak())
    }
  } else if animal_name == "bird"{
    if operation == "eat"{
      fmt.Println("Food eaten: ",bird.Eat())
    } else if operation == "move" {
      fmt.Println("Locomotion method: ",bird.Move())
    } else if operation == "speak" {
      fmt.Println("Spoken sound: ",bird.Speak())
    }
  } else if animal_name == "snake"{
    if operation == "eat"{
      fmt.Println("Food eaten: ",snake.Eat())
    } else if operation == "move" {
      fmt.Println("Locomotion method: ",snake.Move())
    } else if operation == "speak" {
      fmt.Println("Spoken sound: ",snake.Speak())
    }
  } else {
    os.Exit(3)
  }

  }



}
