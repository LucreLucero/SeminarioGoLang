package candyshop

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

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

	list = append(list, &endpoint{ //GET All
		method:   "GET",
		path:     "/candies",
		function: getAll(s),
	})
	list = append(list, &endpoint{ //GET One by id
		method:   "GET",
		path:     "/candy/:id",
		function: getCandy(s),
	})
	list = append(list, &endpoint{ //Add one
		method:   "POST",
		path:     "/candy",
		function: addCandy(s),
	})
	list = append(list, &endpoint{ //Update one
		method:   "PUT",
		path:     "/candy/:id",
		function: updateCandy(s),
	})
	list = append(list, &endpoint{ //Delete one
		method:   "DELETE",
		path:     "/candy/:id",
		function: deleteCandy(s),
	})
	return list
}

//GET all the candies...
func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"candies": s.FindAll(),
		})
	}
}

//GET one candy...
func getCandy(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		c.JSON(http.StatusOK, gin.H{
			"candies": s.FindByID(id),
		})
	}
}

//ADD one candy...
func addCandy(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body) //para agregar necesito pasarle el cuerpo del objeto
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var candy Candy                                     //lo que tengo que agregar
		if err = json.Unmarshal(body, &candy); err != nil { //le paso el cuerpo y el valor en memoria de candy
			fmt.Println(err)
			os.Exit(1)
		}
		s.AddCandy(candy)
		c.JSON(http.StatusOK, gin.H{
			"Message": "Great! You have a new candy!",
		})
	}
}

//UPDATE one candy by id...
func updateCandy(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id")) //me traigo un objeto
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(c.Request.Body) //para agregar necesito pasarle el cuerpo del objeto
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var candy Candy                                     //lo que tengo que agregar
		if err = json.Unmarshal(body, &candy); err != nil { //le paso el cuerpo y el valor en memoria de candy
			fmt.Println(err)
			os.Exit(1)
		}
		s.Update(id, candy) //para actualizar tengo que pasarle el id del que modifico y el candy nuevo
		c.JSON(http.StatusOK, gin.H{
			"Message": "Great! You updated a candy!",
		})
	}
}

//DELETE one candy by id...
func deleteCandy(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id")) //me traigo un objeto
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		s.Delete(id) //para borrar tengo que pasarle el id del que quiero eliminar
		c.JSON(http.StatusOK, gin.H{
			"Message": "Oh! You deleted a candy :(",
		})

	}
}

//Register...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}

}
