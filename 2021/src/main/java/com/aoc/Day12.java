package com.aoc;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Day12 extends AdventOfCode {
    ArrayList<Vertex> graph = new ArrayList<>();
    ArrayList<String> pathsTaken = new ArrayList<>();
    int totalPaths = 0;

    @Override
    public void solve() {
        Scanner in = super.input;
        while (in.hasNextLine()) {
            String line = in.nextLine();
            String[] vertices = line.split("-");
            Vertex fromNode =
                    graph.stream().filter(vertex -> vertex.key.equals(vertices[0])).findFirst().orElse(new Vertex(vertices[0], new ArrayList<>()));
            Vertex toNode =
                    graph.stream().filter(vertex -> vertex.key.equals(vertices[1])).findFirst().orElse(new Vertex(vertices[1], new ArrayList<>()));
            fromNode.adjacentVertices.add(toNode);
            toNode.adjacentVertices.add(fromNode);

            if (!graph.stream().anyMatch(vertex -> vertex.key.equals(fromNode.key))) {
                graph.add(fromNode);
            }
            if (!graph.stream().anyMatch(vertex -> vertex.key.equals(toNode.key))) {
                graph.add(toNode);
            }
        }

        Vertex startNode = graph.stream().filter(vertex -> vertex.key.equals("start")).findFirst().get();
//        visit(startNode, new LinkedList<>(), null);

        List<Vertex> smallCaveNodes = graph.stream().filter(vertex -> isSmallCave(vertex)).collect(Collectors.toList());
        for (Vertex smallCaveNode: smallCaveNodes) {
            visit(startNode, new LinkedList<>(), smallCaveNode);
        }
        List<String> distinctPaths = pathsTaken.stream().distinct().collect(Collectors.toList());
        printPaths(distinctPaths);
        System.out.println("Total Paths: " + distinctPaths.size());
    }

    public void visit(Vertex vertex, LinkedList<Vertex> currentVisitedPath, Vertex useTwice) {
        if (vertex.key.equals("end")) {
            currentVisitedPath.add(vertex);
            printPath(currentVisitedPath);
            totalPaths++;
            currentVisitedPath.removeLastOccurrence(vertex);
            return;
        }
        currentVisitedPath.add(vertex);
        List<Vertex> adjacentVertices = vertex.adjacentVertices;
        for (Vertex adjacentVertex: adjacentVertices) {
            if (!adjacentVertex.key.equals("start")
                    &&
                    !(isSmallCave(adjacentVertex)
                                    && visitedAtLeastOnce(adjacentVertex, currentVisitedPath)
                                    && !(allowedToUseTwice(adjacentVertex, useTwice) && visitedOnlyOnce(adjacentVertex, currentVisitedPath)))) {
                visit(adjacentVertex, currentVisitedPath, useTwice);
            }
        }
        currentVisitedPath.removeLastOccurrence(vertex);
    }

    private boolean isSmallCave(Vertex vertex) {
        return !vertex.key.equals("start") && !vertex.key.equals("end") && vertex.key.matches(".*[a-z].*");
    }

    private boolean allowedToUseTwice(Vertex current, Vertex toBeUsedTwice) {
        return current.equals(toBeUsedTwice);
    }

    private boolean visitedAtLeastOnce(Vertex vertex, LinkedList<Vertex> currentVisitedPath) {
        List<Vertex> numberOfVisits = currentVisitedPath.stream().filter(filterVertex -> filterVertex.key.equals(vertex.key)).collect(Collectors.toList());
        return numberOfVisits.size() >= 1;
    }

    private boolean visitedOnlyOnce(Vertex vertex, LinkedList<Vertex> currentVisitedPath) {
        List<Vertex> numberOfVisits = currentVisitedPath.stream().filter(filterVertex -> filterVertex.key.equals(vertex.key)).collect(Collectors.toList());
        return numberOfVisits.size() < 2;
    }

    void printPath(List<Vertex> visitedVertices) {
        String pathString = "";
        for (Vertex vertex : visitedVertices) {
//            System.out.print(vertex.key + ",");
            pathString += vertex.key + ",";
        }
        pathString = pathString.substring(0, pathString.length() - 1);
        pathsTaken.add(pathString);
//        System.out.print("\b");
//        System.out.println();
    }

    void printPaths(List<String> paths) {
        for (String path: paths)
            System.out.println(path);
    }

    class Vertex {
        public String key;
        public List<Vertex> adjacentVertices;

        public Vertex(String key, List<Vertex> adjacentVertices) {
            this.key = key;
            this.adjacentVertices = adjacentVertices;
        }
    }
}
