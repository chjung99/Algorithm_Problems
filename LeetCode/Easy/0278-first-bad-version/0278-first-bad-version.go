/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    left := 0
    right := n + 1

    for left + 1 < right {
        mid := int((left + right) / 2)
        if (isBadVersion(mid)) {
            right = mid
        } else {
            left = mid
        }
    }
    return left + 1
}