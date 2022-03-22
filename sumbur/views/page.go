package views

type MenuItem struct {
	link  []byte
	label []byte
}

var Menu = []MenuItem{
	{[]byte(""), []byte("Сумбур")},
	{[]byte("things"), []byte("Всячина")},
}

const (
	PageBlog = iota
	PageThings
)
