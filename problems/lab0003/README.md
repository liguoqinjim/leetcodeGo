# [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

## NOTICE
 - 一开始做的是最简单的遍历
 - 优化：`abcabcbb`,到第二个a的时候，可以不用从第二位开始，而是把a的index移动到3，再加上一个起始位计算

## 参考资料
 - https://leetcode.com/problems/longest-substring-without-repeating-characters/discuss/1731/A-Python-solution-85ms-O(n)