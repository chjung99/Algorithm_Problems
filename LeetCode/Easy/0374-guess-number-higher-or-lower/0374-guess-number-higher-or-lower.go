/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    num := n
    lb := 1
    ub := n
    for true {
        flag := guess(num)
        if (flag < 0) {
            ub = num
            num = (lb + num) / 2
        } else if (flag > 0) {
            lb = num
            num = (num + ub) / 2
        } else {
            break
        }
    }
    return num 
}