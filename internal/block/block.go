package block

import (
	"fmt"
	"strings"
	"math/rand"
)

var colors = map[int]int{
	2:1,
	4:1,
	8:2,
	16:2,
	32:3,
	64:3,
	128:4,
	256:4,
	512:5,
	1024:6,
	2048:7,
}

type Block struct {
	values [][]int
	size   int
}

func NewBlock(value [][]int, gameSize int) *Block {
	block := new(Block)
	block.values = value
	block.size = gameSize

	return block
}

func (b *Block) String() string {
	block := make([][][]string, b.size)
	for i := 0; i < b.size; i++ {
		block[i] = make([][]string, b.size)
	}
	var strArr []string

	for x, row := range b.values {
		for y, value := range row {
			block[x][y] = b.format(value)
		}
	}

	for _, row := range block {
		for i := 0; i < 5; i++ {
			var piece []string
			for j := 0; j < b.size; j++ {
				piece = append(piece, row[j][i])
			}
			strArr = append(strArr, strings.Join(piece, ""))
		}
	}

	return strings.Join(strArr, "\n")
}

func (b *Block) format(val int) []string {
	block := []string{"+----+", "|    |", "|    |", "|    |", "+----+"}
	if val == 0 {
		return block
	}
	reset := ColorSlice[0]
	colorIdx := colors[val]
	valueStr := fmt.Sprintf("%v%4d%v",ColorSlice[colorIdx],val,reset)
	valueDisplay := "|"+ valueStr + "|"
	block[2] = valueDisplay
	return block
}

func (b *Block) Update(newVal [][]int, gameSize int) {
	if gameSize != b.size {
		b.size = gameSize
	}
	b.values = newVal
	return
}

func compress(values []int) []int {
	var tmp []int
	for i := 0; i < len(values); i++ {
		if values[i] == 0 {
			continue
		}
		tmp = append(tmp, values[i])
	}

	return tmp
}

func merge(values []int) []int {
	var tmp []int

	for i := 0; i < len(values); i++ {
		if values[i] == 0 {
			continue
		} else if len(tmp) == 0 {
			tmp = append(tmp, values[i])
		} else if tmp[len(tmp)-1] == values[i] {
			tmp[len(tmp)-1] *= -2
		} else {
			tmp = append(tmp, values[i])
		}
	}
	for i := 0; i< len(tmp);i++ {
		if tmp[i] < 0 {
			tmp[i] = -1 * tmp[i]
		}
	}

	return tmp
}

func (b *Block) MoveLeft() {
	for i := 0; i < b.size; i++ {
		var tmp []int
		for j := 0; j < b.size; j++ {
			tmp = append(tmp,b.values[i][j])
			b.values[i][j] = 0
		}
		tmp = compress(tmp)
		tmp = merge(tmp)
		tmp = compress(tmp)
		for j := 0; j < len(tmp); j++ {
			b.values[i][j] = tmp[j]
		}
	}
	return
}

func (b *Block) MoveRight() {
	for i := 0; i < b.size; i++ {
		var tmp []int
		for j := b.size - 1; j >= 0; j-- {
			tmp = append(tmp,b.values[i][j])
			b.values[i][j] = 0
		}
		tmp = compress(tmp)
		tmp = merge(tmp)
		tmp = compress(tmp)
		for j := 0; j < len(tmp); j++ {
			b.values[i][b.size - 1 - j] = tmp[j]
		}
	}
	return
}
func (b *Block) MoveUp() {
	for i := 0; i < b.size; i++ {
		var tmp []int
		for j := 0; j < b.size; j++ {
			tmp = append(tmp,b.values[j][i])
			b.values[j][i] = 0
		}
		tmp = compress(tmp)
		tmp = merge(tmp)
		tmp = compress(tmp)
		for j := 0; j < len(tmp); j++ {
			b.values[j][i] = tmp[j]
		}
	}
	return
}
func (b *Block) MoveDown() {
	for i := 0; i < b.size; i++ {
		var tmp []int
		for j := b.size - 1; j >= 0; j-- {
			tmp = append(tmp,b.values[j][i])
			b.values[j][i] = 0
		}
		tmp = compress(tmp)
		tmp = merge(tmp)
		tmp = compress(tmp)
		for j := len(tmp) - 1; j >= 0; j-- {
			b.values[b.size - 1 - j][i] = tmp[j]
		}
	}
	return
}

func (b *Block) GenerateNewValue() bool {
	empty := 0
	for _,row := range b.values {
		for _,val := range row {
			if val == 0 {
				empty++
			}
		}
	}
	if empty == 0 {
		return false
	}

	pos := rand.Int() % empty
	newVal := 2
	if rand.Int() % 10 == 0 {
		newVal = 4
	}
	count := 0

	for i := 0 ; i < b.size;i++ {
		for j := 0; j < b.size;j++ {

			if b.values[i][j] == 0 {
				if count == pos {
					b.values[i][j] = newVal
				}
				count++
			}
		}
	}
	return true
}

func (b *Block) CheckWin() bool { 
	for _,row := range b.values {
		for _,val := range row {
			if val == 2048 {
				return true
			}
		}
	}
	return false
}
/*
+----+
|    |
|2048|
|    |
+----+

*/
