class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int[] result = new int[m+n];
        
        int L = 0;
        int R = 0;

        while ( L + R < m+n) {
            if (L >= m || R >= n) {
                break;
            }
            if (nums1[L] <= nums2[R]) {
                result[L+R] = nums1[L];
                L++;
                continue;
            }

            if (nums1[L] > nums2[R]) {
                result[L+R] = nums2[R];
                R++;
            }
        }

        if (L + R < m+n && L >= m) {
            for (int i = L+R; i < m+n; i++) {
                result[i] = nums2[R++];
            }
        }

        if (L + R < m+n &&  R >= n) {
            for (int i = L+R; i < m+n; i++) {
                result[i] = nums1[L++];
            }
        }

        for (int i = 0; i < m+n; i++) {
            nums1[i] = result[i];
        }

    }
}