package candyshop //el package siempre es el nombre de la carpeta
import (
	"SeminarioGoLang/apiREST/internal/config"

	"github.com/jmoiron/sqlx"
)

//Candy
type Candy struct {
	ID   int64
	Text string
}

//CandyShopService
type CandyShopService interface { //para que una estructura implemente una interfaz
	AddCandy(Candy) error //puedo usarlos como anonimos
	FindByID(int) *CandyShopService
	FindAll() []*Candy
}

type service struct { //no lo voy a exportar --esta en minuscula
	db     *sqlx.DB
	config *config.Config
}

//New
func New(db *sqlx.DB, c *config.Config) (CandyShopService, error) {
	return service{db, c}, nil
}

func (s service) AddCandy(c Candy) error {
	return nil
}

func (s service) FindByID(ID int) *CandyShopService { //tengo que nombrarlos en las funciones
	return nil
}
func (s service) FindAll() []*Candy {
	var list []*Candy
	if err := s.db.Select(&list, "SELECT * FROM candies"); err != nil{
		panic(err)
	}
	return list
}
