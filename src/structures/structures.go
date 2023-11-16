package structures

type Person struct {
	FirstName string
	Lastname  string
	Age       int
	Phone     Phone
}

type Phone struct {
	Areacode string
	Prefix   string
	Suffix   string
}
