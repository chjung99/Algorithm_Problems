class Solution {
    public int hIndex(int[] citations) {
        int n = citations.length;

        Arrays.sort(citations);
        
        TreeMap<Integer, Integer> indexMap = new TreeMap<>();

        int maxValue = 0;

        for (int i = 0 ; i < n; i++) {
            maxValue = Math.max(maxValue, citations[i]);

            if (!indexMap.containsKey(citations[i])) {
                indexMap.put(citations[i], i);
            }
        }

        int left = 0;
        int right = 1001;
        int hMid = 0;
        int hCnt = 0;

        while (left + 1 < right) {
            
            hMid = (left + right) / 2;
            // System.out.println(hMid);
            // System.out.println(indexMap.ceilingEntry(hMid).getValue());
            if (indexMap.ceilingEntry(hMid) == null) {
                right = hMid;
                continue;
            }
            hCnt = n - indexMap.ceilingEntry(hMid).getValue();
            // System.out.println(left+" " +right + " " + hMid + " " + hCnt);
            if (hMid > hCnt) {
                right = hMid;
            } else {
                left = hMid;
            }
        }

        return left;
        // for (Integer key : indexMap.keySet()) {
        //     System.out.println(key + " : " + indexMap.get(key));
        // }


        

        // 0 1 3 5 6
    }
}