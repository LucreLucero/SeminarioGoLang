package candyshop

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HTTPService
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

//NewHTTPTransport
func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEnpoints(s)
	return httpService{endpoints}
}

func makeEnpoints(s Service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{ //All
		method:   "GET",
		path:     "/candies",
		function: getAll(s),
	})
	list = append(list, &endpoint{ //One
		method:   "GET",
		path:     "/candy",
		function: getOne(s),
	})
	list = append(list, &endpoint{ //Add
		method:   "POST",
		path:     "/candy",
		function: addCandy(s),
	})
	list = append(list, &endpoint{ //Update
		method:   "PUT",
		path:     "/candy/:id",
		function: updateCandy(s),
	})
	list = append(list, &endpoint{ //Delete
		method:   "DELETE",
		path:     "/candy/:id",
		function: deleteCandy(s),
	})
	return list
}

//GET all
func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"candies": s.FindAll(),
		})
	}
}

//GET one
func getOne(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

//ADD
func addCandy(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

//UPDATE
func updateCandy(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

//DELETE
func deleteCandy(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {}
}

//Register
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}

}
