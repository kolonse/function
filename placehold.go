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

// 创建站位符 必须返回是一个
func PH(p uint) *Placeholder {
	var ph Placeholder
	ph.Set(p)
	return &ph
}

var globalPH *Placeholder
