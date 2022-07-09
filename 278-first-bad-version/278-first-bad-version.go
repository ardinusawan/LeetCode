/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    // binary search iterative
    
    left, right := 1, n
    
    for left <= right {
        pivot := left + (right - left)/2
        
        if isBadVersion(pivot) == true && pivot == 1 {
            return 1
        } else if isBadVersion(pivot) == true && isBadVersion(pivot - 1) == false {
            return pivot
        }
        
        if isBadVersion(pivot) == false {
            left = pivot + 1    
        } else {
            right = pivot - 1
        }
    }
    
    return 1
}