class Solution {
    public String longestCommonPrefix(String[] strs) {
        int n = strs.length;
        StringBuilder sb = new StringBuilder();
        int minLen = strs[0].length();
        for (String str: strs) {
            minLen = Math.min(minLen, str.length());
        }
        for (int i = 0; i < minLen; i++) {
            Character c = null;
            boolean flag = true;
            for (String str: strs) {
                if (c == null) {
                    c = str.charAt(i);
                } else{
                    if (c != str.charAt(i)){
                        flag = false;
                        break;
                    }
                }
            }
            if (flag){
                sb.append(c);
            }
            else {
                break;
            }
        }
        return sb.toString();
    }
}