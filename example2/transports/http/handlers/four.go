/*
* @author Colton J. McCurdy
*	GitHub: mccurdyc
* Email:  mccurdyc22@gmail.com
* Date: 2017-12-11
 */

package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mccurdyc/go-examples/example2/transports/http/response"
)

// Four expects a POST with a Person
// in the request body. It then takes this
// and adds it to the global array of people.
func Four(w http.ResponseWriter, r *http.Request) {
	var p response.Person

	fmt.Printf("Request: %+v\n", r)

	// create a new json decoder
	decoder := json.NewDecoder(r.Body)

	// wait until the surrounding function returns
	// defer req.Body.Close()

	// decode into a struct
	if err := decoder.Decode(&p); err != nil {
		// this is also another way you will commonly see error handling
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Person: %+v\n", p)

	r.Body.Close() // we can also use a defer instead

	w.WriteHeader(http.StatusOK)
	return
}
