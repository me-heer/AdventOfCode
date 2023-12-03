package com.aoc;

import java.util.ArrayList;
import java.util.Scanner;

public class Day1 extends AdventOfCode {

    public static int countSlidingWindow(ArrayList<Integer> input) {
        int result = 0;
        int prevWindowValue = Integer.MAX_VALUE;
        for (int i = 2; i < input.size(); i++) {
            int currentWindowValue = input.get(i) + input.get(i - 1) + input.get(i - 2);
            if (currentWindowValue > prevWindowValue)
                result++;
            prevWindowValue = currentWindowValue;
        }
        return result;
    }

    @Override
    public void solve() {
        Scanner in = super.input;
        ArrayList<Integer> input = new ArrayList<>();
        while (in.hasNextLine()) {
            input.add(in.nextInt());
        }

        int result = countSlidingWindow(input);
        System.out.println("Result: " + result);
    }
}
