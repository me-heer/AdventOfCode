package com.aoc;

import java.util.ArrayList;
import java.util.Objects;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Day3 extends AdventOfCode {

    private final Byte EQUAL_COUNT = 2;

    @Override
    public void solve() {
        Scanner in = super.input;
        ArrayList<ArrayList<Byte>> input = new ArrayList<>();
        while (in.hasNextLine()) {
            ArrayList<Byte> row = new ArrayList<>();
            String inputStr = in.nextLine();
            for (int i = 0; i < inputStr.length(); i++) {
                row.add(Byte.valueOf(inputStr.substring(i, i + 1)));
            }
            input.add(row);
        }

        int rowLength = input.get(0).size();
        StringBuilder gammaRate = new StringBuilder();
        StringBuilder epsilonRate = new StringBuilder();
        for (int index = 0; index < rowLength; index++) {
            String mostCommon = findMostCommonAtIndex(input, index);
            gammaRate.append(mostCommon);
            if (Objects.equals(mostCommon, "0")) {
                epsilonRate.append("1");
            } else {
                epsilonRate.append("0");
            }
        }

        Long gammaRateCount = Long.parseLong(gammaRate.toString(), 2);
        Long epsilonRateCount = Long.parseLong(epsilonRate.toString(), 2);
        System.out.println("Gamma Rate: " + gammaRateCount);
        System.out.println("Epsilon Rate: " + epsilonRateCount);
        System.out.println("Gamma * Epsilon: " + gammaRateCount * epsilonRateCount);

        //--------------------------------Part 2------------------------------------
        String oxygenGeneratorRatingStr = findOxygenGeneratorRating(input, rowLength);
        String CO2ScrubberRatingStr = findCO2ScrubberRating(input, rowLength);
        Long oxygenGeneratorRating = Long.parseLong(oxygenGeneratorRatingStr, 2);
        Long CO2ScrubberRating = Long.parseLong(CO2ScrubberRatingStr, 2);
        System.out.println("Oxygen Generator Rating: " + oxygenGeneratorRating);
        System.out.println("CO2 Scrubber Rating: " + CO2ScrubberRating);
        System.out.println("Result: " + oxygenGeneratorRating * CO2ScrubberRating);
    }

    public String findOxygenGeneratorRating(ArrayList<ArrayList<Byte>> input, int rowLength) {
        ArrayList<ArrayList<Byte>> inputCopy = new ArrayList<>(input);
        for (int index = 0; index < rowLength; index++) {
            Byte mostCommonByte = findMostCommonByteAtIndex(inputCopy, index);
            if (Objects.equals(mostCommonByte, EQUAL_COUNT))
                mostCommonByte = 1;


            inputCopy = filterInputAtIndexWithByte(inputCopy, index, mostCommonByte);
            if (inputCopy.size() == 1)
                return byteArrayToString(inputCopy.get(0));
        }
        return byteArrayToString(inputCopy.get(0));
    }

    public String findCO2ScrubberRating(ArrayList<ArrayList<Byte>> input, int rowLength) {
        ArrayList<ArrayList<Byte>> inputCopy = new ArrayList<>(input);
        for (int index = 0; index < rowLength; index++) {
            Byte leastCommonByte = findLeastCommonByteAtIndex(inputCopy, index);
            if (Objects.equals(leastCommonByte, EQUAL_COUNT))
                leastCommonByte = 0;


            inputCopy = filterInputAtIndexWithByte(inputCopy, index, leastCommonByte);
            if (inputCopy.size() == 1)
                return byteArrayToString(inputCopy.get(0));
        }
        return byteArrayToString(inputCopy.get(0));
    }

    private ArrayList<ArrayList<Byte>> filterInputAtIndexWithByte(ArrayList<ArrayList<Byte>> inputCopy, int index, Byte filterByte) {
        inputCopy = (ArrayList<ArrayList<Byte>>) inputCopy
                .stream()
                .filter(row -> Objects.equals(row.get(index), filterByte))
                .collect(Collectors.toList());
        return inputCopy;
    }


    private String byteArrayToString(ArrayList<Byte> input) {
        StringBuilder output = new StringBuilder();
        for (Byte digit : input) {
            output.append(digit);
        }
        return output.toString();
    }

    public String findMostCommonAtIndex(ArrayList<ArrayList<Byte>> input, int index) {
        int sum = 0;
        for (ArrayList<Byte> row : input) {
            sum += row.get(index);
        }
        int countOf1s = sum;
        int countOf0s = input.size() - countOf1s;

        return countOf0s > countOf1s ? "0" : "1";
    }

    public Byte findMostCommonByteAtIndex(ArrayList<ArrayList<Byte>> input, int index) {
        int sum = 0;
        for (ArrayList<Byte> row : input) {
            sum += row.get(index);
        }
        int countOf1s = sum;
        int countOf0s = input.size() - countOf1s;
        if (countOf0s == countOf1s)
            return EQUAL_COUNT;
        return countOf0s > countOf1s ? Byte.valueOf((byte) 0) : Byte.valueOf((byte) 1);
    }


    public Byte findLeastCommonByteAtIndex(ArrayList<ArrayList<Byte>> input, int index) {
        int sum = 0;
        for (ArrayList<Byte> row : input) {
            sum += row.get(index);
        }
        int countOf1s = sum;
        int countOf0s = input.size() - countOf1s;
        if (countOf0s == countOf1s)
            return EQUAL_COUNT;
        return countOf0s < countOf1s ? Byte.valueOf((byte) 0) : Byte.valueOf((byte) 1);
    }

}
