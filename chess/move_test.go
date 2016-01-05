package chess

import (
	"log"
	"testing"
)

func moveIsValid(pos *Position, m *Move) bool {
	for _, move := range pos.ValidMoves() {
		if move.s1 == m.s1 && move.s2 == m.s2 && move.promo == m.promo {
			return true
		}
	}
	return false
}

type moveTest struct {
	pos *Position
	m   *Move
}

var (
	validMoves = []moveTest{
		// pawn moves
		{m: &Move{s1: E2, s2: E4}, pos: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		{m: &Move{s1: A2, s2: A3}, pos: unsafeFEN("8/8/8/8/8/8/P7/8 w - - 0 1")},
		{m: &Move{s1: A7, s2: A6}, pos: unsafeFEN("8/p7/8/8/8/8/8/8 b - - 0 1")},
		{m: &Move{s1: A7, s2: A5}, pos: unsafeFEN("8/p7/8/8/8/8/8/8 b - - 0 1")},
		{m: &Move{s1: C4, s2: B5}, pos: unsafeFEN("8/8/8/1p1p4/2P5/8/8/8 w - - 0 1")},
		{m: &Move{s1: C4, s2: D5}, pos: unsafeFEN("8/8/8/1p1p4/2P5/8/8/8 w - - 0 1")},
		{m: &Move{s1: C4, s2: C5}, pos: unsafeFEN("8/8/8/1p1p4/2P5/8/8/8 w - - 0 1")},
		{m: &Move{s1: C5, s2: B4}, pos: unsafeFEN("8/8/8/2p5/1P1P4/8/8/8 b - - 0 1")},
		{m: &Move{s1: C5, s2: D4}, pos: unsafeFEN("8/8/8/2p5/1P1P4/8/8/8 b - - 0 1")},
		{m: &Move{s1: C5, s2: C4}, pos: unsafeFEN("8/8/8/2p5/1P1P4/8/8/8 b - - 0 1")},
		{m: &Move{s1: A4, s2: B3}, pos: unsafeFEN("2r3k1/1q1nbppp/r3p3/3pP3/pPpP4/P1Q2N2/2RN1PPP/2R4K b - b3 0 23")},
		{m: &Move{s1: A2, s2: A1, promo: Queen}, pos: unsafeFEN("8/8/8/8/8/8/p7/8 b - - 0 1")},
		// knight moves
		{m: &Move{s1: E4, s2: F6}, pos: unsafeFEN("8/8/8/3pp3/4N3/8/5B2/8 w - - 0 1")},
		{m: &Move{s1: E4, s2: D6}, pos: unsafeFEN("8/8/8/3pp3/4N3/8/5B2/8 w - - 0 1")},
		{m: &Move{s1: E4, s2: C3}, pos: unsafeFEN("8/8/8/3pp3/4N3/8/5B2/8 w - - 0 1")},
		// bishop moves
		{m: &Move{s1: E4, s2: H7}, pos: unsafeFEN("8/8/8/3pp3/4B3/5N2/8/8 w - - 0 1")},
		{m: &Move{s1: E4, s2: D5}, pos: unsafeFEN("8/8/8/3pp3/4B3/5N2/8/8 w - - 0 1")},
		{m: &Move{s1: E4, s2: B1}, pos: unsafeFEN("8/8/8/3pp3/4B3/5N2/8/8 w - - 0 1")},
		// rook moves
		{m: &Move{s1: B2, s2: B4}, pos: unsafeFEN("8/1p5b/4N3/4p3/8/8/1R6/1B6 w - - 0 1")},
		{m: &Move{s1: B2, s2: B7}, pos: unsafeFEN("8/1p5b/4N3/4p3/8/8/1R6/1B6 w - - 0 1")},
		{m: &Move{s1: B2, s2: A2}, pos: unsafeFEN("8/1p5b/4N3/4p3/8/8/1R6/1B6 w - - 0 1")},
		{m: &Move{s1: B2, s2: H2}, pos: unsafeFEN("8/1p5b/4N3/4p3/8/8/1R6/1B6 w - - 0 1")},
	}

	invalidMoves = []moveTest{
		// out of turn moves
		{m: &Move{s1: E7, s2: E5}, pos: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		{m: &Move{s1: E2, s2: E4}, pos: unsafeFEN("rnbqkbnr/1ppppppp/p7/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1")},
		// pawn moves
		{m: &Move{s1: E2, s2: D3}, pos: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		{m: &Move{s1: E2, s2: F3}, pos: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		{m: &Move{s1: E2, s2: E5}, pos: unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")},
		{m: &Move{s1: A2, s2: A1}, pos: unsafeFEN("8/8/8/8/8/8/p7/8 b - - 0 1")},
		{m: &Move{s1: E6, s2: E5}, pos: unsafeFEN(`2b1r3/2k2p1B/p2np3/4B3/8/5N2/PP1K1PPP/3R4 b - - 2 1`)},
		{m: &Move{s1: H7, s2: H5}, pos: unsafeFEN(`2bqkbnr/rpppp2p/2n2p2/p5pB/5P2/4P3/PPPP2PP/RNBQK1NR b KQk - 4 6`)},
		// knight moves
		{m: &Move{s1: E4, s2: F2}, pos: unsafeFEN("8/8/8/3pp3/4N3/8/5B2/8 w - - 0 1")},
		{m: &Move{s1: E4, s2: F3}, pos: unsafeFEN("8/8/8/3pp3/4N3/8/5B2/8 w - - 0 1")},
		// bishop moves
		{m: &Move{s1: E4, s2: C6}, pos: unsafeFEN("8/8/8/3pp3/4B3/5N2/8/8 w - - 0 1")},
		{m: &Move{s1: E4, s2: E5}, pos: unsafeFEN("8/8/8/3pp3/4B3/5N2/8/8 w - - 0 1")},
		{m: &Move{s1: E4, s2: E4}, pos: unsafeFEN("8/8/8/3pp3/4B3/5N2/8/8 w - - 0 1")},
		{m: &Move{s1: E4, s2: F3}, pos: unsafeFEN("8/8/8/3pp3/4B3/5N2/8/8 w - - 0 1")},
		// rook moves
		{m: &Move{s1: B2, s2: B1}, pos: unsafeFEN("8/1p5b/4N3/4p3/8/8/1R6/1B6 w - - 0 1")},
		{m: &Move{s1: B2, s2: C3}, pos: unsafeFEN("8/1p5b/4N3/4p3/8/8/1R6/1B6 w - - 0 1")},
		{m: &Move{s1: B2, s2: B8}, pos: unsafeFEN("8/1p5b/4N3/4p3/8/8/1R6/1B6 w - - 0 1")},
		{m: &Move{s1: B2, s2: G7}, pos: unsafeFEN("8/1p5b/4N3/4p3/8/8/1R6/1B6 w - - 0 1")},
	}
)

func unsafeFEN(s string) *Position {
	pos, err := decodeFEN(s)
	if err != nil {
		log.Fatal(err)
	}
	return pos
}

func TestValidMoves(t *testing.T) {
	for _, mt := range validMoves {
		if !moveIsValid(mt.pos, mt.m) {
			log.Println(mt.pos.String())
			log.Println(mt.pos.board.Draw())
			log.Println(mt.pos.ValidMoves())
			t.Fatalf("expected move %s to be valid", mt.m)
		}
	}
}

func TestInvalidMoves(t *testing.T) {
	for _, mt := range invalidMoves {
		if moveIsValid(mt.pos, mt.m) {
			log.Println(mt.pos.String())
			log.Println(mt.pos.board.Draw())
			t.Fatalf("expected move %s to be invalid", mt.m)
		}
	}
}

func BenchmarkValidMoves(b *testing.B) {
	pos := unsafeFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		pos.ValidMoves()
	}
}
