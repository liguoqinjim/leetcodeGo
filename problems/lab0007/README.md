### [Reverse Integer](https://leetcode.com/problems/reverse-integer/#/description)
把一个数字反一下，比如321变成123、-456变成-654，但是要注意，还是超过了32位有符号数字的范围的话，就返回0

|方法|简介|说明|
|---|---|---|
|reverse||时间复杂度是O(n^2)|
|reverse2|优化|不用另一个slice参与计算，判断数字是否超过范围的判断也更简单|

#### 参考资料
##### 一行解决overflow的问题
https://discuss.leetcode.com/topic/6104/my-accepted-15-lines-of-code-for-java/20