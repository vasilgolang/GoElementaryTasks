package task1

func ChessBoard(width, height int, symbol rune) (board string) {
	if width < 0 || height < 0 {
		return ""
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if (i+j)%2 == 0 {
				board += string(symbol)
			} else {
				board += ` `
			}
		}
		board += "\r\n"
	}
	
	return
}
