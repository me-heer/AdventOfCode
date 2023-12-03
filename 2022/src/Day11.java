import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Scanner;

public class Day11 {
    static List<Monkey> monkeys = new ArrayList<>();

    static class Monkey {
        List<Long> items = new ArrayList<>();
        String operation;
        Long testDivisibleBy;
        Integer ifTrueThrowTo;
        Integer ifFalseThrowTo;
        Integer totalItemsInspected = 0;

        public Monkey() {}

        public Monkey(List<Long> items, String operation, Long testDivisibleBy, Integer ifTrueThrowTo, Integer ifFalseThrowTo) {
            this.items = items;
            this.operation = operation;
            this.testDivisibleBy = testDivisibleBy;
            this.ifTrueThrowTo = ifTrueThrowTo;
            this.ifFalseThrowTo = ifFalseThrowTo;
        }
    }

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        List<String> inputLines = new ArrayList<>();
        while (in.hasNextLine()) {
            inputLines.add(in.nextLine());
        }

        for (int input = 0; input < inputLines.size(); input++) {
            Monkey monkey = new Monkey();
            input++;

            String line = inputLines.get(input);
            String[] items = line.substring(line.indexOf(":") + 2).split(",");
            for (String item : items) {
                monkey.items.add(Long.parseLong(item.trim()));
            }

            input++;
            line = inputLines.get(input);
            monkey.operation = line.substring(line.indexOf(": ") + 2);

            input++;
            line = inputLines.get(input);
            monkey.testDivisibleBy = Long.parseLong(line.split("by")[1].trim());


            input++;
            line = inputLines.get(input);
            monkey.ifTrueThrowTo = Integer.parseInt(line.split("monkey")[1].trim());

            input++;
            line = inputLines.get(input);
            monkey.ifFalseThrowTo = Integer.parseInt(line.split("monkey")[1].trim());

            monkeys.add(monkey);
            input++;
        }
        for (int i = 0; i < 10000; i++)
            simulate();
        List<Integer> itemsSorted = new ArrayList<>();
        for (Monkey m: monkeys) {
            System.out.println(m.totalItemsInspected);
            itemsSorted.add(m.totalItemsInspected);
        }
        Collections.sort(itemsSorted);
//        System.out.println(itemsSorted.get(itemsSorted.size() - 1));
//        System.out.println(itemsSorted.get(itemsSorted.size() - 2));
//        System.out.println(itemsSorted.get(itemsSorted.size() - 1) * itemsSorted.get(itemsSorted.size() - 2));
    }

    private static void simulate() {
        for (Monkey monkey: monkeys) {
            List<Long> items = monkey.items;
            for (Long item: items) {
                monkey.totalItemsInspected++;
                String[] operation = monkey.operation.split(" = ")[1].split(" ");
                Long leftOperand = Long.parseLong(operation[0].replace("old", String.valueOf(item)));
                Long rightOperand = Long.parseLong(operation[2].replace("old", String.valueOf(item)));
                Long worryLevel = Long.valueOf(0);
                switch (operation[1]) {
                    case "+":
                        worryLevel = leftOperand + rightOperand;
                        break;
                    case "*":
                        worryLevel = leftOperand * rightOperand;
                        break;
                }
                Long worryLevelAfterBored = worryLevel / 3;
                /*
                Part 2
                 */
                worryLevelAfterBored = worryLevel;
                worryLevelAfterBored = worryLevelAfterBored % 9699690;
                Monkey monkeyToThrowTo;
                if (worryLevelAfterBored % monkey.testDivisibleBy == 0) {
                    monkeyToThrowTo = monkeys.get(monkey.ifTrueThrowTo);
                } else {
                    monkeyToThrowTo = monkeys.get(monkey.ifFalseThrowTo);
                }
                monkeyToThrowTo.items.add(worryLevelAfterBored);
            }
            monkey.items.clear();
        }
    }
}
