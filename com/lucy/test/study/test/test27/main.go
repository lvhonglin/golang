package main

func solve(board [][]byte) {
	mmap := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		mmap[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if mmap[i][j] {
				continue
			}
			if i == 0 || j == 0 || i == len(board)-1 || j == len(board[i])-1 {
				dfs(board, i, j, mmap)
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' && !mmap[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}
func dfs(board [][]byte, i, j int, mmap [][]bool) {
	if i < 0 || j < 0 || i > len(board)-1 || j > len(board[0])-1 {
		return
	}
	if mmap[i][j] {
		return
	}
	mmap[i][j] = true
	dfs(board, i+1, j, mmap)
	dfs(board, i-1, j, mmap)
	dfs(board, i, j+1, mmap)
	dfs(board, i, j-1, mmap)
}
