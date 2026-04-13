class Solution {
    public boolean canConstruct(String ransomNote, String magazine) {
        HashMap<Character, Integer> rMap = new HashMap<>();
        HashMap<Character, Integer> mMap = new HashMap<>();

        for (int i = 0; i < ransomNote.length(); i++) {
            Character key = ransomNote.charAt(i);
            rMap.put(key, rMap.getOrDefault(key, 0) + 1);
        }

        for (int i = 0; i < magazine.length(); i++) {
            Character key = magazine.charAt(i);
            mMap.put(key, mMap.getOrDefault(key, 0) + 1);
        }

        for (Character key: rMap.keySet()) {
            if (rMap.get(key) > mMap.getOrDefault(key, 0)) {
                return false;
            }
        }
        return true;
    }
}