package handlers

/*
* @author Colton J. McCurdy
*	GitHub: mccurdyc
* Email:  mccurdyc22@gmail.com
* Date: 2018-02-15
 */

import (
	"encoding/json"
	"net/http"

	"github.com/mccurdyc/go-examples/example2/transports/http/response"
)

var people = []response.Person{
	{Name: "Colton", Age: 22},
	{Name: "Dan", Age: 32},
	{Name: "Bob", Age: 54},
}

// Three will use a value passed in as a query paremeter to find a person by name
// in the slice of people.
func Three(w http.ResponseWriter, r *http.Request) {
	p := response.NewPerson()

	// ParseForm updates r.Form with values
	if err := r.ParseForm(); err != nil {
		json.NewEncoder(w).Encode(p)
	}

	name := r.FormValue("name")
	p = response.FindPersonByName(name, people)

	json.NewEncoder(w).Encode(p)
	return
}
