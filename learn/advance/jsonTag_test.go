package advance

import (
	"encoding/json"
	"fmt"
	"testing"
)

type prod struct {
	Name      string  `json:"name"`           //json输出name,type string
	Desc      string  `json:"desc,omitempty"` //omitempty表示忽略0值或者空值
	ProductID int64   `json:"product_id,string"`
	GoodsID   int64   `json:"goods_id"`
	Price     float64 `json:"price,string"` //string 以string类型输出、读入

	//Additional
	Additional string `json:"-"` //- 不输出字段
}

func TestMarshal(t *testing.T) {
	var prduct = &prod{
		Name:      "音响",
		Desc:      "",
		ProductID: 103131441,
		GoodsID:   210313140,
		Price:     99.99,

		Additional: "不输出Json段",
	}

	str, err := json.Marshal(prduct)
	if err != nil {
		t.Logf("Marshal Err[%v]", err) //打印日志并终止测试
	}

	//Log	打印日志，同时结束测试
	//Logf	格式化打印日志，同时结束测试
	//Error	打印错误日志，同时结束测试
	//Errorf	格式化打印错误日志，同时结束测试
	//Fatal	打印致命日志，同时结束测试
	//Fatalf	格式化打印致命日志，同时结束测试
	t.Log(string(str))
}

type JExchange struct {
	E    string    `json:"e,string"`
	Data []float64 `json:"data,omitempty"`
	Pair string    `json:"pair,omitempty"`
}

func TestUnMarshal(t *testing.T) {
	var Unpack JExchange

	var strByte = []byte(`{"e":"\"ohlcv24\"","data":[7469.6,7541.1,7156,7213.4,64856543672],"pair":"BTC:USD"}`)

	Unpack = JExchange{}
	err := json.Unmarshal(strByte, &Unpack)

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(Unpack.E, Unpack.Pair, Unpack.Data)
}

func BenchmarkUnMarshal(b *testing.B) {
	var Unpack JExchange
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var strByte = []byte(`{"e":"\"ohlcv24\"","data":[7469.6,7541.1,7156,7213.4,64856543672],"pair":"BTC:USD"}`)

		Unpack = JExchange{}
		err := json.Unmarshal(strByte, &Unpack)

		if err != nil {
			b.Error(err.Error())
		}
	}
	b.StopTimer()
	fmt.Println(Unpack.E, Unpack.Pair)
}

func TestFailNow(t *testing.T) {
	t.FailNow() //终止当前测试用例
}

func TestFail(t *testing.T) {
	println("before fail") //不能使用t.Log,会终止当前测试
	t.Fail()               //表示测试失败,但不终止当前测试
	fmt.Println("after fail")
}
