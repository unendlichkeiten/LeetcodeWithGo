#### 解题思路
- 题目：[1668. 最大重复子字符串](https://leetcode.cn/problems/maximum-repeating-substring/description/)
- 易错点：开始看题目的时候，没有注意连续重复几个字的描述，直接使用模糊匹配解答了
- 解题思路：
  1. 计算每个位置开始出现连续字符串的个数，一旦匹配失败就终止匹配逻辑
  2. 统计所有位置出现重复字符串的次数，选择最大值
