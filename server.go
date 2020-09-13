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
	/* Since we are not interacting with database layer
	so we are not injecting any repository in service object */
	primeService service.PrimeService = service.NewPrimeService()

	//Injectig service object(PrimeService) to controller object (PrimeController)
	primeController controller.PrimeController = controller.NewPrimeController(primeService)
	muxRouter       router.Router              = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"
	//var port string = os.Getenv("PRIME_PORT")
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
