问题:import 导包时包名前加"_"?
解释:让go语言对包做初始化操作(调用相应的init函数),但是并不使用包里的标识符

问题:init函数神魔时候执行?
解释:程序中每个代码文件的init函数都会在main函数执行前调用

问题:变量没有定义在任何函数作用域内,会被当成什么变量?
解释:会被当成包级别的变量
    go中大写开头标识符是公开的,小写不公开
    
问题:go语言零值?
解释:map变量默认零值是nil
    在go中,所有变量都被初始化为其零值。
    数值类型,是0
    字符串类型,是空字符串
    bool类型,false
    指针类型,nil
    引用类型,所引用的底层数据结构会被初始化为其对应的零值

问题:log.Fatal(err)?
解释:Fatal接受一个错误值,并将这个错误在终端窗口里输出,随后终止程序

问题:通道?
解释:通道本身实现的是一组带类型的值,这组值用于在go routine之间传递数据

问题； for range
解释: for range 迭代时返回两个值,第一个值是元素在切片的索引位,第二个是元素值的一个副本
      "_"代表占位符
问题:接口命名事项?
解释:如果接口类型只包含一个方法,那么这个类型的名字以er结尾,如果接口内部
    生命了多个方法,其名字需要与其行为关联
问题:实现接口?
解释:如果要让一个用户定义的类型实现一个接口,这个用户定义的类型要实现接口类型
     里声明的所有方法

问题:空结构？
解释:空结构在创建实例时，不会分配任何内存。


rss文档格式:
<rss xmlns:npr="http://www.npr.org/rss/" xmlns:nprml="http://api" <channel>
<title>News</title> <link>...</link> <description>...</description>
log.Fatalln(feedType, "Matcher already registered")
<language>en</language>
<copyright>Copyright 2014 NPR - For Personal Use <image>...</image>
<item>
<title>
Putin Says He'll Respect Ukraine Vote But U.S.
                </title>
                <description>
The White House and State Department have called on the </description>

问题:函数参数是数组该怎么传?
解释:在函数间传递数组开销很大,很占内存,因为函数之间传递不安良是以值的方式传递
    ,所以在需要数组时，应该传入指向数组的指针,这样只需要对应类型的字节数即可

问题:切片?
解释:切片是围绕动态数组概念构建的,可以按需自动增长和缩小.
    切片动态增张是通过append来实现的.
