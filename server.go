package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/mygocode/prime/controller"
	"github.com/mygocode/prime/router"
	"github.com/mygocode/prime/service"
)

var (
	// postRepository repository.PostRepository = repository.NewFirestoreRepository()

	/* Since we are not interacting with database layer
	so we are not injecting any repository in service object */
	primeService service.PrimeService = service.NewPrimeService()

	//Injectig service object(PrimeService) to controller object (PrimeController)
	primeController controller.PrimeController = controller.NewPrimeController(primeService)
	muxRouter       router.Router              = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"

	muxRouter.GET("/", home)
	muxRouter.POST("/postprime", primeController.PostPrime)
	muxRouter.SERVE(port)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Up and running...")
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
