import java.util.*;

public class Day8 {
    static List<List<Integer>> trees = new ArrayList<>();
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        while (in.hasNextLine()) {
            List<Integer> row = new ArrayList<>();
            String line = in.nextLine();
            for (int i = 0; i < line.length(); i++) {
                row.add(Integer.parseInt(String.valueOf(line.charAt(i))));
            }
            trees.add(row);
        }

        int totalVisibleTrees = 0;
        for (int i = 0; i < trees.size(); i++) {
            for (int j = 0; j < trees.get(0).size(); j++) {
                if (isVisible(i, j))
                    totalVisibleTrees++;
            }
        }

        System.out.println(totalVisibleTrees);

        List<Integer> scenicScores = new ArrayList<>();
        for (int i = 0; i < trees.size(); i++) {
            for (int j = 0; j < trees.get(0).size(); j++) {
                scenicScores.add(countScenicScore(i, j));
            }
        }

        System.out.println(Collections.max(scenicScores));
    }


    private static int countScenicScore(int row, int col) {
        int currentTree = trees.get(row).get(col);
        // go through 4 directions

        // go up
        int up = 0;
        for (int i = row - 1; i >= 0; i--) {
            up++;
            if (trees.get(i).get(col) >= currentTree) {
                break;
            }
        }

        int down = 0;
        for (int i = row + 1; i < trees.size(); i++) {
            down++;
            if (trees.get(i).get(col) >= currentTree) {
                break;
            }
        }

        int left = 0;
        for (int i = col - 1; i >= 0; i--) {
            left++;
            if (trees.get(row).get(i) >= currentTree) {
                break;
            }
        }

        int right = 0;
        for (int i = col + 1; i < trees.get(0).size(); i++) {
            right++;
            if (trees.get(row).get(i) >= currentTree) {
                break;
            }
        }

        return up * down * left * right;
    }

    private static boolean isVisible(int row, int col) {
        if (row == 0 || col == 0 || row == trees.size() - 1 || col == trees.get(0).size() - 1)
            return true;

        int currentTree = trees.get(row).get(col);
        // go through 4 directions

        // go up
        boolean isVisibleFromUp = true;
        for (int i = row - 1; i >= 0; i--) {
            if (trees.get(i).get(col) >= currentTree) {
                isVisibleFromUp = false;
                break;
            }
        }

        boolean isVisibleFromDown = true;
        for (int i = row + 1; i < trees.size(); i++) {
            if (trees.get(i).get(col) >= currentTree) {
                isVisibleFromDown = false;
                break;
            }
        }

        boolean isVisibleFromLeft = true;
        for (int i = col - 1; i >= 0; i--) {
            if (trees.get(row).get(i) >= currentTree) {
                isVisibleFromLeft = false;
                break;
            }
        }

        boolean isVisibleFromRight = true;
        for (int i = col + 1; i < trees.get(0).size(); i++) {
            if (trees.get(row).get(i) >= currentTree) {
                isVisibleFromRight = false;
                break;
            }
        }

        return isVisibleFromUp || isVisibleFromDown || isVisibleFromLeft || isVisibleFromRight;
    }
}
