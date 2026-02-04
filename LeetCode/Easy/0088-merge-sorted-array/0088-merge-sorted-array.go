func merge(nums1 []int, m int, nums2 []int, n int)  {
    var arr []int
    for i, j := 0, 0; len(arr) != m+n;  {
        if (i >= m && j < n) {
            arr = append(arr, nums2[j])
            j += 1
        } else if (i < m && j >= n) {
            arr = append(arr, nums1[i])
            i += 1
        } else if (i < m && j < n){
            if (nums1[i] >= nums2[j]){
                arr = append(arr, nums2[j])
                j += 1
            } else {
                arr = append(arr, nums1[i])
                i += 1
            }
        }
    }
    for i := 0; i < m+n; i++ {
        nums1[i] = arr[i]
    }
}