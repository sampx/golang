package advance

import (
	"encoding/json"
	"fmt"
)

// Product 商品信息
type Product struct {
	Name      string  `json:"name,string"`                 //json输出name,type string
	ProductID int64   `json:"product_id,string,omitempty"` //"-"表示不进行序列化,omitempty表示忽略0值或者空值
	Stock     int     `json:"stock,string"`
	Price     float64 `json:"price,number"`
	IsOnSale  bool    `json:"is_on_sale,bool"` //改名
}

func JsonTest() {
	p := &Product{}
	p.Name = "Xiao mi 6"
	p.IsOnSale = true
	p.Stock = 10000
	p.Price = 2499.00
	p.ProductID = 1
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
}
