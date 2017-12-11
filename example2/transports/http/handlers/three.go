/*
* @author Colton J. McCurdy
*	GitHub: mccurdyc
* Email:  mccurdyc22@gmail.com
* Date: 2017-12-11
 */

package handlers

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

func Three(w http.ResponseWriter, r *http.Request) {
	p := response.NewPerson()

	// ParseForm updates r.Form with values
	if err := r.ParseForm(); err != nil {
		json.NewEncoder(w).Encode(p)
	}

	name := r.FormValue("name")
	// log.Printf("%+v", r)
	p = response.FindPersonByName(name, people)

	json.NewEncoder(w).Encode(p)
	return
}
