package chessboard

type File []bool

type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	
	if col, ok := cb[file]; !ok{
		return 0
	} else {
		count := 0
		for _, occupied := range col {
			if occupied {
				count++
			}
		}
		return count
	}
}

func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	idx := rank - 1
	count := 0
	for _, file := range cb {
		if file[idx] {
			count++
		}
	}
	return count
}

func CountAll(cb Chessboard) int {
	count := 0
	for range cb {
		count += 8
	}
	return count
}

func CountOccupied(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		for _, occupied := range file {
			if occupied {
				count++
			}
		}
	}
	return count
}