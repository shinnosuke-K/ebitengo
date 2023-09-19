package main

type Arbiter struct {
	board     Board
	direction [8][2]int
	end       bool
}

func newArbiter(board Board) *Arbiter {
	return &Arbiter{
		board: board,
		direction: [8][2]int{
			{1, 0},   // 右
			{1, 1},   // 右下
			{0, 1},   // 下
			{-1, 1},  // 左下
			{-1, 0},  // 左
			{-1, -1}, // 左上
			{0, -1},  // 上
			{1, -1},  // 右上
		},
	}
}

func (a *Arbiter) IsEndGame() bool { return a.end }

// Judge は勝敗を判定する
func (a *Arbiter) Judge(turn Piece) Piece {
	if a.CanPlace(turn) || a.CanPlace(turn.Opponent()) {
		return PieceNone
	}
	black := 0
	white := 0
	for _, row := range a.board {
		for _, piece := range row {
			switch piece {
			case PieceBlack:
				black++
			case PieceWhite:
				white++
			}
		}
	}
	if black > white {
		a.end = true
		return PieceBlack
	} else if black < white {
		a.end = true
		return PieceWhite
	} else {
		return PieceNone
	}
}

// CountPiece は盤面上の指定した色の石の数を数える
func (a *Arbiter) CountPiece(piece Piece) int {
	count := 0
	for _, row := range a.board {
		for _, p := range row {
			if p == piece {
				count++
			}
		}
	}
	return count
}

// CanPlace プレイヤーが駒を置ける場所があるかどうかをチェックする関数
// 1. 盤面の全てのマスに対して、プレイヤーの駒が置けるかどうかをチェックする
// 2. 1つでも置ける場所があればtrueを返す
// 3. 1つも置ける場所がなければfalseを返す
func (a *Arbiter) CanPlace(piece Piece) bool {
	for y, row := range a.board {
		for x := range row {
			if a.CanSet(x, y, piece) {
				return true
			}
		}
	}
	return false
}

func (a *Arbiter) CanSet(x, y int, piece Piece) bool {
	if a.board[y][x] != PieceNone {
		return false
	}
	for _, d := range a.direction {
		if a.checkDirection(x, y, d[0], d[1], piece) {
			return true
		}
	}
	return false
}

// checkDirection は指定した方向に相手の駒があるかどうかをチェックする関数
// 1. 指定した方向に相手の駒があるかどうかをチェックする
// 2. 相手の駒があればtrueを返す
// 3. 相手の駒がなければfalseを返す
func (a *Arbiter) checkDirection(x, y, dx, dy int, piece Piece) bool {
	// 1つ目のマスをチェック
	x += dx
	y += dy

	// 盤面の範囲外に出たらfalseを返す
	if x < 0 || x > len(a.board)-1 || y < 0 || y > len(a.board)-1 {
		return false
	}
	if a.board[y][x] != piece.Opponent() {
		return false
	}
	// 2つ目以降のマスをチェック
	for {
		x += dx
		y += dy

		// 盤面の範囲外に出たらfalseを返す
		if x < 0 || x > len(a.board)-1 || y < 0 || y > len(a.board)-1 {
			return false
		}

		switch a.board[y][x] {
		case PieceNone:
			return false
		case piece:
			return true
		}
	}
}

func (a *Arbiter) PlaceAndFlip(x, y int, piece Piece) {
	a.board[y][x] = piece
	for _, d := range a.direction {
		a.flipDirection(x, y, d[0], d[1], piece)
	}
}

func (a *Arbiter) flipDirection(x, y, dx, dy int, piece Piece) {
	// 1つ目のマスをチェック
	x += dx
	y += dy

	// 盤面の範囲外に出たらfalseを返す
	if x < 0 || x > len(a.board)-1 || y < 0 || y > len(a.board)-1 {
		return
	}
	if a.board[y][x] != piece.Opponent() {
		return
	}
	// 2つ目以降のマスをチェック
	for i := 0; i < len(a.board); i++ {
		x += dx
		y += dy

		// 盤面の範囲外に出たらfalseを返す
		if x < 0 || x > len(a.board)-1 || y < 0 || y > len(a.board)-1 {
			return
		}

		switch a.board[y][x] {
		case PieceNone:
			return
		case piece:
			for j := 0; j < i+1; j++ {
				x -= dx
				y -= dy
				if a.board[y][x] == piece {
					return
				}
				a.board.Set(x, y, piece)
			}
		}
	}
}
