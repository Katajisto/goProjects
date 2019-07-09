package mp

import(
	"fmt"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	fmt.Fprint(w, "Tuomas Katajisto")
	
=======
	fmt.Fprint(w, "|- Tuomas Katajisto -|")	
>>>>>>> 00d8ac71debbae0248bd4437fe5613e1b7e868f9
}
