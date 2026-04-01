type ValidBoard [][]int64 

func NewValidBoard(board [][]byte) (ValidBoard, bool) {
    b := make(ValidBoard, 9)
    rows := [9]map[int64]struct{}{}
    cols := [9]map[int64]struct{}{}
    boxes := [9]map[int64]struct{}{}

    for i := range rows {
        rows[i] = make(map[int64]struct{})
        cols[i] = make(map[int64]struct{})
        boxes[i] = make(map[int64]struct{})
    }

    for i := range b {
        b[i] = make([]int64, 9)
        for j := range b[i] {
            val := board[i][j]
            if val == '.' {
                continue
            }
            conv, err := strconv.ParseInt(string(val), 10, 64)
            if err != nil {
                return ValidBoard{}, false
            }

            boxIdx := (i/3)*3 + j/3

            for _, set := range []map[int64]struct{}{rows[i], cols[j], boxes[boxIdx]} {
                if _, exists := set[conv]; exists {
                    return ValidBoard{}, false
                }
                set[conv] = struct{}{}
            }

            b[i][j] = conv
        }
    }

    return b, true
}

func isValidSudoku(board [][]byte) bool {
	_, isValid := NewValidBoard(board)
	return isValid
}