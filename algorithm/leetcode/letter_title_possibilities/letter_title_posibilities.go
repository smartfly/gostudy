package letter_title_possibilities

func numTilePossibilities(tiles string) int {
	counts := make([]int, 26)
	for _, char := range tiles {
		counts[char-65]++
	}
	return dfs(counts)
}

// dfs 这个在递归分别计算a,b,c开头的
func dfs(counts []int) int {
	res := 0 // 返回值，组合个数
	for i := 0; i < 26; i++ {
		if counts[i] == 0 {
			continue
		}
		res++              // 返回结果加1
		counts[i]--        //消耗当前字符
		res += dfs(counts) //递归到下一层子问题解(消耗当前字符后的解)
		counts[i]++        //还原当前字符继续循环
	}
	return res
}
