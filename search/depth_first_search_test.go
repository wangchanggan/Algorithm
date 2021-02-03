package search

import "testing"

func TestDFS(t *testing.T) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if Visit[j][j] != 1 && Mat[i][j] == 1 {
				dfs(i, j)
			}
		}
	}
}
