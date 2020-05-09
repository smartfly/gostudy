package regular_expression_matching

/*
 * @author taohu
 * @date 2020/5/4
 * @time 上午11:53
 * @desc 正则表达式匹配
**/

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}

	sLen, pLen := len(s), len(p)
	dp := make([][]bool, sLen+1)
	for i := 0; i < sLen+1; i++ {
		dp[i] = make([]bool, pLen+1)
	}
	// init
	dp[0][0] = true
	// deal with pattern like a*, a*b* or a*b*c*
	for i := 2; i < pLen+1; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}

	for i := 1; i < sLen+1; i++ {
		for j := 1; j < pLen+1; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if s[i-1] != p[j-2] && p[j-2] != '.' {
					dp[i][j] = dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-2] || dp[i][j-1] || dp[i-1][j]
				}
			}
		}
	}
	return dp[sLen][pLen]
}
