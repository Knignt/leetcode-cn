class Solution {

    /**
     * @param Integer $num
     * @return Boolean
     */
    function isPerfectSquare($num) {
        if($num == 0 || $num == 1) return $num;

        $cnt = 1;
        $cursor = 1;
        while($cnt <= $num){
            if($cnt == $num){
                return true;
            }
            $cursor += 2;
            $cnt += $cursor;

            if($cnt  > $num){
                break;
            }
        }

        return false;
    }

    /**
    利用 1+3+5+7+9+…+(2n-1)=n^2，即完全平方数肯定是前n个连续奇数的和
     */
}
