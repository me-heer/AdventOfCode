import java.util.*;

public class Day6 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String line = in.nextLine();
        int markerIndex;
        for (markerIndex = 13; markerIndex < line.length(); markerIndex++) {
            // check 0,1,2,3 and see if they are different
            Set<Character> characters = new HashSet<>();
            for (int j = 0; j < 14; j++)
                characters.add(line.charAt(markerIndex - j));
            if (characters.size() == 14)
                break;
        }
        System.out.println(markerIndex + 1);
    }
}
