package com.aoc;

import java.util.Scanner;

public class Day2 extends AdventOfCode {

    private int horizontalPosition = 0;
    private int depth = 0;
    private int aim = 0;

    @Override
    public void solve() {
        Scanner in = super.input;
        while (in.hasNextLine()) {
            String input = in.nextLine();
            Instruction instruction = Instruction.valueOf(input.split(" ")[0].toUpperCase());
            int units = Integer.parseInt(input.split(" ")[1]);

            processInstruction(instruction, units);
        }

        System.out.println("Horizontal Position: " + horizontalPosition);
        System.out.println("Depth: " + depth);
        System.out.println("Aim: " + aim);
        System.out.println("Multiplication: " + (horizontalPosition * depth));
    }

    private void processInstruction(Instruction instruction, int units) {
        switch (instruction) {
            case FORWARD -> updateHorizontalPosition(units);
            case UP, DOWN -> updateAim(units, instruction);
        }
    }

    private void updateHorizontalPosition(int units) {
        horizontalPosition = horizontalPosition + units;
        depth = depth + (aim * units);
    }

    private void updateAim(int units, Instruction instruction) {
        switch (instruction) {
            case DOWN -> aim = aim + units;
            case UP -> aim = aim - units;
        }
    }

    enum Instruction {
        FORWARD("forward"),
        UP("up"),
        DOWN("down");

        private final String instruction;

        Instruction(final String instruction) {
            this.instruction = instruction;
        }
    }

}
