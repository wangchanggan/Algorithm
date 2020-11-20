package sort

type SortInterface interface {
	// 数据的长度
	Len() int
	// 元素i是否在元素j之前，用于升序或降序，例如p[i]<p[j]为降序（把元素j放到元素i前面）
	Less(i, j int) bool
	// 元素i和j交换
	Swap(i, j int)
	//获取元素i
	Get(i int) (data interface{})
	//覆盖元素i
	Coverage(i int, data interface{})
	//元素比较
	LessData(iData, jData interface{}) bool
}