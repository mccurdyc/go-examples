package response

/*
* @author Colton J. McCurdy
*	GitHub: mccurdyc
* Email:  mccurdyc22@gmail.com
* Date: 2018-02-15
 */

// PersonResponse contains a response with a slice of Person types.
type PersonResponse struct {
	data []Person `json:"data"`
}

// Person is a struct that will hold basic information about a person i.e., name and age.
type Person struct {
	Name string `json:"name"` // we will use this for encoding and decoding
	Age  int    `json:"age"`
}

// NewPerson will create a new Person with the nil values of Name and Age.
func NewPerson() *Person {
	return &Person{}
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
