type ValidBoard [][]int64

func (v ValidBoard) HasCompleteColumns() bool {
    for col := range v[0] {
        set := make(map[int64]struct{})
        for row := range v {
            num := v[row][col]
            if num == 0 {
                continue
            }
            if _, exists := set[num]; exists {
                return false
            }
            set[num] = struct{}{}
        }
    }
    return true
}

func (v ValidBoard) HasCompleteBoxes() bool {
    for boxRow := 0; boxRow < 9; boxRow += 3 {
        for boxCol := 0; boxCol < 9; boxCol += 3 {
            if !v.isBoxValid(boxRow, boxCol) {
                return false
            }
        }
    }
    return true
}

func (v ValidBoard) isBoxValid(boxRow, boxCol int) bool {
    set := make(map[int64]struct{})
    for row := boxRow; row < boxRow+3; row++ {
        for col := boxCol; col < boxCol+3; col++ {
            num := v[row][col]
            if num == 0 {
                continue 
            }
            if _, exists := set[num]; exists {
                return false 
            }
            set[num] = struct{}{}
        }
    }
    return true
}

func (v ValidBoard) HasCompleteRows() bool {
    for row := range v {
        set := make(map[int64]struct{})
        for col := range v[row] {
            num := v[row][col]
            if num == 0 {
                continue 
            }
            if _, exists := set[num]; exists {
                return false
            }
            set[num] = struct{}{}
        }
    }
    return true
}
  
func isValidSudoku(board [][]byte) bool {
     v, ok := NewValidBoard(board)
     if !ok {
         return false
     }

     return v.HasCompleteColumns() && v.HasCompleteRows() && v.HasCompleteBoxes()
}

func NewValidBoard(board [][]byte) (ValidBoard, bool) {
    b := make(ValidBoard, 9)
    for i := range b {
        b[i] = make([]int64, 9)
        for j := range b[i] {
            val :=  board[i][j]
            if  val == '.' {
                b[i][j] = 0
                continue
            }
            conv, err := strconv.ParseInt(string(val), 10, 64) 
            if err != nil {
                return ValidBoard{}, false
            }

            if conv < 1 || conv > 9 {
                return ValidBoard{}, false
            } 
            b[i][j] = conv 
        }
    }

    return b, true
}
