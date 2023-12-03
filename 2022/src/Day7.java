import javax.sound.midi.SysexMessage;
import java.util.*;
import java.util.stream.Collectors;

public class Day7 {
    static Set<Node> directories = new HashSet<>();


    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        List<String> input = new ArrayList<>();
        while (in.hasNextLine()) {
            input.add(in.nextLine());
        }

        // Create root directory
        Node root = new Node(null, "/", null, true, null);
        directories.add(root);
        Node current = root;
        for (int i = 1; i < input.size(); i++) {
            String inputLine = input.get(i);
            if (inputLine.startsWith("$")) {
                // is an operation
                String operation = inputLine.split(" ")[1];
                switch (operation) {
                    case "cd":
                        // if '..', move current back
                        // else move forward
                        if (inputLine.split(" ")[2].equals("..")) {
                            // go back
                            current = current.parent;
                        } else {
                            String childDirectoryName = inputLine.split(" ")[2];
                            // find dir with name == childDirectoryName in current.children
                            for (Node child: current.children) {
                                if (child.isDirectory && child.name.equals(childDirectoryName)) {
                                    current = child;
                                    break;
                                }
                            }
                        }



                        break;
                    case "ls":
                        // Parse the next statements till you encounter $
                        int dirContentIndex;
                        for (dirContentIndex = i + 1;
                             dirContentIndex < input.size() && !input.get(dirContentIndex).startsWith("$");
                             dirContentIndex++) {
                            String nodeStr = input.get(dirContentIndex);
                            Node node;
                            if (nodeStr.startsWith("dir")) {
                                node = new Node(current,
                                        nodeStr.split(" ")[1],
                                        null,
                                        true,
                                        null);
                                directories.add(node);
                            } else {
                                node = new Node(current,
                                        nodeStr.split(" ")[1],
                                        Integer.parseInt(nodeStr.split(" ")[0]),
                                        false,
                                        null);
                            }
                            current.children.add(node);
                        }
                        i = dirContentIndex - 1;
                        break;
                }
            }
        }
        countSize(root);
        Integer result = 0;
        for (Node dir: directories) {
            if (dir.size < 100000) {
                result += dir.size;
            }
        }
        System.out.println("Answer Part-1: " + result);
        Integer remainingSpace = 70000000 - root.size;
        System.out.println("Remaining space: " + remainingSpace);
        Integer minRequired = 30000000 - remainingSpace;
        System.out.println("Min required space for dir: " + minRequired);

        List<Node> sorted = directories.stream().sorted((o1, o2) -> o1.size.compareTo(o2.size)).collect(Collectors.toList());
        for (Node dir: sorted) {
            if (dir.size >= minRequired) {
                System.out.println(dir.size);
                break;
            }
        }

    }

    private static Integer countSize(Node dir) {
        if (!dir.isDirectory)
            return dir.size;
        Integer total = 0;
        for (Node child: dir.children) {
            total += countSize(child);
        }
        dir.size = total;
        return total;
    }

    static class Node {
        Node parent;
        String name;
        Integer size;
        boolean isDirectory;
        List<Node> children = new ArrayList<>();



        public Node(Node parent, String name, Integer size, boolean isDirectory, List<Node> children) {
            this.parent = parent;
            this.name = name;
            this.size = size;
            this.isDirectory = isDirectory;
            if (children != null)
                this.children.addAll(children);
        }
    }
}
