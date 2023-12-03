package com.aoc;

import java.lang.reflect.Array;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Day8 extends AdventOfCode {

    @Override
    public void solve() {
        Scanner in = super.input;
        String[] inputLines = new String[200];
        for(int i = 0; in.hasNextLine(); i++)
            inputLines[i] = in.nextLine();

        int sumOfAllUniqueNumbers = 0;

        for (String inputLine: inputLines) {
            String leftSide = inputLine.split(" \\| ")[0];
            String rightSide = inputLine.split(" \\| ")[1];
            String[] allDigits = leftSide.split(" ");
            String[] uniqueDigits = rightSide.split(" ");

            String uniqueNumber = getUniqueDigitsNumber(allDigits, uniqueDigits);
            sumOfAllUniqueNumbers += Integer.parseInt(uniqueNumber);
            System.out.println(uniqueNumber);
        }
        System.out.println("Answer: " + sumOfAllUniqueNumbers);
    }

    private String getUniqueDigitsNumber(String[] allDigits, String[] uniqueDigits) {
        String[] indexedDigits = new String[10];
        solveEasyDigits(allDigits, indexedDigits);
        solveDigit3(allDigits,indexedDigits);
        solveDigit069(allDigits, indexedDigits);
        solve2and5(allDigits, indexedDigits);
        return findUniqueDigits(indexedDigits, uniqueDigits);
    }

    private String findUniqueDigits(String[] indexedDigits, String[] uniqueDigits) {
        String finalResult = "";
        for(String uniqueDigit: uniqueDigits) {
            for(int i = 0; i < indexedDigits.length; i++) {
                String common = commonChars(uniqueDigit, indexedDigits[i]);
                if(common.length() == uniqueDigit.length() && common.length() == indexedDigits[i].length())
                    finalResult += i;
            }
        }
        return finalResult;
    }

    private void solveDigit3(String[] allDigits, String[] indexedDigits) {
        String commonBetween7and1 = commonChars(indexedDigits[1], indexedDigits[7]);
        String digit3 = Arrays.stream(allDigits).filter(digit -> digit.length() == 5 & (commonChars(digit, commonBetween7and1).length() == commonBetween7and1.length())).findFirst().get();
        indexedDigits[3] = digit3;
    }

    private void solveDigit069(String[] allDigits, String[] indexedDigits) {
        List<String> potential2and5 = getPotential2and5(allDigits, indexedDigits);
        String x = commonChars(potential2and5.get(0), potential2and5.get(1));
        String y = indexedDigits[1];
        String digit9 = Arrays.stream(allDigits).filter(digit -> firstContainsSecond(digit, x) && firstContainsSecond(digit, y) && digit.length() == 6).findFirst().get();
        String digit6 = Arrays.stream(allDigits).filter(digit -> firstContainsSecond(digit, x) && !firstContainsSecond(digit, y) && digit.length() == 6).findFirst().get();
        String digit0 = Arrays.stream(allDigits).filter(digit -> !firstContainsSecond(digit, x) && firstContainsSecond(digit, y) && digit.length() == 6).findFirst().get();
        indexedDigits[0] = digit0;
        indexedDigits[6] = digit6;
        indexedDigits[9] = digit9;
    }

    private boolean firstContainsSecond(String first, String second) {
        return commonChars(first, second).length() == second.length();
    }

    private List<String> getPotential2and5(String[] allDigits, String[] indexedDigits) {
        return Arrays.stream(allDigits).filter(digit -> digit.length() == 5 && !digit.equals(indexedDigits[3])).collect(Collectors.toList());
    }

    private void solve2and5(String[] allDigits, String[] indexedDigits) {
        String digit5 = Arrays.stream(allDigits).filter(digit -> digit.length() == 5 && firstContainsSecond(indexedDigits[9], digit) && !digit.equals(indexedDigits[3])).findFirst().get();
        String digit2 = Arrays.stream(allDigits).filter(digit -> digit.length() == 5 && !digit.equals(digit5) && !digit.equals(indexedDigits[3])).findFirst().get();
        indexedDigits[2] = digit2;
        indexedDigits[5] = digit5;
    }

    private String commonChars(String str1, String str2) {

        if (str1.length() > 0 & str2.length() > 0) {
            List<Character> s1 = new ArrayList<>();
            List<Character> s2 = new ArrayList<>();

            for (int i = 0; i < str1.length(); i++) {
                s1.add(str1.charAt(i));
            }

            for (int i = 0; i < str2.length(); i++) {
                s2.add(str2.charAt(i));
            }

            // Finding intersection of both lists
            s1.retainAll(s2);

            StringBuilder sb = new StringBuilder();

            for (Character c : s1) {
                sb.append(c);
            }

            return sb.toString();
        } else
            return "";
    }

    private void solveEasyDigits(String[] allDigits, String[] indexedDigits) {
        for(String digit: allDigits) {
            int digitLength = digit.length();
            switch (digitLength) {
                case 2 -> indexedDigits[1] = digit;
                case 4 -> indexedDigits[4] = digit;
                case 3 -> indexedDigits[7] = digit;
                case 7 -> indexedDigits[8] = digit;
            }
        }
    }
}
