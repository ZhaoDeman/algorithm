go string类型
> 设计：string 类型是不可修改的，无论是java还是go,或者其他语言，字符串都是不可修改的。
> 由于字符串是不可变的，所以可以安全地共享底层数据，而不必担心副本的创建和修改。