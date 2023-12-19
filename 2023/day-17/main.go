package main

//
//import (
//	"bufio"
//	"container/heap"
//	"os"
//	"strconv"
//)
//
//var graph [][]Point
//
//var dist map[PointLocation]DirectedDistance
//
//type DirectedDistance struct {
//	dist int
//	path string
//}
//
//type PointLocation struct {
//	i int
//	j int
//}
//
//type Point struct {
//	location     PointLocation
//	heat         int //used as priority in the priority queue
//	neighbors    []*Point
//	neighborDirs []string
//	dist         int
//	dirHistory   string
//	lastDir      string
//	n            int
//	index        int // used by heap
//}
//
//type PriorityQueue []*Point
//
//func (pq PriorityQueue) Len() int { return len(pq) }
//
//func (pq PriorityQueue) Less(i, j int) bool {
//	return pq[i].dist < pq[j].dist
//}
//
//func (pq PriorityQueue) Swap(i, j int) {
//	pq[i], pq[j] = pq[j], pq[i]
//	pq[i].index = i
//	pq[j].index = j
//}
//
//func (pq *PriorityQueue) Push(x any) {
//	n := len(*pq)
//	point := x.(*Point)
//	point.index = n
//	*pq = append(*pq, point)
//}
//
//func (pq *PriorityQueue) Pop() any {
//	old := *pq
//	n := len(old)
//	point := old[n-1]
//	old[n-1] = nil
//	point.index = -1
//	*pq = old[0 : n-1]
//	return point
//}
//
//func main() {
//	input, _ := os.Open("day-17/input.txt")
//	defer input.Close()
//	scanner := bufio.NewScanner(input)
//
//	var matrix []string
//
//	for scanner.Scan() {
//		matrix = append(matrix, scanner.Text())
//	}
//
//	graph = make([][]Point, len(matrix))
//
//	for i := 0; i < len(matrix); i++ {
//		row := make([]Point, len(matrix[0]))
//		for j := 0; j < len(matrix[0]); j++ {
//			heatLoss, _ := strconv.Atoi(string(matrix[i][j]))
//			row[j] = Point{
//				location: PointLocation{
//					i: i,
//					j: j,
//				}, heat: heatLoss}
//		}
//		graph[i] = row
//	}
//
//	for i := 0; i < len(matrix); i++ {
//		for j := 0; j < len(matrix[0]); j++ {
//			p := &graph[i][j]
//			if i > 0 {
//				// add up neighbour to current
//				p.neighbors = append(p.neighbors, &graph[i-1][j])
//				p.neighborDirs = append(p.neighborDirs, "U")
//			}
//			if j > 0 {
//				// left
//				p.neighbors = append(p.neighbors, &graph[i][j-1])
//				p.neighborDirs = append(p.neighborDirs, "L")
//			}
//			if i < len(matrix)-1 {
//				// down
//				p.neighbors = append(p.neighbors, &graph[i+1][j])
//				p.neighborDirs = append(p.neighborDirs, "D")
//			}
//			if j < len(matrix[0])-1 {
//				// right
//				p.neighbors = append(p.neighbors, &graph[i][j+1])
//				p.neighborDirs = append(p.neighborDirs, "R")
//			}
//		}
//	}
//	dist = make(map[PointLocation]DirectedDistance)
//	djikstra(graph, graph[0][0])
//	for k, v := range result {
//		if k.location.i == len(graph)-1 && k.location.j == len(graph[0])-1 {
//			println(k.location.i, " ", k.location.j, " ", v, " ", k.lastDir, " ", k.n)
//		}
//	}
//	println("end")
//
//}
//
//var queue PriorityQueue
//
//type SeenPoint struct {
//	i          int
//	j          int
//	heat       int
//	dist       int
//	dirHistory string
//}
//
//var seen = make(map[SeenPoint]bool)
//
//type Info struct {
//	location PointLocation
//	lastDir  string
//	n        int
//}
//
//var result = make(map[Info]int)
//
//func djikstra(graph [][]Point, source Point) {
//	queue = make(PriorityQueue, 0)
//	heap.Push(&queue, &Point{
//		location:     source.location,
//		heat:         source.heat,
//		neighbors:    source.neighbors,
//		neighborDirs: source.neighborDirs,
//		dist:         0,
//		dirHistory:   "",
//		index:        0,
//	})
//
//	for len(queue) > 0 {
//		minV := heap.Pop(&queue).(*Point)
//
//		sp := Info{
//			location: minV.location,
//			lastDir:  minV.lastDir,
//			n:        minV.n,
//		}
//		if _, ok := result[sp]; ok {
//			continue
//		}
//		result[sp] = minV.dist
//
//		for i, neighbour := range minV.neighbors {
//			distance := minV.dist + neighbour.heat
//			direction := minV.dirHistory + minV.neighborDirs[i]
//
//			notGoingBackwards := true
//			if len(minV.dirHistory) > 0 {
//				notGoingBackwards = string(minV.dirHistory[len(minV.dirHistory)-1]) != opposingPath(minV.neighborDirs[i])
//			}
//
//			//result[neighbour.location THIS PLUS LAST 3 DIRSS
//
//			//less := true
//			//d, ok := result[Info{location: neighbour.location, lastDir: last3Dir(direction)}]
//			//if !ok {
//			//	less = true
//			//} else {
//			//	less = distance < d
//			//}
//			if !areDirSame(direction) && notGoingBackwards {
//				//println(direction)
//				heap.Push(&queue, &Point{
//					location:     neighbour.location,
//					heat:         neighbour.heat,
//					neighbors:    neighbour.neighbors,
//					neighborDirs: neighbour.neighborDirs,
//					dist:         distance,
//					dirHistory:   direction,
//					lastDir:      last3Dir(direction),
//					n:            nLast3Dir(direction),
//				})
//			}
//		}
//	}
//}
//
//func nLast3Dir(dirStr string) int {
//	c := last3Dir(dirStr)
//	if len(dirStr) == 1 {
//		counter := 1
//		for i, iter := len(dirStr)-2, 0; i >= 0 && iter < 3; i, iter = i-1, iter+1 {
//			if string(dirStr[i]) == c {
//				counter++
//			}
//		}
//		return counter
//	}
//	return 0
//}
//
//func last3Dir(dirStr string) string {
//	if len(dirStr) < 1 {
//		return dirStr
//	}
//	return dirStr[len(dirStr)-1:]
//}
//
//func areDirSame(dirStr string) bool {
//	if len(dirStr) > 3 {
//		result := string(dirStr[len(dirStr)-1]) == string(dirStr[len(dirStr)-2]) && string(dirStr[len(dirStr)-2]) == string(dirStr[len(dirStr)-3]) && string(dirStr[len(dirStr)-3]) == string(dirStr[len(dirStr)-4])
//		return result
//	}
//	return false
//}
//
//func opposingPath(p string) string {
//	if p == "U" {
//		return "D"
//	}
//	if p == "D" {
//		return "U"
//	}
//	if p == "L" {
//		return "R"
//	}
//	if p == "R" {
//		return "L"
//	}
//	return ""
//}
