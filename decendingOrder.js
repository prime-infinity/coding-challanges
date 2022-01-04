/**
 *
CodeWars

Your task is to make a function that can take any
non-negative integer as an argument and return
it with its digits in descending order. 
Essentially, rearrange the digits to create the highest 
possible number.

Examples:
Input: 42145 Output: 54421

Input: 145263 Output: 654321

Input: 123456789 Output: 987654321
 */

function decendingOrder(n){

    /**
     * first of all, we turn the number into
     * a string and then to an array useing
     * either the array.from, or the map function
    */
    let convertToNum = num => Number(num)
    let nowArray = Array.from(String(n),convertToNum)
    
    //let nowArray2 = String(n).split("").map((item) => Number(item))

    /**
     * next, we have to arrange it in decending order
     */
   
}

decendingOrder(85738263)