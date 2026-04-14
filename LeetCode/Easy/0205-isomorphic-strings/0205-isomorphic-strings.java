class Solution {
    public boolean isIsomorphic(String s, String t) {
        Map<Character, Character> trans = new HashMap<>();
        Set<Character> used = new HashSet<>();
        for (int i = 0; i < s.length(); i++) {
            Character sKey = s.charAt(i);
            Character tKey = t.charAt(i);

            if (!trans.containsKey(sKey)) {
                if (used.contains(tKey)) {
                    return false;
                }
                trans.put(sKey, tKey);
                used.add(tKey);
                continue;
            }
            if (trans.get(sKey) != tKey) {
                return false;
            }
        }
        return true;
    }
}