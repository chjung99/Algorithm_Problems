class Solution {
    public int[] plusOne(int[] digits) {
        int carry = 1;
        List<Integer> ls = new ArrayList<>();

        for (int i = digits.length - 1; i >= 0; i--) {
            int cur = digits[i] + carry;
            if (cur < 10) {
                carry = 0;
            }
            ls.addFirst(cur % 10);
        }

        if (carry > 0) {
            ls.addFirst(carry);
        }
        return ls.stream()
                           .mapToInt(Integer::intValue)
                           .toArray();
    }
}