package candyshop //el package siempre es el nombre de la carpeta
import (
	"fmt"

	"github.com/LucreLucero/SeminarioGoLang/internal/config"

	"github.com/jmoiron/sqlx"
)

//Candy...
type Candy struct {
	ID   int64
	Text string //Name
}

//Service...
type Service interface { //para que una estructura implemente una interfaz
	AddCandy(Candy) error //puedo usarlos como anonimos
	FindByID(int) *Candy
	FindAll() []*Candy

	Update(int, Candy)
	Delete(int)
}

type service struct { //no lo voy a exportar --esta en minuscula
	db     *sqlx.DB
	config *config.Config
}

//New...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

//Add one candy ...
func (s service) AddCandy(c Candy) error {
	query := "INSERT INTO candies (text) VALUES (?)"
	_, err := s.db.Exec(query, c.Text)
	if err != nil {
		fmt.Println(err.Error()) //imprimo el error --esta bien asi o debe ser un panic(err)
		return err
	}
	return nil
}

//Find one candy by id...
func (s service) FindByID(ID int) *Candy { //tengo que nombrarlos en las funciones
	var candy []*Candy                           //
	query := "SELECT * FROM candies WHERE id =?" //consulta a la db para que me traiga un objeto por id
	if err := s.db.Select(&candy, query, ID); err != nil {
		fmt.Println(err.Error()) //imprimo el error --esta bien asi o debe ser un panic(err)
	}
	return candy[0] //retorno el primer valor obtenido
}

//FindAll the candies...
func (s service) FindAll() []*Candy { // en base a la clase
	var list []*Candy
	if err := s.db.Select(&list, "SELECT * FROM candies"); err != nil {
		panic(err) //se puede usar siempre o cuando es mejor hacerlo ?
	}
	return list
}

//Update a candy...
func (s service) Update(ID int, c Candy) {
	query := "UPDATE candies SET text=?  WHERE id = ?"
	_, err := s.db.Exec(query, c.Text, ID)
	if err != nil {
		fmt.Println(err.Error()) //imprimo el error --en que se diferecia el panic(err) del error comun ??
	}
}

//Delete a candy...
func (s service) Delete(ID int) {
	query := "DELETE FROM candies WHERE id = ?"
	_, err := s.db.Exec(query, ID)
	if err != nil {
		fmt.Println(err.Error()) //imprimo el error --en que se diferecia el panic(err) del error comun ??
	}

}
