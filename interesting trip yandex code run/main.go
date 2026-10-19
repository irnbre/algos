package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
)

type Town struct {
	x int 
	y int
}

func main() {
	towns := 0
	file, _:= os.Open("data.txt")
	reader := bufio.NewReader(file)
	fmt.Fscan(reader, &towns)
	var x int 
	var y int 
	sl := []Town{}
	for i := 0; i < towns; i++ {
		fmt.Fscan(reader, &x, &y)
		sl = append(sl, Town{x: x, y: y})
	}
	var benz int
	var from, to int
	fmt.Fscan(reader, &benz)
	fmt.Fscan(reader, &from, &to)

	queue := []Town{{x: sl[from - 1].x, y: sl[from - 1].y}}
	visited := make(map[Town]bool)
	pred := make(map[Town]Town)
	
	for len(queue) != 0 {
		
		current := queue[0]
		_, ok := visited[Town{current.x, current.y}]
		if ok {
			queue = queue[1:]
			continue
		}
		visited[Town{current.x, current.y}] = true
		queue = queue[1:]
		if current.x != sl[to - 1 ].x || current.y != sl[to - 1].y{
			for _, town := range sl {
				_, ok := visited[Town{town.x, town.y}]
				if current.x == town.x && current.y == town.y || ok{
					continue
				} else {
					if math.Abs(math.Abs(float64(current.x - town.x)) + math.Abs(float64(current.y - town.y))) <= float64(benz) {
						queue = append(queue, town)
						_ , ok :=  pred[Town{town.x, town.y}] 
						if ok {
							continue
						}
						pred[Town{town.x, town.y}] = Town{current.x, current.y}
						
					}
				}
			}
		} else {
			
			path := []Town{{sl[to - 1].x, sl[to - 1].y}}
			cur := Town{current.x, current.y}
			
			for cur.x != sl[from - 1].x || cur.y != sl[from - 1].y {
				path = append(path, pred[Town{cur.x, cur.y}])
				buf := Town{cur.x, cur.y}
				cur.x = pred[Town{buf.x, buf.y}].x
				cur.y = pred[Town{buf.x, buf.y}].y
			}
			fmt.Println(len(path) - 1)
			return 
		}
	}
	fmt.Println("-1")
}