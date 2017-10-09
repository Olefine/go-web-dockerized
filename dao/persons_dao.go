package dao

// Person is a representation of DB entry
type Person struct {
	Firstname   string
	Lastname    string
	YearOfBirth uint16
}

var persons = []Person{
	newPerson("Egor", "Gorodov", uint16(1993)),
	newPerson("Oleg", "Papricko", uint16(1987)),
}

func newPerson(fn string, ln string, yob uint16) Person {
	return Person{fn, ln, yob}
}

// GetAll returns the all persons available
func GetAll() *[]Person {
	return &persons
}
