/*
* @author Colton J. McCurdy
*	GitHub: mccurdyc
* Email:  mccurdyc22@gmail.com
* Date: 2017-12-11
 */

package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// Two prints "hello, {name}" where {name} is a parameter
// that is passed in by the caller
func Two(w http.ResponseWriter, r *http.Request) {

	// a map of route variables
	vars := mux.Vars(r)
	s := fmt.Sprintf("hello, %s", vars["name"])

	io.WriteString(w, s)
}
