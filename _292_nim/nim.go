package _292_nim

/*
** 解题思路：
** 每次最多只能拿 3 个石头
** 石头个数    先手    对手    剩余石头   成功/失败
**    1         1      0         0        win
**    2         2      0         0        win
**    3         3      0         0        win
**    4         1      3         0        failed
**    4         2      2         0        failed
**    4         3      1         0        failed
**    5         1      *         4        win
**    6         2      *         4        win
**    7         3      *         4        win
**    8         1      3         4        failed
**    8         2      2         4        failed
**    8         3      1         4        failed
**    9         1      *         8        win
**
** 从上面的递推可以发现，只要留给对手的石头个数为 4 的倍数就可以赢
** 代码如下：
 */

func canWinNim(n int) bool {
	return n%4 == 0
}
