package mp

import(
	"fmt"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "- Tuomas Katajisto - \n This server contains all of my projects")
}
