package com.aoc;

import java.util.Scanner;

public class Day5 extends AdventOfCode {

    public static final int MATRIX_SIZE = 1000;

    @Override
    public void solve() {
        Scanner in = super.input;
        Floor floor = new Floor();
        while (in.hasNextLine()) {
            String line = in.nextLine();
            floor.readLine(line);
        }
        floor.printFloor();
        System.out.println("Points where at least two lines overlap: " + floor.getPointsWhereAtLeast2LinesOverlap());
    }


    private class Floor {
        Point[][] matrix = new Point[MATRIX_SIZE][MATRIX_SIZE];

        Floor() {
            for (int i = 0; i < MATRIX_SIZE; i++)
                for (int j = 0; j < MATRIX_SIZE; j++)
                    matrix[i][j] = new Point();
        }

        public void readLine(String line) {
            String[] coordinates = line.split(" -> ");

            int x1 = Integer.parseInt(coordinates[0].split(",")[0]);
            int y1 = Integer.parseInt(coordinates[0].split(",")[1]);
            int x2 = Integer.parseInt(coordinates[1].split(",")[0]);
            int y2 = Integer.parseInt(coordinates[1].split(",")[1]);

            Coordinate coord1 = new Coordinate(x1, y1);
            Coordinate coord2 = new Coordinate(x2, y2);
            coverLine(coord1, coord2);
        }

        public void coverLine(Coordinate coordinate1, Coordinate coordinate2) {
            if (isLineHorizontal(coordinate1, coordinate2)) {
                int y = coordinate1.y;
                for (int i = Math.min(coordinate1.x, coordinate2.x); i <= Math.max(coordinate1.x, coordinate2.x); i++) {
                    matrix[y][i].linesCovered++;
                }
            } else if (isLineVertical(coordinate1, coordinate2)) {
                int x = coordinate1.x;
                for (int i = Math.min(coordinate1.y, coordinate2.y); i <= Math.max(coordinate1.y, coordinate2.y); i++) {
                    matrix[i][x].linesCovered++;
                }
            } else if (isLineDiagonal(coordinate1, coordinate2)) {
                int distance = Math.abs(coordinate1.x - coordinate2.x);

                int xDiagonalDirection = -(coordinate1.x - coordinate2.x) / distance;
                int yDiagonalDirection = -(coordinate1.y - coordinate2.y) / distance;
                int xInitial = coordinate1.x;
                int yInitial = coordinate1.y;

                for (int i = 0; i <= distance; i++) {
                    matrix[yInitial][xInitial].linesCovered++;
                    yInitial += yDiagonalDirection;
                    xInitial += xDiagonalDirection;
                }
            }
        }

        public void printFloor() {
            for (int i = 0; i < MATRIX_SIZE; i++) {
                for (int j = 0; j < MATRIX_SIZE; j++) {
                    System.out.format("%5d", matrix[i][j].linesCovered);
                }
                System.out.println();
            }
        }

        private int getPointsWhereAtLeast2LinesOverlap() {
            int count = 0;
            for (int i = 0; i < MATRIX_SIZE; i++)
                for (int j = 0; j < MATRIX_SIZE; j++)
                    if (matrix[i][j].linesCovered >= 2)
                        count++;
            return count;
        }

        private boolean isLineHorizontal(Coordinate coordinate1, Coordinate coordinate2) {
            return coordinate1.y == coordinate2.y;
        }

        private boolean isLineVertical(Coordinate coordinate1, Coordinate coordinate2) {
            return coordinate1.x == coordinate2.x;
        }

        private boolean isLineDiagonal(Coordinate coordinate1, Coordinate coordinate2) {
            return Math.abs(coordinate1.x - coordinate2.x) == Math.abs(coordinate1.y - coordinate2.y);
        }
    }

    private class Coordinate {
        int x;
        int y;

        Coordinate(int x, int y) {
            this.x = x;
            this.y = y;
        }
    }

    private class Point {
        int linesCovered = 0;
    }

}
