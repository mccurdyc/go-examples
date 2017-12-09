/*
* @author Colton J. McCurdy
*	GitHub: mccurdyc
* Email:  mccurdyc22@gmail.com
* Date: 2017-12-11
 */

package response

type PersonResponse struct {
	data []Person `json:"data"`
}

type Person struct {
	Name string `json:"name"` // we will use this for encoding and decoding
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
