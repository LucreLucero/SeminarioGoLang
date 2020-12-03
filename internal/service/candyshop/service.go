package candyshop //el package siempre es el nombre de la carpeta
import (
	"github.com/LucreLucero/SeminarioGoLang/internal/config"

	"github.com/jmoiron/sqlx"
)

//Candy
type Candy struct {
	ID   int64
	Text string
}

//Service
type Service interface { //para que una estructura implemente una interfaz
	AddCandy(Candy) error //puedo usarlos como anonimos
	FindByID(int) *Service
	FindAll() []*Candy
}

type service struct { //no lo voy a exportar --esta en minuscula
	db     *sqlx.DB
	config *config.Config
}

//New
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddCandy(c Candy) error {
	return nil
}

func (s service) FindByID(ID int) *Service { //tengo que nombrarlos en las funciones
	return nil
}
func (s service) FindAll() []*Candy {
	var list []*Candy
	if err := s.db.Select(&list, "SELECT * FROM candies"); err != nil {
		panic(err)
	}
	return list
}
