# Minimum Difference Sum
Given an array of n integers, rearrange them so that the sum of the
absolute differences of all adjacent elements is minimized. Then,
compute the sum of those absolute differences.

### Example
    n= 5
    arr = [1, 3, 3, 2, 4]

If the list is rearranged as arr' = [1, 2, 3, 3, 4], the absolute
differences are |1 -2| = 1, |2-3| = 1, |3-3| = 0, |3 - 4| = 1. The
sum of those differences is 1 + 1 + 0 + 1 = 3.

### Function Description
Complete the function minDiff in the editor below.

minDiff has the following parameter:

    arr: an integer array

Returns:

int: the sum of the absolute differences of adjacent elements

### Constraints

- 2 <= n <= 10^5
- O <= arr[i] <= 10^9, where O <= i < n

### Input Format For Custom Testing
The first line of input contains an integer, n, the size of arr.
Each of the following n lines contains an integer that describes
arr[i] (where O < i < n).

### Sample Case 0
Sample Input For Custom Testing

````
STDIN	Function
-----   --------
5     -> arr[] size n = 5
5	  -> arr[] = [5, 1, 3, 7, 3]
1
3
7
3
````

#### Sample Output
6

#### Explanation
n= 5
arr = [5, 1, 3, 7, 3]

If arr is rearranged as arr' = [1, 3, 3, 5, 71, the differences are
minimized.

The final answer is |1-3|+ |3-3| +|3-5|+|5-7|=6.

### Sample Case 1
Sample Input For Custom Testing
````
STDIN  Function
-----  --------
2    -> arr[] size n = 2
3	 -> arr[] = [3, 2]
2
````
#### Sample Output
1
#### Explanation
n = 2
arr = [3, 2]

There is no need to rearrange because there are only two
elements. The final answer is |3-2| = 1.