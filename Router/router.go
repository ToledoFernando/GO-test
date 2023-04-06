package rutas

import (
	"fmt"
	"net/http"
)

func Respuesta(w http.ResponseWriter, c *http.Request) {
	fmt.Fprint(w, "Hello world")
}
