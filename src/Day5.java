import com.sun.org.apache.xerces.internal.impl.io.UCSReader;

import java.util.*;

public class Day5 {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        ArrayList<String> lines = new ArrayList<>();
        while (in.hasNextLine()) {
            String line = in.nextLine();
            lines.add(line);
        }

        // Find number of stacks
        List<Stack<String>> stacks = new ArrayList<>();
        int numberOfStacks = ((lines.get(0).length()) + 1) / 4;
        for (int i = 0; i < numberOfStacks; i++)
            stacks.add(new Stack<>());

        List<List<String>> inputList = new ArrayList<>();
        for (int i = 0; i < numberOfStacks; i++)
            inputList.add(new ArrayList<>());


        int readIndex;
        // Read stacks
        for (readIndex = 0; readIndex < lines.size(); readIndex++) {
            if (Character.isDigit(lines.get(readIndex).charAt(1)))
                break;
            for (int j = 0; j < numberOfStacks; j++) {
                String crate = String.valueOf(lines.get(readIndex).charAt(1 + (j * 4)));
                if (!crate.equals(" "))
                    inputList.get(j).add(crate);
            }
        }

        for (int i = 0; i < inputList.size(); i++) {
            Collections.reverse(inputList.get(i));
            Stack<String> s = new Stack<>();
            s.addAll(inputList.get(i));
            stacks.set(i, s);
        }

        readIndex += 2;

        for (; readIndex < lines.size(); readIndex++) {
            // parse operation
            String[] input = lines.get(readIndex).split(" ");
            int numberOfCrates = Integer.parseInt(input[1]);
            int from = Integer.parseInt(input[3]);
            int to = Integer.parseInt(input[5]);


            // Part 1
//            for (int j = 0; j < numberOfCrates; j++) {
//                String crate = stacks.get(from - 1).pop();
//                stacks.get(to - 1).push(crate);
//            }

            // Part 2
            ArrayList<String> cratesToBeMoved = new ArrayList<>();
            for (int j = 0; j < numberOfCrates; j++) {
                String crate = stacks.get(from - 1).pop();
                cratesToBeMoved.add(crate);
            }

            Collections.reverse(cratesToBeMoved);
            for (String crate: cratesToBeMoved)
                stacks.get(to - 1).push(crate);
        }

        String result = "";
        for (int i = 0; i < stacks.size(); i++)
            result = result.concat(stacks.get(i).pop());
        System.out.println(result);

    }
}
