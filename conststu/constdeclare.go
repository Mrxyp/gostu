package conststu
//常量的声明

//方式1
const NAME string = "xyp"
//方式2
const AGE = 2
//方式3
const IPHONE,HUAIWEI = 1,"huaiwei"
//方式4
const(
	cat="猫"
	dog string ="狗"
)
//方式5
const alen2  = len(NAME) //只能用库函数  自定义的函数返回值不能给常量