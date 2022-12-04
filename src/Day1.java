import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;

public class Day1 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        ArrayList<String> lines = new ArrayList<>();
        while (in.hasNextLine()) {
            String line = in.nextLine();
            lines.add(line);
        }
        int[] elfCalories = new int[lines.size()];
        int j = 0;
        for (int i = 0; i < lines.size(); i++) {
            if (lines.get(i).equals("")) {
                j++; continue;
            }
            Integer prev = elfCalories[j];
            prev = prev + Integer.parseInt(lines.get(i));
            elfCalories[j] = prev;
        }
        System.out.println(elfCalories[elfCalories.length - 1]);
        Arrays.sort(elfCalories);
        System.out.println(elfCalories[elfCalories.length - 1]);
        System.out.println(elfCalories[elfCalories.length - 2]);
        System.out.println(elfCalories[elfCalories.length - 3]);
    }
}
