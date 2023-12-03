package com.aoc;

import java.lang.reflect.Array;
import java.util.*;
import java.util.stream.Collectors;

public class Day9 extends AdventOfCode {

    ArrayList<ArrayList<Position>> matrix = new ArrayList<>();
    ArrayList<LowPoint> lowPoints = new ArrayList<>();
    Stack<Position> buffer = new Stack<>();
    ArrayList<Basin> basins = new ArrayList<>();

    @Override
    public void solve() {
        Scanner in = super.input;
        ArrayList<String> input = new ArrayList<>();
        while (in.hasNextLine()) {
            input.add(in.nextLine());
        }

        ArrayList<LowPoint> lowPoints = new ArrayList<>();
        for (int i = 0; i < input.size(); i++) {
            for (int j = 0; j < input.get(i).length(); j++) {
                int current = Integer.parseInt(String.valueOf(input.get(i).charAt(j)));
                int down, up, left, right;
                if (i == 0 && j == 0) {
                    right = Integer.parseInt(String.valueOf(input.get(i).charAt(j + 1)));
                    down = Integer.parseInt(String.valueOf(input.get(i + 1).charAt(j)));
                    if (current < right && current < down) {
                        lowPoints.add(new LowPoint(i, j, current));
                    }
                } else if (i == 0 && j == input.get(i).length() - 1) {
                    down = Integer.parseInt(String.valueOf(input.get(i + 1).charAt(j)));
                    left = Integer.parseInt(String.valueOf(input.get(i).charAt(j - 1)));
                    if (current < left && current < down)
                        lowPoints.add(new LowPoint(i, j, current));
                } else if (i == input.size() - 1 && j == 0) {
                    up = Integer.parseInt(String.valueOf(input.get(i - 1).charAt(j)));
                    right = Integer.parseInt(String.valueOf(input.get(i).charAt(j + 1)));
                    if (current < up && current < right)
                        lowPoints.add(new LowPoint(i, j, current));
                } else if (i == input.size() - 1 && j == input.get(i).length() - 1) {
                    left = Integer.parseInt(String.valueOf(input.get(i).charAt(j - 1)));
                    up = Integer.parseInt(String.valueOf(input.get(i - 1).charAt(j)));
                    if (current < left && current < up)
                        lowPoints.add(new LowPoint(i, j, current));
                } else {
                    down = i != input.size() - 1 ? Integer.parseInt(String.valueOf(input.get(i + 1).charAt(j))) : Integer.MAX_VALUE;
                    up = i != 0 ? Integer.parseInt(String.valueOf(input.get(i - 1).charAt(j))) : Integer.MAX_VALUE;
                    left = j != 0 ? Integer.parseInt(String.valueOf(input.get(i).charAt(j - 1))) : Integer.MAX_VALUE;
                    right = j != input.get(i).length() - 1 ? Integer.parseInt(String.valueOf(input.get(i).charAt(j + 1))) : Integer.MAX_VALUE;
                    if (current < left && current < down && current < right && current < up)
                        lowPoints.add(new LowPoint(i, j, current));
                }
            }
        }

        int answer = 0;
        for (LowPoint lowPoint : lowPoints)
            answer += (lowPoint.value + 1);

        System.out.println(answer);

        int i = 0;
        for (String inputLine : input) {
            ArrayList<Position> row = new ArrayList<>();
            int j = 0;
            for (char c : inputLine.toCharArray())
                row.add(new Position(i, j++, Integer.parseInt(String.valueOf(c))));
            matrix.add(row);
            i++;
        }

        for (LowPoint lowPoint : lowPoints)
            search(lowPoint);

        ArrayList<Integer> basinSizes = new ArrayList<>();
        for(Basin basin: basins) {
            basinSizes.add(basin.size);
        }
        List<Integer> basinSizesSorted = basinSizes.stream().sorted(Comparator.reverseOrder()).collect(Collectors.toList());
        System.out.println(basinSizesSorted.get(0) * basinSizesSorted.get(1) * basinSizesSorted.get(2));
    }

    private void search(LowPoint lowPoint) {
        System.out.println("BEGIN");
        Basin basin = new Basin();
        basin.sum += matrix.get(lowPoint.row).get(lowPoint.column).value;
        basin.size++;
        System.out.println(matrix.get(lowPoint.row).get(lowPoint.column).value);
        matrix.get(lowPoint.row).get(lowPoint.column).isMarked = true;
        scanDirections(lowPoint.row, lowPoint.column, basin);
        while(!buffer.isEmpty()) {
            Position position = buffer.pop();
            scanDirections(position.row, position.column, basin);
        }
        basins.add(basin);
        System.out.println("END");
    }

    private void scanDirections(int row, int col, Basin basin) {
        if(col != matrix.get(row).size() - 1)
            scanRight(row, col, basin);
        if(row != matrix.size() - 1)
            scanDown(row, col, basin);
        if(col != 0)
            scanLeft(row, col, basin);
        if(row != 0)
            scanUp(row, col, basin);
    }

    private void scanUp(int row, int col, Basin basin) {
        Position up = matrix.get(row - 1).get(col);
        if (!up.isMarked && up.value != 9)
        {
            System.out.println(up.value);
            basin.sum += up.value;
            basin.size++;
            up.isMarked = true;
            buffer.push(up);
        }
    }

    private void scanLeft(int row, int col, Basin basin) {
        Position left = matrix.get(row).get(col - 1);
        if (!left.isMarked && left.value != 9)
        {
            basin.sum += left.value;
            basin.size++;
            System.out.println(left.value);
            left.isMarked = true;
            buffer.push(left);
        }
    }

    private void scanDown(int row, int col, Basin basin) {
        Position down = matrix.get(row + 1).get(col);
        if (!down.isMarked && down.value != 9)
        {
            basin.sum += down.value;
            basin.size++;
            System.out.println(down.value);
            down.isMarked = true;
            buffer.push(down);
        }
    }

    private void scanRight(int row, int col, Basin basin) {
        Position right = matrix.get(row).get(col + 1);
        if (!right.isMarked && right.value != 9) {
            basin.sum += right.value;
            basin.size++;
            System.out.println(right.value);
            right.isMarked = true;
            buffer.push(right);
        }
    }

    private class Basin {
        int size = 0;
        int sum = 0;
    }

    private class Position {
        boolean isMarked;
        int row;
        int column;
        int value;

        Position(int row, int column, int value) {
            this.row = row;
            this.column = column;
            this.value = value;
        }

    }

    private class LowPoint {
        int row;
        int column;
        int value;

        LowPoint(int r, int c, int v) {
            row = r;
            column = c;
            value = v;
        }
    }
}
