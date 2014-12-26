package dice

import (
	"net/http"
)

func GetDie(rw http.ResponseWriter, r *http.Request) {

	rw.Write([]byte("test"))
}
