package repo

import (
	"seckshop/datasource"
	"seckshop/models"
	"github.com/spf13/cast"
)

type ProductRepository interface {
	List(m map[string]interface{})(total int, products []models.Product)
	Insert(m map[string]interface{}) (goodId int64, err error)
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

func (p productRepository) Insert(m map[string]interface{}) (goodId int64, err error) {
	product := new(models.Product)
	product.ProductName = cast.ToString(m["product_name"])
	product.ProductNum = cast.ToInt8(m["num"])
	product.ProductFirstImage = cast.ToString(m["first_image"])
	if cast.ToString(m["product_url"]) != "" {
		product.ProductUrl = cast.ToString(m["product_url"])
	}
	_, err = engine.Insert(product)
	if err != nil {
		panic("insert error")
	}
	goodId = product.ID
	return
}
