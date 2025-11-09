import java.util.*;

class Solution {

    public int[] twoSum(int[] nums, int target) {
        int[] answer = new int[2];
        Map<Integer, ArrayDeque<Integer>> hashMap = new HashMap<>();

        for (int i=0; i < nums.length; i++){
            int key = nums[i];
            if (hashMap.containsKey(key)){
                ArrayDeque<Integer> bucket =  hashMap.get(key);
                bucket.push(i);
                hashMap.put(nums[i], bucket);
            } else {
                hashMap.put(nums[i], new ArrayDeque<>(List.of(i)));
            }
        }

        for (int i=0; i < nums.length ; i++){
            int key = target - nums[i];
            
            if (hashMap.containsKey(key)){
                ArrayDeque<Integer> bucket = hashMap.get(key);
                answer[0] = i;
                if (key == nums[i]){
                    while (!bucket.isEmpty()){
                        answer[1] = bucket.pop();
                        if (answer[1] != i){
                            break;
                        }
                    }
                    if (answer[0] != answer[1]) break;
                     
                } else {
                    answer[1] = bucket.pop();
                    break;

                }
            }
        }
        
        return answer;
    }
}