import java.util.ArrayList;
import java.util.Scanner;

public class Day2 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        ArrayList<String> lines = new ArrayList<>();
        while (in.hasNextLine()) {
            String line = in.nextLine();
            lines.add(line);
        }
        int total = 0;
        for (String line: lines) {
            // Count score
            String[] match = line.split(" ");
            String opponent = match[0];
            String you = match[1];
            total += countScoreAccordingToWin(opponent, you);
        }
        System.out.println(total);
    }

    private static int countScoreAccordingToWin(String opponent, String you) {
        if (opponent.equals("A") && you.equals("X")) {
            return 0 + 3;
        } else if (opponent.equals("A") && you.equals("Y")) {
            return 3 + 1;
        } else if (opponent.equals("A") && you.equals("Z")) {
            return 6 + 2;
        }

        if (opponent.equals("B") && you.equals("X")) {
            return 1 + 0;
        } else if (opponent.equals("B") && you.equals("Y")) {
            return 2 + 3;
        } else if (opponent.equals("B") && you.equals("Z")) {
            return 6 + 3;
        }

        if (opponent.equals("C") && you.equals("X")) {
            return 0 + 2;
        } else if (opponent.equals("C") && you.equals("Y")) {
            return 3 + 3;
        } else if (opponent.equals("C") && you.equals("Z")) {
            return 6 + 1;
        }
        return 0;
    }

    private static int countScore(String opponent, String you) {
        int winScore = 0;
        if (opponent.equals("A") && you.equals("X")) {
            winScore = 3;
        } else if (opponent.equals("A") && you.equals("Y")) {
            winScore = 6;
        } else if (opponent.equals("A") && you.equals("Z")) {
            winScore = 0;
        }

        if (opponent.equals("B") && you.equals("X")) {
            winScore = 0;
        } else if (opponent.equals("B") && you.equals("Y")) {
            winScore = 3;
        } else if (opponent.equals("B") && you.equals("Z")) {
            winScore = 6;
        }

        if (opponent.equals("C") && you.equals("X")) {
            winScore = 6;
        } else if (opponent.equals("C") && you.equals("Y")) {
            winScore = 0;
        } else if (opponent.equals("C") && you.equals("Z")) {
            winScore = 3;
        }

        if (you.equals("X")) {
            winScore += 1;
        } else if (you.equals("Y")) {
            winScore += 2;
        } else if (you.equals("Z")) {
            winScore += 3;
        }
        return winScore;
    }
}
