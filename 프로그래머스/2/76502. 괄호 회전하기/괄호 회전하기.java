import java.util.*;

class Solution {
    public int solution(String s) {
        int answer = 0;
        ArrayDeque<Character> deque = new ArrayDeque<>();
        char[] arr = s.toCharArray();
        for(char c: arr){
            deque.addLast(c);
        }

        for (int i=0; i<s.length(); i++){
            deque.addLast(deque.pollFirst());
            if (func(deque)){
                answer ++;
            }
        }
        return answer;
    }
    public boolean func(ArrayDeque<Character> deque){
        ArrayDeque<Character> stack = new ArrayDeque<>();
        for (Character c: deque){
            if(c=='('||c=='{'||c=='['){
                stack.push(c);
            }else{
                if(stack.isEmpty()) {
                    return false;
                }else if(c==')'){
                    if(stack.pop()!='('){
                        return false;
                    }
                }else if(c=='}'){
                    if(stack.pop()!='{'){
                        return false;
                    }
                }else if(c==']'){
                    if(stack.pop()!='['){
                        return false;
                    }
                }
        
            }
        }
        return stack.isEmpty();
    }
}