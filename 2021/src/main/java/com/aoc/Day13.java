package com.aoc;

import java.util.ArrayList;
import java.util.Collections;
import java.util.Scanner;

public class Day13 extends AdventOfCode {
    ArrayList<Operation> foldOperations = new ArrayList<>();
    ArrayList<ArrayList<Point>> paper = new ArrayList<>();
    ArrayList<Integer> xCoords = new ArrayList<>();
    ArrayList<Integer> yCoords = new ArrayList<>();
    int rowBoundary, colBoundary, previousRowBoundary, previousColBoundary;

    @Override
    public void solve() {
        Scanner in = input;
        while (in.hasNextLine()) {
            String line = in.nextLine();
            String[] coords = line.split(",");
            if (coords.length == 2) {
                createCoordinate(xCoords, yCoords, coords);
            } else if (coords.length == 1 && !coords[0].equals("")) {
                createFoldOperation(line);
            }
        }
        populateMatrix(xCoords, yCoords);
        markDots(xCoords, yCoords);
        System.out.println("\n Executing Fold Operations... \n");
        for (Operation operation: foldOperations)
            executeOperations(operation);
        printPaper();
        System.out.println("Total Dots: " + getTotalDots(paper));
    }

    private int getTotalDots(ArrayList<ArrayList<Point>> paper) {
        int totalDots = 0;
        for (int rowIndex = 0; rowIndex < rowBoundary; rowIndex++) {
            for (int colIndex = 0; colIndex < colBoundary; colIndex++) {
                Point point = paper.get(rowIndex).get(colIndex);
                if (point.isDot)
                    totalDots++;
            }
        }
        return totalDots;
    }

    private void executeOperations(Operation operation) {
        String axis = operation.axis;
        Integer foldLine = operation.value;
        if (axis.equals("y")) {
            foldUp(foldLine);
        } else {
            foldLeft(foldLine);
        }
    }

    private void foldLeft(Integer foldLine) {
        previousColBoundary = colBoundary;
        colBoundary = foldLine;
        for (int colIndex = colBoundary + 1; colIndex < previousColBoundary; colIndex++) {
            int colIndexToBeOverwritten = foldLine - (colIndex - colBoundary);
            for (int rowIndex = 0; rowIndex < rowBoundary; rowIndex++) {
                boolean wasPreviouslyDot = paper.get(rowIndex).get(colIndexToBeOverwritten).isDot;
                Point pointToBeOverwritten = paper.get(rowIndex).get(colIndexToBeOverwritten);
                Point pointToOverwriteFrom = paper.get(rowIndex).get(colIndex);
                pointToBeOverwritten.isDot = wasPreviouslyDot ? pointToBeOverwritten.isDot :
                        pointToOverwriteFrom.isDot;
            }
        }
    }

    private void foldUp(Integer foldLine) {
        previousRowBoundary = rowBoundary;
        rowBoundary = foldLine;
        for (int rowIndex = rowBoundary + 1; rowIndex < previousRowBoundary; rowIndex++) {
            int rowIndexToBeOverwritten = foldLine - (rowIndex - rowBoundary);
            ArrayList<Point> rowToOverwriteFrom = paper.get(rowIndex);
            ArrayList<Point> rowToBeOverwritten = paper.get(rowIndexToBeOverwritten);
            for (int colIndex = 0; colIndex < colBoundary; colIndex++) {
                boolean wasPreviouslyDot = rowToBeOverwritten.get(colIndex).isDot;
                rowToBeOverwritten.get(colIndex).isDot = wasPreviouslyDot ?
                        rowToBeOverwritten.get(colIndex).isDot : rowToOverwriteFrom.get(colIndex).isDot;
            }
        }
    }

    private void printPaper() {
        for (int rowIndex = 0; rowIndex < rowBoundary; rowIndex++) {
            for (int colIndex = 0; colIndex < colBoundary; colIndex++) {
                Point point = paper.get(rowIndex).get(colIndex);
                if (point.isDot)
                    System.out.printf("%2s", "#");
                else
                    System.out.printf("%2s", ".");
            }
            System.out.println();
        }
    }

    private void createFoldOperation(String line) {
        String[] operation = line.split(" ");
        String[] operationStr = operation[operation.length - 1].split("=");
        Operation foldOp = new Operation(operationStr[0], Integer.parseInt(operationStr[1]));
        foldOperations.add(foldOp);
    }

    private void createCoordinate(ArrayList<Integer> xCoords, ArrayList<Integer> yCoords, String[] coords) {
        xCoords.add(Integer.parseInt(coords[0]));
        yCoords.add(Integer.parseInt(coords[1]));
    }

    private void populateMatrix(ArrayList<Integer> xCoords, ArrayList<Integer> yCoords) {
        Integer maxRows = Collections.max(yCoords) + 1;
        Integer maxCols = Collections.max(xCoords) + 1;
        colBoundary = maxCols;
        rowBoundary = maxRows;
        for (int i = 0; i < maxRows; i++) {
            ArrayList<Point> row = new ArrayList<>();
            for (int j = 0; j < maxCols; j++) {
                boolean isDot = xCoords.contains(j) && yCoords.contains(i);
                row.add(new Point(false));
            }
            paper.add(row);
        }
    }

    private void markDots(ArrayList<Integer> xCoords, ArrayList<Integer> yCoords) {
        for (int i = 0; i < xCoords.size(); i++) {
            int x = xCoords.get(i);
            int y = yCoords.get(i);
            paper.get(y).get(x).isDot = true;
        }
    }

    class Point {
        boolean isDot;

        public Point(boolean isDot) {
            this.isDot = isDot;
        }
    }

    class Operation {
        String axis;
        int value;

        public Operation(String axis, int value) {
            this.axis = axis;
            this.value = value;
        }
    }
}
