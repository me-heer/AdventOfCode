package com.aoc;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;

public class Day7 extends AdventOfCode {
    @Override
    public void solve() {
        Scanner in = super.input;
        String[] crabPositionsStr = in.nextLine().split(",");
        int[] crabPositions = new int[crabPositionsStr.length];
        int index = 0;
        for(String crabPosition: crabPositionsStr)
            crabPositions[index++] = Integer.parseInt(crabPosition);

        int min = Arrays.stream(crabPositions).min().getAsInt();
        int max = Arrays.stream(crabPositions).max().getAsInt();
        int[] fuelSpent = new int[max - min];
        for(int i = min; i < max; i++) {
            int fuel = 0;
            for(int j = 0; j < crabPositions.length; j++)
                fuel += (Math.abs(crabPositions[j] - i));
            fuelSpent[i - min] = fuel;
        }

        int sum = 0;
        long answer = 0;
        for(int i = 0; i < crabPositions.length;i++) {
            sum += crabPositions[i];
            answer += sumOfDigits(Math.abs(crabPositions[i] - 478));
        }
        float avg = sum/crabPositions.length;
        System.out.println(Arrays.stream(fuelSpent).min().getAsInt());
        System.out.println("Average: " + avg);
        System.out.println("Answer: " + answer);
    }

    private long sumOfDigits(long n) {
        return (n * (n + 1))/2;
    }

}
