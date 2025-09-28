import java.util.*;

class Solution {
    
    public int[] solution(int N, int[] stages) {
        
        int M = stages.length;
        Arrays.sort(stages);
        Map<Integer, Double> map = new HashMap<Integer, Double>();
        for (int i=1; i<=N; i++){
            int upper = getUpperBound(stages, i);
            int lower = getLowerBound(stages, i);
            System.out.println(lower+ " "+ upper);
            if (lower==upper){
                map.put(i, Double.valueOf(0));
            }else{
                map.put(i, Double.valueOf(upper-lower)/(M-lower));
            }
        }
        System.out.println(map);
        
        List<Integer> keySet = new ArrayList<>(map.keySet());

        // Value 값으로 오름차순 정렬
        keySet.sort(new Comparator<Integer>() {
            @Override
            public int compare(Integer o1, Integer o2) {
                if (map.get(o1).equals(map.get(o2))){
                    return o1.compareTo(o2);
                }
                else{
                    return map.get(o2).compareTo(map.get(o1));
                }
            }
        });
        // System.out.print(keySet);
        int[] answer = keySet.stream()
            .mapToInt(x->x.intValue())
            .toArray();
        return answer;
    }
    public int getLowerBound(int[] stages, int x){
        int left = 0;
        int right = stages.length;
        
        while (left<right){
            int mid = (left + right) / 2;
            if (stages[mid] >= x){
                right = mid;
            } else{
                left = mid + 1; 
            }
        }
        return left;
    }
    public int getUpperBound(int[] stages, int x){
        int left = 0;
        int right = stages.length;
        
        while (left<right){
            int mid = (left + right) / 2;
            if (stages[mid] > x){
                right = mid;
            } else{
                left = mid+1; 
            }
        }
        return left;
    }
}