class Solution {
    public boolean isPalindrome(String s) {
        String filterd = s.toLowerCase();
        StringBuilder filterdBuilder = new StringBuilder();
        for (int i = 0; i < filterd.length(); i++){
            Character c = filterd.charAt(i);
            if ((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')||(c >= '0' && c <= '9')){
                filterdBuilder.append(c.toString());
            }
        }
        String ori = filterdBuilder.toString();
        String rev = new StringBuilder(ori).reverse().toString();

        System.out.println(ori);
        System.out.println(rev);
        
        return rev.equals(ori);
    }
}