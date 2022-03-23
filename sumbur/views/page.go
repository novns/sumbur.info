package views

// Menu

type MenuItem struct {
	link  []byte
	label []byte
}

const (
	PageBlog = iota
	PageThings
)

var Menu = []MenuItem{
	{[]byte(""), []byte("Сумбур")},
	{[]byte("things"), []byte("Всячина")},
}

// Auth

const (
	AuthDefault = iota
	AuthOK
	AuthFail
)

var AuthClass = [][]byte{
	[]byte("default"),
	[]byte("ok"),
	[]byte("fail"),
}

var AuthState int = AuthDefault
