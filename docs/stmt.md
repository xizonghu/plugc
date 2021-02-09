使用token类型描述语句格式

### token类型：
key：关键字
op：操作符
sym：符号
div：分解符
const：常量

# 语句：

### 声明：变量
> 格式： key{1,} (op{0,} sym div){1,} div

### 声明：结构体
> 格式： key sym div (变量){1,} div div

### 声明：数组
> 格式： (变量 | 结构体) div const div

### 声明：自定义类型
> 格式： key key sym div (变量){1,} div div

### 声明：函数

### 定义：标签

### 定义：变量
> 格式： key{1,} (op{0,} sym (op [const | sym]){0,1} op{0,}) div

### 定义：结构体
> 格式： key sym op{0,} sym (op div ([const | sym] op){0,} div){0,1} div

### 定义： 数组

### 定义：自定义类型
> 格式： sym op{0,} sym (op div ([const | sym] op){0,} div){0,1} div

### 定义：函数
> 格式： key{1,} sym div (key{1,} sym){0,} div

### 表达式
> 格式： sym op{1,} [const | sym]{1,} div

### 函数调用
> (变量 =){0,1} 函数名 表达式 ;

### 控制：条件
> if 表达式 (else if){0,} else{0,}
> switch 表达式 (case 常量: 段 break{0,1}){0,} (default: 段 break{0,1}){0,1}

### 控制：循环
> for 表达式{0,1} ; 表达式{0,1} ; 表达式{0,1} 段
> do 段 while 表达式
> while 表达式 段

### 控制：跳转
> goto 标签
> return 表达式

### 段
> (语句 ;){0,}

### 空
> ;

