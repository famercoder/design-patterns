import models

//
//golang单例模式
//定义：单例对象的类必须保证只有一个实例存在，全局有唯一接口访问
//分类：
//[1]懒汉方式：指全局的单例实例在第一次被使用时构建
//[2]饿汉方式：指全局的单例实例在类装载时构建

//懒汉方式
type singleton struct{}
var ins *singleton
func GetIns() *singleton {
	if ins == nil {
		ins = &singleton{}
	}
	return ins
}
