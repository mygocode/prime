package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/mygocode/prime/controller"
	"github.com/mygocode/prime/router"
	"github.com/mygocode/prime/service"
)

var (
	primeCache = map[uint32]uint32{0: 0}

	/* Creating and injecting local cache.
	If we had database layer then we would have injected here in service */
	primeService service.PrimeService = service.NewPrimeService(primeCache)

	//Injectig service object(PrimeService) to controller object (PrimeController)
	primeController controller.PrimeController = controller.NewPrimeController(primeService)

	muxRouter router.Router = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"

	muxRouter.GET("/", home)
	muxRouter.POST("/postprime", primeController.PostPrime)
	muxRouter.SERVE(port)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/home.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, nil); err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
}
