import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;

public class Day4 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        ArrayList<String> lines = new ArrayList<>();
        while (in.hasNextLine()) {
            String line = in.nextLine();
            lines.add(line);
        }
        Integer totalFullyContained = 0;
        Integer totalOverlapAtAll = 0;
        for (String line: lines) {
            String firstElf = line.split(",")[0];
            String secondElf = line.split(",")[1];

            Integer firstStart = Integer.parseInt(firstElf.split("-")[0]);
            Integer firstEnd = Integer.parseInt(firstElf.split("-")[1]);

            Integer secondStart = Integer.parseInt(secondElf.split("-")[0]);
            Integer secondEnd = Integer.parseInt(secondElf.split("-")[1]);

            boolean isFullyContained = isFullyContained(firstStart, firstEnd, secondStart, secondEnd);
            boolean overlapAtAll = overlapAtAll(firstStart, firstEnd, secondStart, secondEnd);
            System.out.println(line + " overlapAtAll: " + overlapAtAll);
            if (isFullyContained)
                totalFullyContained++;
            if (overlapAtAll)
                totalOverlapAtAll++;
        }
        System.out.println(totalOverlapAtAll);
    }

    private static boolean isFullyContained(Integer a, Integer b, Integer y, Integer z) {
        return (y <= a && z >= b) || (a <= y && b >= z);
    }

    private static boolean overlapAtAll(Integer a, Integer b, Integer y, Integer z) {
        return (a <= y && y <= b) || (y <= a && a <= z);
    }
}
