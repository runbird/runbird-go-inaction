package search

// defaultMatcher实现默认匹配器
type defaultMatcher struct{}

//init函数将默认匹配器注册
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

//Search实现默认匹配器的行为
//方法声明为使用defaultMatcher类型的值作为接收者
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

//声明一个指向defaultMatcher类型值的指针
//dm := new(defaultMatcher)

//编译器会解开dm指针的引用，使用对应的值调用方法
//dm.Search(feed,"test")

//方法声明为使用指向defaultMatcher类型值的指针作为接收者
//func (m *defaultMatcher) Search(feed *Feed,searchTerm string)

//声明一个defaultMatcher
//var dm defaultMatcher

//编译器会自动生成指针引用dm值，使用指针调用方法
//dm.Search(feed,"test")
