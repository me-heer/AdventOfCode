package com.aoc;

import java.util.Scanner;

public class Day11 extends AdventOfCode {
    private static final int SIZE = 10;
    Octopus[][] matrix = new Octopus[SIZE][SIZE];
    private static int totalFlashes = 0;

    @Override
    public void solve() {
        Scanner in = input;

        int row = 0;
        while (in.hasNextLine()) {
            String line = in.nextLine();
            for(int i = 0; i < line.length(); i++) {
                Octopus octopus = new Octopus(Integer.parseInt(String.valueOf(line.charAt(i))));
                matrix[row][i] = octopus;
            }
            row++;
        }

        int iteration = 0;
        while(true) {
            System.out.println(iteration++);
            simulate();
            //print();
        }
    }


    private void print() {
        for(int i = 0; i < SIZE; i++) {
            for(int j = 0; j < SIZE; j++)
                System.out.format("%3d", matrix[i][j].energyLevel);
            System.out.println();
        }

    }

    private void simulate() {
        for(int i = 0; i < SIZE; i++) {
            for(int j = 0; j < SIZE; j++) {
                Octopus octopus = matrix[i][j];
                if(!octopus.isFlashed) {
                    octopus.energyLevel++;
                    if(octopus.energyLevel > 9) {
                        flashOctopus(i, j);
                    }
                }
            }
        }

        checkForAllFlash();

        for(int i = 0; i < SIZE; i++) {
            for(int j = 0; j < SIZE; j++)
                if(matrix[i][j].isFlashed)
                {
                    matrix[i][j].isFlashed = false;
                    matrix[i][j].energyLevel = 0;
                }
        }
    }

    private void checkForAllFlash() {
        int flashCounter = 0;
        for(int i = 0; i < SIZE; i++)
            for(int j = 0; j < SIZE; j++) {
                if(matrix[i][j].isFlashed == true)
                    flashCounter++;
            }
        if(flashCounter == SIZE * SIZE)
        {
            System.out.println("ALL FLASH");
            System.exit(0);
        }
    }

    private void flashOctopus(int row, int column) {
        matrix[row][column].isFlashed = true;
        totalFlashes++;
        illuminateAround(row, column);
    }

    private void illuminateAround(int row, int column) {
        scanLeft(row, column);
        scanUpperLeft(row, column);
        scanUp(row,column);
        scanUpperRight(row,column);
        scanRight(row,column);
        scanLowerRight(row,column);
        scanDown(row,column);
        scanLowerLeft(row,column);
    }

    private void scanLowerLeft(int row, int column) {
        if(column != 0 && row != SIZE - 1) {
            Octopus octopus = matrix[row + 1][column - 1];
            if(!octopus.isFlashed) {
                octopus.energyLevel++;
                if(octopus.energyLevel > 9) {
                    flashOctopus(row + 1, column - 1);
                }
            }
        }
    }
    private void scanDown(int row, int column) {
        if(row != SIZE - 1) {
            Octopus octopus = matrix[row + 1][column];
            if(!octopus.isFlashed) {
                octopus.energyLevel++;
                if(octopus.energyLevel > 9) {
                    flashOctopus(row + 1, column);
                }
            }
        }
    }
    private void scanLowerRight(int row, int column) {
        if(row != SIZE - 1 && column != SIZE - 1) {
            Octopus octopus = matrix[row + 1][column + 1];
            if(!octopus.isFlashed) {
                octopus.energyLevel++;
                if(octopus.energyLevel > 9) {
                    flashOctopus(row + 1, column + 1);
                }
            }
        }
    }
    private void scanRight(int row, int column) {
        if(column != SIZE - 1) {
            Octopus octopus = matrix[row][column + 1];
            if(!octopus.isFlashed) {
                octopus.energyLevel++;
                if(octopus.energyLevel > 9) {
                    flashOctopus(row, column + 1);
                }
            }
        }
    }
    private void scanUpperRight(int row, int column) {
        if(row != 0 && column != SIZE - 1) {
            Octopus octopus = matrix[row - 1][column + 1];
            if(!octopus.isFlashed) {
                octopus.energyLevel++;
                if(octopus.energyLevel > 9) {
                    flashOctopus(row - 1, column + 1);
                }
            }
        }
    }
    private void scanUp(int row, int column) {
        if(row != 0) {
            Octopus octopus = matrix[row - 1][column];
            if(!octopus.isFlashed) {
                octopus.energyLevel++;
                if(octopus.energyLevel > 9) {
                    flashOctopus(row - 1, column);
                }
            }
        }
    }
    private void scanUpperLeft(int row, int column) {
        if(row != 0 && column != 0) {
            Octopus octopus = matrix[row - 1][column - 1];
            if(!octopus.isFlashed) {
                octopus.energyLevel++;
                if(octopus.energyLevel > 9) {
                    flashOctopus(row - 1, column - 1);
                }
            }
        }
    }

    private void scanLeft(int row, int column) {
        if(column != 0) {
            Octopus octopus = matrix[row][column - 1];
            if(!octopus.isFlashed) {
                octopus.energyLevel++;
                if(octopus.energyLevel > 9) {
                    flashOctopus(row, column - 1);
                }
            }
        }
    }

    private class Octopus {
        boolean isFlashed = false;
        int energyLevel = 0;

        Octopus(int e) {
            this.energyLevel = e;
        }
    }
}
