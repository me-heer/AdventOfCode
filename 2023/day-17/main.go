package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

var graph [][]Point

var dist map[PointLocation]DirectedDistance

type DirectedDistance struct {
	dist int
	path string
}

type PointLocation struct {
	i int
	j int
}

type Point struct {
	location     PointLocation
	heat         int
	neighbors    []*Point
	neighborDirs []string
	dist         int
	dirHistory   string
}

func main() {
	input, _ := os.Open("day-17/input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var matrix []string

	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}

	graph = make([][]Point, len(matrix))

	for i := 0; i < len(matrix); i++ {
		row := make([]Point, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			heatLoss, _ := strconv.Atoi(string(matrix[i][j]))
			row[j] = Point{
				location: PointLocation{
					i: i,
					j: j,
				}, heat: heatLoss}
		}
		graph[i] = row
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			p := &graph[i][j]
			if i > 0 {
				// add up neighbour to current
				p.neighbors = append(p.neighbors, &graph[i-1][j])
				p.neighborDirs = append(p.neighborDirs, "U")
			}
			if j > 0 {
				// left
				p.neighbors = append(p.neighbors, &graph[i][j-1])
				p.neighborDirs = append(p.neighborDirs, "L")
			}
			if i < len(matrix)-1 {
				// down
				p.neighbors = append(p.neighbors, &graph[i+1][j])
				p.neighborDirs = append(p.neighborDirs, "D")
			}
			if j < len(matrix[0])-1 {
				// right
				p.neighbors = append(p.neighbors, &graph[i][j+1])
				p.neighborDirs = append(p.neighborDirs, "R")
			}
		}
	}
	dist = make(map[PointLocation]DirectedDistance)
	djikstra(graph, graph[0][0])
	p := dist[PointLocation{i: 1, j: 8}]
	println(p.path)
	println(p.dist)

}

var queue []Point

func djikstra(graph [][]Point, source Point) {
	dist[source.location] = DirectedDistance{
		dist: 0,
		path: "START ",
	}
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[0]); j++ {
			if graph[i][j].location != source.location {
				dist[graph[i][j].location] = DirectedDistance{
					dist: math.MaxInt,
					path: "",
				}
			}
			queue = append(queue, graph[i][j])
		}
	}

	for len(queue) > 0 {
		minDist := math.MaxInt
		var minV Point
		var minVIndex int
		for i, v := range queue {
			if dist[v.location].dist < minDist {
				minDist = dist[v.location].dist
				minV = v
				minVIndex = i
			}
		}

		queue = append(queue[:minVIndex], queue[minVIndex+1:]...)

		minV = graph[minV.location.i][minV.location.j]

		for i, neighbour := range minV.neighbors {
			distance := dist[minV.location].dist + neighbour.heat
			direction := dist[minV.location].path + minV.neighborDirs[i]
			// TODO: Check if direction is valid
			if distance < dist[neighbour.location].dist {
				dist[neighbour.location] = DirectedDistance{dist: distance, path: direction}
			}
		}
	}
}

func areDirSame(dirStr string) bool {
	if len(dirStr) > 3 {
		result := string(dirStr[len(dirStr)-1]) == string(dirStr[len(dirStr)-2]) && string(dirStr[len(dirStr)-2]) == string(dirStr[len(dirStr)-3]) && string(dirStr[len(dirStr)-3]) == string(dirStr[len(dirStr)-4])
		return result
	}
	return false
}

func opposingPath(p string) string {
	if p == "U" {
		return "D"
	}
	if p == "D" {
		return "U"
	}
	if p == "L" {
		return "R"
	}
	if p == "R" {
		return "L"
	}
	return ""
}
