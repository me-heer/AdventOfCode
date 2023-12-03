import java.util.*;

public class Day10 {
    static int cycles = 1;
    static int X = 1;
    static Map<Integer, Integer> result = new HashMap<>();
    static Map<Integer, Integer> spriteByCycle = new HashMap<>();

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        List<String> operations = new ArrayList<>();
        while (in.hasNextLine()) {
            operations.add(in.nextLine());
        }
        for (String operation: operations) {
            if (operation.startsWith("noop")) {
                result.put(cycles, cycles * X); printResult(operation);
                spriteByCycle.put(cycles, X);
                cycles++;
                result.put(cycles, cycles * X); printResult(operation);
                spriteByCycle.put(cycles, X);
            } else {
                result.put(cycles, cycles * X); printResult(operation);
                spriteByCycle.put(cycles, X);
                cycles++;
                result.put(cycles, cycles * X); printResult(operation);
                spriteByCycle.put(cycles, X);
                cycles++;
                X += Integer.parseInt(operation.split(" ")[1]);
                result.put(cycles, cycles * X); printResult(operation);
                spriteByCycle.put(cycles, X);
            }
        }
        printCrt();
    }

    private static void printCrt() {
        int fortyCounter = 1;
        for (int i = 0; i < cycles; i++) {
            Integer currentX = spriteByCycle.get(i + 1);
            currentX = currentX % 40;
            if (i % 40 >= currentX - 1 && i % 40 <= currentX + 1)
                System.out.printf("%1s", "#");
            else
                System.out.printf("%1s", ".");

            fortyCounter++;
            if (fortyCounter == 41) {
                System.out.println();
                fortyCounter = 1;
            }
        }
    }

    private static void printResult(String currentOperation) {
//        System.out.println(currentOperation + " Cycle: " + cycles + " X: " + X + ", Multiplication: " + cycles * X);
    }
}
