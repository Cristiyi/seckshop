package repo

import (
	"seckshop/datasource"
	"seckshop/models"
	"github.com/spf13/cast"
)

type ProductRepository interface {
	List(m map[string]interface{})(total int, products []models.Product)
}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

var engine = datasource.GetDB()

type productRepository struct {

}

func (p productRepository) List(m map[string]interface{})(total int, products []models.Product) {
	err := engine.Limit(cast.ToInt(m["size"]), (cast.ToInt(m["page"])-1)*cast.ToInt(m["size"])).Find(&products)
	if err != nil {
		panic("select error")
	}
	total = len(products)
	return
}
