package table

type Data [][]string

type Table struct {
	data    Data
	indexed bool
	tWidth  int
	tHeight int
}

func New() *Table {
	return &Table{
		data:    Data{},
		indexed: false,
		tWidth:  80,
		tHeight: 20,
	}
}

func (t *Table) Print() {
}
