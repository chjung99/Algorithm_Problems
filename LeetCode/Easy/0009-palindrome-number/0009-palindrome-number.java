class Solution {
    public boolean isPalindrome(int x) {
        String src = String.valueOf(x);
        String rev = new StringBuilder(src).reverse().toString();
        return src.equals(rev);
    }
}