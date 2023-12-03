package com.aoc;

import java.lang.reflect.Array;
import java.util.ArrayList;
import java.util.Scanner;

public class Day4 extends AdventOfCode {
    @Override
    public void solve() {
        Scanner in = super.input;
        String[] numbers = in.nextLine().split(",");
        ArrayList<Board> boards = new ArrayList<>();
        while (in.hasNextLine()) {
            in.nextLine();
            Board board = new Board();
            board.readBoardFromInput(in);
            boards.add(board);
        }

        for (String number : numbers) {
                System.out.println(number + " is drawn");
                for (Board board : boards) {
                    if(!board.getIsWinner()) {
                        board.markNumber(Integer.parseInt(number));
                        board.validateWinner(Integer.parseInt(number));
                    }
                }
                if(Board.totalWinners == boards.size())
                    break;
        }
    }


    private class Board {

        ArrayList<ArrayList<BoardNumber>> board = new ArrayList<>();
        boolean isWinner = false;
        public int winnerOrder = -1;
        public static int totalWinners = 0;

        public void readRow(Scanner in) {
            ArrayList<BoardNumber> boardRow = new ArrayList<>();
            for (int i = 0; i < 5; i++) {
                BoardNumber number = new BoardNumber();
                number.readNumber(in);
                boardRow.add(number);
            }
            board.add(boardRow);
        }

        public void readBoardFromInput(Scanner in) {
            for (int i = 0; i < 5; i++) {
                readRow(in);
            }
        }

        public void markNumber(int number) {
            for (ArrayList<BoardNumber> row : board) {
                for (BoardNumber boardNumber : row) {
                    if (boardNumber.getNumber() == number) {
                        boardNumber.setMarked(true);
                        return;
                    }
                }
            }
        }

        private boolean getIsWinner() {
            return isWinner;
        }

        public boolean validateWinner(int number) {
            this.isWinner = isWinner();
            if (isWinner) {
                int sumOfAllUnmarkedNumbers = getSumOfUnmarkedBoardNumbers();
                int result = sumOfAllUnmarkedNumbers * number;
                totalWinners++;
                this.winnerOrder = totalWinners;
                printBoard();
                System.out.println("Board won.");
                System.out.println("Answer: " + result);
            }
            return isWinner;
        }

        private int getSumOfUnmarkedBoardNumbers() {
            int sum = 0;
            for (ArrayList<BoardNumber> row : board) {
                for (BoardNumber boardNumber : row) {
                    if (!boardNumber.isMarked())
                        sum += boardNumber.getNumber();
                }
            }
            return sum;
        }

        public boolean isWinner() {
            return isAnyRowComplete() || isAnyColumnComplete();
        }

        private boolean isAnyRowComplete() {
            for (ArrayList<BoardNumber> row : board) {
                int totalMarked = 0;
                for (BoardNumber boardNumber : row) {
                    if (boardNumber.isMarked())
                        totalMarked++;
                }
                if (totalMarked == 5)
                    return true;
            }
            return false;
        }

        private boolean isAnyColumnComplete() {
            for (int i = 0; i < 5; i++) {
                int totalMarked = 0;
                for (ArrayList<BoardNumber> row : board) {
                    if (row.get(i).isMarked())
                        totalMarked++;
                }
                if (totalMarked == 5)
                    return true;
            }
            return false;
        }

        private void printBoard() {
            System.out.println("-------------BOARD------------\n");
            for (ArrayList<BoardNumber> row : board) {
                for (BoardNumber number : row) {
                    if (number.isMarked())
                        System.out.format("%10d*", number.getNumber());
                    else
                        System.out.format("%10d", number.getNumber());
                }
                System.out.println();
            }
            System.out.println("-------------BOARD------------\n");
        }
    }

    private class BoardNumber {

        private boolean isMarked;
        private int number;

        public BoardNumber() {
            this.isMarked = false;
        }

        public boolean isMarked() {
            return isMarked;
        }

        public void setMarked(boolean marked) {
            isMarked = marked;
        }

        public int getNumber() {
            return number;
        }

        public void setNumber(int number) {
            this.number = number;
        }

        public void readNumber(Scanner in) {
            this.number = in.nextInt();
        }

    }
}
