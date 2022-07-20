package znet

import "Zinx_Study/zinx/ziface"

// BaseRouter 实现Router时,  先嵌入这个BaseRouter基类, 然后根据需要对这个基类的方法进行重写就行
type BaseRouter struct {
}

// 这里之所有BaseRouter方法为空，是因为有的Router不希望有PreHandle和PostHandle这两个业务
// 所有Router继承BaseRouter的好处就是不再需要是需要实现PreHandle和PostHandle

// PreHandle 在处理coon业务之前的钩子方法Hook
func (br *BaseRouter) PreHandle(request ziface.IRequest) {

}

// Handle 在处理conn业务的主方法
func (br *BaseRouter) Handle(request ziface.IRequest) {

}

// PostHandle 在处理conn业务之后的钩子方法
func (br *BaseRouter) PostHandle(request ziface.IRequest) {

}
