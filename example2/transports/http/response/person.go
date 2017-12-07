package response

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewPerson() *Person {
	return &Person{
		Name: "",
		Age:  0,
	}
}

// FindPersonByName returns a Person if a Person with
// that name exists in the slice of Person provided
// to search through
func FindPersonByName(name string, people []Person) *Person {
	for _, p := range people {
		// case sensitive comparison
		if p.Name == name {
			return &p
		}
	}
	return nil
}
