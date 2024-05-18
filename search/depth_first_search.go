package search

const N = 250

var Mat [N][N]int
var Visit [N][N]int

func dfs(x, y int) {
	if Mat[x][y] != 1 || Visit[x][y] == 1 {
		return
	}
	Visit[x][y] = 1
	dfs(x-1, y-1)
	dfs(x-1, y)
	dfs(x-1, y+1)
	dfs(x+1, y-1)
	dfs(x+1, y)
	dfs(x+1, y+1)
}
