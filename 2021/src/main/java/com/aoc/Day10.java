package com.aoc;

import java.util.*;
import java.util.stream.Collectors;

public class Day10 extends AdventOfCode{
    @Override
    public void solve() {
        Scanner in = super.input;

        ArrayList<String> inputLines = new ArrayList<>();
        while (in.hasNextLine()) {
            inputLines.add(in.nextLine());
        }

        long points = 0;


        for(Iterator<String> inputLineIterator= inputLines.iterator(); inputLineIterator.hasNext();) {
            String inputLine = inputLineIterator.next();
            Stack<Character> stack = new Stack<>();
            for(Character c: inputLine.toCharArray()) {
                switch (c) {
                    case '{':
                    case '[':
                    case '(':
                    case '<': stack.push(c); break;
                    case '}':
                        Character matchCurly = stack.pop();
                        if(matchCurly != '{')
                            inputLineIterator.remove();
                        break;
                    case ')':
                        Character matchRound = stack.pop();
                        if(matchRound != '(')
                            inputLineIterator.remove();
                        break;
                    case '>':
                        Character matchAngled = stack.pop();
                        if(matchAngled != '<')
                            inputLineIterator.remove();
                        break;
                    case ']':
                        Character matchSquare = stack.pop();
                        if(matchSquare != '[')
                            inputLineIterator.remove();
                        break;
                }
            }


        }

        ArrayList<Long> totalPoints = new ArrayList<>();
        for(String inputLine: inputLines) {
            totalPoints.add(findExpectedBracketPoints(inputLine));
        }

        List<Long> sortedTotalPoints = totalPoints.stream().sorted().collect(Collectors.toList());
        System.out.println(sortedTotalPoints.get(sortedTotalPoints.size() / 2));
    }

    private Long findExpectedBracketPoints(String inputLine) {
        Stack<Character> stack = new Stack<>();
        for(Character c: inputLine.toCharArray())
        {
            switch (c) {
                case '{':
                case '[':
                case '(':
                case '<': stack.push(c); break;
                case '}':
                    Character matchCurly = stack.pop();
                    if(matchCurly != '{')
                        System.out.println("Error");
                    break;
                case ')':
                    Character matchRound = stack.pop();
                    if(matchRound != '(')
                        System.out.println("Error");
                    break;
                case '>':
                    Character matchAngled = stack.pop();
                    if(matchAngled != '<')
                        System.out.println("Error");
                    break;
                case ']':
                    Character matchSquare = stack.pop();
                    if(matchSquare != '[')
                        System.out.println("Error");
                    break;
            }

        }

        long points = 0;
        while (!stack.isEmpty()) {
            Character c = stack.pop();
            switch (c) {
                case '(': points = (points * 5) + 1; break;
                case '[': points = (points * 5) + 2; break;
                case '{': points = (points * 5) + 3; break;
                case '<': points = (points * 5) + 4; break;
            }
        }
        return points;
    }
}
