package main

import (
  "fmt"
  "sync"
  "time"
  "math/rand"
)


type ChopS struct{ sync.Mutex }

type Philo struct {
  id              int
	leftCS, rightCS *ChopS
}

func (p Philo) eat() {
   for j := 0; j < 3; j++{
      randomPause(2)

      p.leftCS.Lock()
      p.rightCS.Lock()

      fmt.Printf("starting to eat %d\n", p.id)

      randomPause(5)
      fmt.Printf("finish eating %d\n", p.id)

      p.rightCS.Unlock()
      p.leftCS.Unlock()
   }
   eatWgroup.Done()
}



func randomPause(max int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(max*1000)))
}

var eatWgroup sync.WaitGroup

func main() {

CSticks := make([]*ChopS, 5)
for i := 0; i < 5; i++ {
   CSticks[i] = new(ChopS)
}

philos := make([]*Philo, 5)
for i := 0; i < 5; i++ {
   philos[i] = &Philo{id: i,leftCS: CSticks[i],rightCS: CSticks[(i+1)%5]}
}

for i := 0; i <5; i++ {
    eatWgroup.Add(1)
   go philos[i].eat()
}
eatWgroup.Wait()



}
