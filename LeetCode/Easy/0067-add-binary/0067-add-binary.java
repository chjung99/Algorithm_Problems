class Solution {
    public String addBinary(String a, String b) {
        StringBuilder sb = new StringBuilder();
        a = new StringBuilder(a).reverse().toString();
        b = new StringBuilder(b).reverse().toString();
        int maxLen = Math.max(a.length(), b.length());
        boolean carry = false;

        for (int i = 0; i < maxLen; i++) {
            int setCount = 0;
            if (i < a.length() && i < b.length()) {
                setCount = getSetCount(a.charAt(i), b.charAt(i), carry);
                // System.out.println("#" + 1);

            } else if (i >= a.length() && i < b.length()) {
                setCount = getSetCount(b.charAt(i), carry);
                // System.out.println("#"+ 2);

            } else if (i < a.length() && i >= b.length()){
                setCount = getSetCount(a.charAt(i), carry);
                // System.out.println("#" + 3);

            } else {
                // System.out.println("#" + 4);
                continue;
            }
            
            sb.append(Integer.toString(setCount % 2));
            // System.out.println(setCount);
            if (setCount >= 2) {
                carry = true;
            } else {
                carry = false;
            }

            
        }

        if (carry) {
            sb.append('1');
        }
        return sb.reverse().toString();
    }
    public int getSetCount(Character x, Character y, boolean carry) {
        int cnt = 0;

        if (carry) {
            cnt ++;
        }

        if (x == '1') {
            cnt ++;
        }

        if (y == '1') {
            cnt ++;
        }
        return cnt;
    }

    public int getSetCount(Character x, boolean carry) {
        int cnt = 0;

        if (carry) {
            cnt ++;
        }

        if (x == '1') {
            cnt ++;
        }

        return cnt;
    }
}