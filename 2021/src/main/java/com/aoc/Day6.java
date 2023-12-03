package com.aoc;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.Scanner;

public class Day6 extends AdventOfCode {
    @Override
    public void solve() {
        Scanner in = super.input;
        long[] fishes = new long[10];
        while (in.hasNextLine()) {
            String[] fishesStr = in.nextLine().split(",");
            for(String fishStr : fishesStr) {
                fishes[Integer.parseInt(fishStr)]++;
            }
        }
        int daysToSimulate = 256;
        for(int i = 0; i < daysToSimulate; i++) {
            fishes[7] += fishes[0];
            fishes[9] = fishes[0];
            for(int j = 0; j < 9; j++)
                fishes[j] = fishes[j + 1];
            fishes[9] = 0;
        }

        long sum = 0;
        for(long fish: fishes)
            sum += fish;
        System.out.println("Answer: " + sum);
    }
}
