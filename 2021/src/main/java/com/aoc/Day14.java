package com.aoc;

import java.util.*;

public class Day14 extends AdventOfCode {
    Map<String, String> insertionRules = new HashMap<>();
    String polymerTemplateStr = "";
    LinkedList<Character> polymerLinkedList = new LinkedList<>();

    @Override
    public void solve() {
        Scanner in = input;
        polymerTemplateStr = in.nextLine();
        for (Character c : polymerTemplateStr.toCharArray())
            polymerLinkedList.add(c);

        while (in.hasNextLine()) {
            String line = in.nextLine();
            String[] rules = line.split(" -> ");
            if (rules.length == 2) {
                insertionRules.put(rules[0], rules[1]);
            }
        }

        printLinkedList();
        for (int i = 0; i < 40; i++) {
            System.out.println("STEP: " + i);
            executeStep();
        }
        System.out.println(polymerLinkedList.size());
//        findMostAndLeastCommonElement(polymerTemplateStr);
    }

    private void printLinkedList() {
        for (Character c : polymerLinkedList)
            System.out.print(c);
        System.out.println();
    }

    private void findMostAndLeastCommonElement(String polymerTemplate) {
        Map<Character, Long> elementWithOccurrences = new HashMap<>();
        for (int i = 0; i < polymerTemplate.length(); i++) {
            Character character = polymerTemplate.charAt(i);
            Long previousOccurrences = elementWithOccurrences.getOrDefault(character, 0l);
            Long updatedOccurrences = previousOccurrences + 1;
            elementWithOccurrences.put(character, updatedOccurrences);
        }
        elementWithOccurrences.entrySet().stream()
                .sorted(Map.Entry.comparingByValue())
                .forEach(System.out::println);
    }

    private void executeStep() {
        System.out.println("MAPPING RULES TO INDEXES");
        Map<String, ArrayList<Integer>> rulesToIndex = getRulesToIndexMapping(insertionRules, polymerLinkedList);
        System.out.println("EXECUTING RULES");
        for (Map.Entry<String, ArrayList<Integer>> applicableRule : rulesToIndex.entrySet()) {
            String ruleCondition = applicableRule.getKey();
            String insertionCharacter = insertionRules.get(applicableRule.getKey());

            int[] currentIndexes = rulesToIndex.get(ruleCondition).stream().mapToInt(i -> i).toArray();

            System.out.println("APPLYING RULES");
            for (int i = 0; i < currentIndexes.length; i++) {
                Integer insertIndex = currentIndexes[i];
                System.out.println("ADDING CHAR TO POLYMER");
                polymerLinkedList.add(insertIndex + 1, insertionCharacter.charAt(0));
                System.out.println("ADDED CHAR TO POLYMER");
//                String updatedPolymerState =
//                        new String(currentPolymerState.substring(0, insertIndex + 1) + insertionCharacter + currentPolymerState.substring(insertIndex + 1));
                rulesToIndex = updateIndexes(ruleCondition, insertIndex, rulesToIndex);
                //update remaining
                for (int j = i; j < currentIndexes.length; j++) {
                    currentIndexes[j]++;
                }
            }
            System.out.println("APPLIED RULES");
//            for (Integer insertIndex: rulesToIndex.get(ruleCondition)) {
//                String updatedPolymerState =
//                        currentPolymerState.substring(0, insertIndex + 1) + insertionCharacter + currentPolymerState.substring(insertIndex + 1);
//                currentPolymerState = updatedPolymerState;
//                rulesToIndex = updateIndexes(ruleCondition, insertIndex, rulesToIndex);
//            }
        }
    }

    private Map<String, ArrayList<Integer>> updateIndexes(String currentRuleCondition, Integer currentRuleIndex,
                                                          Map<String, ArrayList<Integer>> rulesToIndexMap) {
        System.out.println("UPDATING INDEXES");
        Map<String, ArrayList<Integer>> updatedRulesToIndexMap = new HashMap<>();
        for (Map.Entry<String, ArrayList<Integer>> ruleToIndexEntry : rulesToIndexMap.entrySet()) {
            String ruleEntry = ruleToIndexEntry.getKey();
            if (!ruleEntry.equals(currentRuleCondition)) {
                ArrayList<Integer> currentIndexes = ruleToIndexEntry.getValue();
                ArrayList<Integer> updatedCurrentIndexes = new ArrayList<>();
                for (Integer currentIndex : currentIndexes) {
                    if (currentIndex > currentRuleIndex)
                        currentIndex++;
                    updatedCurrentIndexes.add(currentIndex);
                }
                updatedRulesToIndexMap.put(ruleEntry, updatedCurrentIndexes);
            }
        }
        System.out.println("INDEXES UPDATED");
        return updatedRulesToIndexMap;
    }

    private Map<String, ArrayList<Integer>> getRulesToIndexMapping(Map<String, String> applicableRules,
                                                                   LinkedList<Character> currentPolymerState) {
        Map<String, ArrayList<Integer>> rulesToIndexMapping = new HashMap<>();
        for (Map.Entry<String, String> applicableRule : applicableRules.entrySet()) {
            String ruleCondition = applicableRule.getKey();
            for (int i = 0; i < currentPolymerState.size(); i++) {
                int upperBound = currentPolymerState.size() > (i + 2) ? (i + 2) : currentPolymerState.size();
                String subStr = getSubstring(i, upperBound, currentPolymerState);
                if (subStr.contains(ruleCondition)) {
                    ArrayList<Integer> occurrenceIndexes = rulesToIndexMapping.getOrDefault(ruleCondition,
                            new ArrayList<Integer>());
                    occurrenceIndexes.add(i);
                    rulesToIndexMapping.put(ruleCondition, occurrenceIndexes);
                }
            }
        }
        return rulesToIndexMapping;
    }

    private String getSubstring(int startIndex, int endIndex, LinkedList<Character> linkedList) {
        String substring = "";
        for (int i = startIndex; i < endIndex; i++)
            substring += linkedList.get(i);
        return substring;
    }

    private Map<String, String> findApplicableRules(Map<String, String> allRules, String currentPolymerState) {
        Map<String, String> applicableRules = new HashMap<>();
        for (Map.Entry<String, String> rule : allRules.entrySet()) {
            String ruleCondition = rule.getKey();
            if (currentPolymerState.indexOf(ruleCondition) != -1)
                applicableRules.put(rule.getKey(), rule.getValue());
        }
        return applicableRules;
    }

}
