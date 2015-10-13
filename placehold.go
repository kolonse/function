// placehold
package function

type Placeholder uint /*{
	index uint
}*/

func (ph *Placeholder) Set(p uint) {
	//ph.index = p
	*ph = Placeholder(p)
}

func (ph *Placeholder) Get() uint {
	return uint(*ph) //.index
}

// 站位符 标志
/*
const (
	P_1 = iota
	P_2
	P_3
	P_4
	P_5
	P_6
	P_7
	P_8
	P_9
)
*/
// 创建站位符 必须返回是一个
func PH(p uint) *Placeholder {
	var ph Placeholder
	ph.Set(p)
	return &ph
}

var globalPH *Placeholder
var P_1 *Placeholder
var P_2 *Placeholder
var P_3 *Placeholder
var P_4 *Placeholder
var P_5 *Placeholder
var P_6 *Placeholder
var P_7 *Placeholder

func init() {
	P_1 = PH(0)
	P_2 = PH(1)
	P_3 = PH(2)
	P_4 = PH(3)
	P_5 = PH(4)
	P_6 = PH(5)
	P_7 = PH(6)
}
