package repo

import (
	"fmt"
	"github.com/spf13/cast"
	"seckshop/datasource"
	"seckshop/models"
	"time"
)


type ProductRepository interface {
	List(m map[string]interface{})(total int, products []models.Product)
	Insert(m map[string]interface{}) (goodId int64, err error)
	GetProduct(productId int64) (hasProduct bool, product models.Product, err error)
	SetAllSeckCount() bool
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
	product.ProductNum = cast.ToInt32(m["num"])
	product.ProductFirstImage = cast.ToString(m["first_image"])
	product.ProductImages = cast.ToString(m["images"])
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

func (p productRepository) GetProduct(productId int64) (hasProduct bool, product models.Product, err error) {
	hasProduct, err = engine.Where("id = ?", productId).Get(&product)
	if err != nil {
		panic("select error")
	}
	return
}

func (p productRepository) SetAllSeckCount() bool {
	product := new(models.Product)
	rows, err := engine.Rows(product)
	if err != nil {
		panic("select error")
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(product)
		ok := datasource.Redis.Set(cast.ToString(product.ID), product.ProductNum, time.Hour*time.Duration(24))
		fmt.Println(ok)
		//if ok.Val() != "ok" {
		//	return false
		//}
	}
	return true
}
