import java.util.*;

public class Day9 {
    static Set<Point> tailVisitedPoints = new HashSet<>();
    static List<Point> rope = new ArrayList<>();

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        ArrayList<String> operations = new ArrayList<>();
        while (in.hasNextLine()) {
            operations.add(in.nextLine());
        }
        for (int i = 0; i < 10; i++)
            rope.add(new Point(0,0));

        tailVisitedPoints.add(new Point(0, 0));
        simulate(operations);
    }

    private static void simulate(List<String> operations) {
        for (String operation: operations) {
            String direction = operation.split(" ")[0];
            int steps = Integer.parseInt(operation.split(" ")[1]);
            System.out.println(operation);
            switch (direction) {
                case "U":
                    for (int i = 0; i < steps; i++) {
                        rope.get(0).row--;
                        moveAllKnots();
//                        printMatrix();
                    }
                    break;
                case "D":
                    for (int i = 0; i < steps; i++) {
                        rope.get(0).row++;
                        moveAllKnots();
//                        printMatrix();
                    }
                    break;
                case "L":
                    for (int i = 0; i < steps; i++) {
                        rope.get(0).col--;
                        moveAllKnots();
//                        printMatrix();
                    }
                    break;
                case "R":
                    for (int i = 0; i < steps; i++) {
                        rope.get(0).col++;
                        moveAllKnots();
//                        printMatrix();
                    }
                    break;
            }
        }
        System.out.println(tailVisitedPoints.size());
    }

    private static void printMatrix() {
        for (int row = -15; row < 20; row++) {
            for (int col = -15; col < 20; col++) {
                String ropeResult = ropeContains(row, col);
                if (!ropeResult.equals("")) {
                    System.out.printf("%2s", ropeResult);
                } else {
                    System.out.printf("%2s", ".");
                }
            }
            System.out.println();
        }
        System.out.println();
        System.out.println();
    }

    private static String ropeContains(int row, int col) {
        for (int i = 0; i < rope.size(); i++) {
            Point p = rope.get(i);
            if (p.row == row && p.col == col) {
                return String.valueOf(i);
            }
        }
        return "";
    }

    private static void moveAllKnots() {
        for (int knotIndex = 1; knotIndex <= 9; knotIndex++) {
            Point movedKnot = moveKnot(rope.get(knotIndex - 1).row,
                    rope.get(knotIndex - 1).col,
                    rope.get(knotIndex).row,
                    rope.get(knotIndex).col,
                    knotIndex == 9);
            rope.get(knotIndex).col = movedKnot.col;
            rope.get(knotIndex).row = movedKnot.row;
        }
    }

    private static Point moveKnot(int headRow, int headCol, int tailRow, int tailCol, boolean isTail) {

        /* head is 2 up/down/left/right

        up by 2 steps: headCol == tailCol, headRow = tailRow - 2
        down by 2 steps: headCol == tailCol, headRow = tailRow + 2
        left by 2 steps: headRow == tailRow, headCol = tailCol - 2
        right by 2 steps: headRow == tailRow, headCol = tailCol + 2
         */

        if (headCol == tailCol && headRow == tailRow - 2) {
            tailRow--;
        } else if (headCol == tailCol && headRow == tailRow + 2) {
            tailRow++;
        } else if (headRow == tailRow && headCol == tailCol - 2) {
            tailCol--;
        } else if (headRow == tailRow && headCol == tailCol + 2) {
            tailCol++;
        } else if (distance(headRow, headCol, tailRow, tailCol) >= 2) {
            if (headCol > tailCol && headRow > tailRow) {
                tailCol++; tailRow++;
            } else if (headCol > tailCol && headRow < tailRow) {
                tailCol++; tailRow--;
            } else if (headCol < tailCol && headRow < tailRow) {
                tailCol--; tailRow--;
            } else if (headCol < tailCol && headRow > tailRow) {
                tailCol--; tailRow++;
            }
        }
        if (isTail)
            tailVisitedPoints.add(new Point(tailRow, tailCol));
        return new Point(tailCol, tailRow);
    }

    private static double distance(int x1, int y1, int x2, int y2) {
        return Math.sqrt(((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)));
    }

    static class Point {
        int col;
        int row;

        public Point(int col, int row) {
            this.col = col;
            this.row = row;
        }

        @Override
        public boolean equals(Object obj) {
            if (this == obj)
                return true;
            if (obj == null)
                return false;
            if (getClass() != obj.getClass())
                return false;
            Point other = (Point) obj;
            return this.col == other.col && this.row == other.row;
        }

        @Override
        public int hashCode() {
            final int prime = 31;
            int result = 1;
            result = prime * result + col;
            result = prime * result + row;
            return result;
        }

    }
}
