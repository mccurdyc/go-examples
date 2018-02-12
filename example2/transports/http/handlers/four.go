package handlers

/*
* @author Colton J. McCurdy
*	GitHub: mccurdyc
* Email:  mccurdyc22@gmail.com
* Date: 2018-02-15
 */

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mccurdyc/go-examples/example2/transports/http/response"
)

// Four expects a POST with a Person
// in the request body. It then takes this
// and adds it to the global array of people.
func Four(w http.ResponseWriter, r *http.Request) {
	p := response.NewPerson()

	// create a new json decoder
	decoder := json.NewDecoder(r.Body)

	// wait until the surrounding function returns
	defer r.Body.Close()

	// decode into a struct
	if err := decoder.Decode(&p); err != nil {
		// this is also another way you will commonly see error handling
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Person: %+v\n", p)

	w.WriteHeader(http.StatusOK)
	return
}
