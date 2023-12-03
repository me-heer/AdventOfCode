import java.util.*;

public class Day3 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        ArrayList<String> rucksacks = new ArrayList<>();
        while (in.hasNextLine()) {
            String line = in.nextLine();
            rucksacks.add(line);
        }
        Integer sumOfPriorities = 0;
        Map<String, Integer> priorities = preparePriorities();
        ArrayList<String> commonItems = new ArrayList<>();
        for (String rucksack: rucksacks) {
            String firstHalf = rucksack.substring(0, rucksack.length() / 2);
            String secondHalf = rucksack.substring(rucksack.length() / 2);

            Set<Character> firstHalfSet = new HashSet<>();
            for (char c:  firstHalf.toCharArray())
                firstHalfSet.add(c);
            Set<Character> secondHalfSet = new HashSet<>();
            for (char c:  secondHalf.toCharArray())
                secondHalfSet.add(c);

            firstHalfSet.retainAll(secondHalfSet);
            String commonItem = String.valueOf(firstHalfSet.stream().findFirst().get());
            sumOfPriorities += priorities.get(commonItem);
        }
        System.out.println(sumOfPriorities);


        // Part 2
        Integer sumOfBadgePriorities = 0;
        for (int i = 0; i < rucksacks.size(); i = i + 3) {
            String first = rucksacks.get(i);
            String second = rucksacks.get(i + 1);
            String third = rucksacks.get(i + 2);

            Set<Character> firstSet = new HashSet<>();
            for (char c:  first.toCharArray())
                firstSet.add(c);

            Set<Character> secondSet = new HashSet<>();
            for (char c:  second.toCharArray())
                secondSet.add(c);

            Set<Character> thirdSet = new HashSet<>();
            for (char c:  third.toCharArray())
                thirdSet.add(c);
            firstSet.retainAll(secondSet);
            firstSet.retainAll(thirdSet);
            String commonItem = String.valueOf(firstSet.stream().findFirst().get());
            sumOfBadgePriorities += priorities.get(commonItem);
        }
        System.out.println(sumOfBadgePriorities);
    }

    private static Map<String, Integer> preparePriorities() {
        Map<String, Integer> priorities = new HashMap<>();
        char c = 'a';
        for (int i = 0; i < 26; i++) {
            priorities.put(String.valueOf(c), i + 1);
            c++;
        }
        c = 'A';
        for (int i = 0; i < 26; i++) {
            priorities.put(String.valueOf(c), i + 1 + 26);
            c++;
        }
        return priorities;
    }
}
