package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

type Reversi struct {
	arbiter      *Arbiter
	screenWidth  int
	screenHeight int
	turn         Piece
	players      map[Piece]*ManualPlayer
	winner       *ManualPlayer
	looser       *ManualPlayer
}

func newReversi() *Reversi {
	return &Reversi{
		arbiter:      newArbiter(newBoard()),
		screenWidth:  800,
		screenHeight: 480,
		turn:         PieceBlack,
		players: map[Piece]*ManualPlayer{
			PieceBlack: newManualPlayer("Player1", PieceBlack, colornames.Red),
			PieceWhite: newManualPlayer("Player2", PieceWhite, colornames.Blue),
		},
	}
}

func (r *Reversi) Title() string { return "Reversi" }

func (r *Reversi) WindowSize() (width, height int) { return r.screenWidth, r.screenHeight }

func (r *Reversi) Layout(_, _ int) (screenWidth, screenHeight int) {
	return r.screenWidth, r.screenHeight
}

func (r *Reversi) Update() error {
	if r.arbiter.IsEndGame() {
		return nil
	}
	if winner := r.arbiter.Judge(r.turn); winner != PieceNone {
		r.winner = r.players[winner]
		r.looser = r.players[winner.Opponent()]
		return nil
	}

	nextTurn := r.turn.Opponent()

	// カーソルの移動
	r.players[r.turn].MoveCursor()

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft):
		r.players[r.turn].ResetColor()
		return nil
	case !inpututil.IsKeyJustPressed(ebiten.KeyA):
		return nil
	}

	if !r.arbiter.CanSet(r.players[r.turn].CursorX, r.players[r.turn].CursorY, r.turn) {
		r.players[r.turn].SetColor(colornames.Yellow)
		return nil
	}
	r.arbiter.PlaceAndFlip(r.players[r.turn].CursorX, r.players[r.turn].CursorY, r.turn)

	// プレイヤーのターン
	if !r.arbiter.CanPlace(nextTurn) {
		r.players[r.turn].ResetCursor()
		return nil
	}
	r.turn = nextTurn
	r.players[r.turn].ResetCursor()
	r.players[PieceBlack].SetCount(r.arbiter.CountPiece(PieceBlack))
	r.players[PieceWhite].SetCount(r.arbiter.CountPiece(PieceWhite))
	return nil
}

func (r *Reversi) Draw(screen *ebiten.Image) {
	// 盤面を描画
	r.arbiter.board.DrawPlate(screen)
	// 盤面のマスを描画
	r.arbiter.board.DrawBoard(screen)

	// プレイヤー名を描画
	r.players[PieceBlack].DrawDisplayName(screen)
	r.players[PieceWhite].DrawDisplayName(screen)
	// プレイヤーの石の数を描画
	r.players[PieceBlack].DrawPieceCount(screen)
	r.players[PieceWhite].DrawPieceCount(screen)
	// カーソルを描画
	r.players[r.turn].DrawCursor(screen)

	if r.arbiter.IsEndGame() {
		r.winner.DrawWinner(screen)
		r.looser.DrawLooser(screen)
	}
}
