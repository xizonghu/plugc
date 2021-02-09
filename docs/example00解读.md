源代码：
void sayHi(int x, char y) {
    int a;
    a = a + 21;
    printf("hello, world: %d, %d, %d\n", a, x, y);
}

void main() {
	int x;
	int y;
	x=17;
	y=19;
    sayHi(x, y);
}

字节码：
4C4156120000000000000000000000003C00203B5000003E0F0002030D0082000F0D0045150035380D68656C6C6F2C20776F726C643A2025642C2025642C2025640A000F0D000F05000E09000104823F3E0900000305008200011135380307008200011335380F05000F07003D17000040

名词：
数据类型：Char(1 byte), Int(2 byte), Addr(3 byte), Long(4 byte)
栈元素(4 byte)

RAM: 静态存储空间
RTM: 运行时存储空间，寻址空间为[0x0000, 0xFFFF]，RTM是RAM空间中的某段，是动态的，不真实的只有空间，也不对空间进行读写操作，读写全由RAM来完成。RTM只记录start和end
APP: 代码存储空间，寻址空间为[0x000000, 0xFFFFFF], 最大1MB
STA: 栈存储空间
STR: 字符存储空间
这些存储空间只有STA是按4字节来读写的，其他都按1字节来读写


解读：
addr 0x000000:
    4C415612000000000000000000000000        文件头信息
    3C0020                                  初始化RTM: RTM.start=0x2000; RTM.end=0x2000
    3B500000                                goto 0x000050

addr 0x000017:
void sayHi(int x, char y) {
    3E0F0002                                函数入口，运行内存大小为0x000F，函数参数个数为0x02：
                                            //保存父函数现场RTM信息，载入子函数现场RTM信息
                                            RAM.addr(RTM.end+3,2,RTM.start);  //在RAM中操作：定位父RTM.end，隔离一个地址(3字节)，保存父RTM.start
                                            RTM.start=RTM.end;  //设置子RTM.start
                                            RTM.end=RTM.start+APP.Int;  //设置子RTM.end
                                            //在RTM限定的RAM中，初始化参数值
                                            params=APP.Char;
                                            while(--params>=0) RAM.addr(RTM.start+5+4*params,4,pop());  //把STA中的参数写入RAM，释放STA

int a;                                      //声明并不产生实际代码，只会影响函数RTM的大小
a=a+21;
    030D008200                              push(0x0082000D);  //变量a在RTM中的位置入STA
    0F0D00                                  val=RAM(0x000D+RTM.start).Int; push(val);  //将变量a从RAM中读到STA中，准备做运算
    451500                                  加法运算： arg1=pop(); arg2=(arg1+0x0015);push(b)
    35                                      为一块内存初始化值:
                                            data=pop();
                                            offset=pop();  //对应push(0x0082000D)
                                            addr=(offset>>0)&0xFFFF; len=(offset>>16)&0xFFFF; memset(addr,data,len);  //变量a赋值
    38                                      pop()

printf("hello, world: %d, %d, %d\n", a, x, y);
    0D68656C6C6F2C20776F726C643A2025642C2025642C2025640A00      将字符串"hello, world: %d, %d, %d\n"载入到字符内存中
                                            STR(APP.seek);
                                            push(STR|0x00100000);  //从APP中载入字符串到STA，STA地址入STA【将高位做地址类型标记？】
    0F0D00                                  var=RAM(RTM.start+0x000D).Int; push(var);  //从RAM中读出变量a，入STA
    0F0500                                  var=RAM(RTM.start+0x0005).Int; push(var);  //从RAM中读出变量x，入STA
    0E0900                                  var=RAM(RTM.start+0x0009).Char; push(var);  //从RAM中读出变量y，入STA
    0104                                    push(0x04);  //函数参数个数，入STA
    82                                      系统调用printf()：
                                            params=pop(Char); seek(-params); addr=peek(i++); 获取字符内存地址addr中的字符ch，if(ch==%d) {num=peek(i++)}，调用系统打印函数打印num
}
    3F                                      调用返回：
                                            //还原父函数现场
                                            addr=RAM(RTM.start).Addr;  //invoke点
                                            RTM.end=RTM.start;  //设置父函数RTM.end
                                            RTM.start=RAM((RTM.end+3)&0xffff).Int;  //设置父函数RTM.start
                                            APP.seek(addr);  //pc指向invoke点

addr 0x000050:
void main() {
    3E090000                                函数入口，运行内存大小为0x0009，函数参数个数为0x00
int x;
x=17;
    0305008200                              push(0x00820005)
    0111                                    push(0x11)
    35                                      data=pop(); offset=pop(); addr=(offset>>0)&0xFFFF; len=(offset>>16)&0xFFFF; memset(addr,data,len);
    38                                      pop()

int y;
y=19;
    0307008200                              push(0x00820007)
    0113                                    push(0x13)
    35                                      data=pop(); offset=pop(); addr=(offset>>0)&0xFFFF; len=(offset>>16)&0xFFFF; memset(addr,data,len);
    38                                      pop()

sayHi(x, y);
    0F0500                                  var=RAM(RTM.start+0x0005).Int; push(var);  //从RAM中读出变量x，入STA
    0F0700                                  var=RAM(RTM.start+0x0007).Int; push(var);  //从RAM中读出变量y，入STA
    3D170000                                函数调用invoke：
                                            //保存现场
                                            RAM.Addr(RTM.end,APP.seek);  //将invoke点写入RAM
                                            goto 0x000017;
}
    40                                      结束



### RAM内容：
0x00(n bytes): 全局变量
0xXX(n bytes): RTM
0xXX(n bytes): RTM
...
0xXX(n bytes): RTM

函数sayHi(x,y)的RTM所指存储空间内容：
0x00(3 bytes): invoke点地址
0x03(2 bytes): 父RTM.start
0x05(4 bytes): 函数参数x
0x09(4 bytes): 函数参数y
0x0D(2 bytes): 函数内部变量a

### APP内容：
0x000000(16 bytes): 头信息
0x000017(n bytes): 初始化全局变量(0x41指令)

### 程序片段

#### 加法语句：
z=x+y;                                  //x的地址是0x0005，y的地址是0x0007，z的地址是0x0009
    0309008200                          push(0x00820009);  //z地址入栈【关于0x008XXXXX，这里的8是标记局部变量的，】
    0F0500                              var=RAM(RTM.start+0x0005).Int; push(var);  //从RAM中读出变量x，入STA
    0F0700                              var=RAM(RTM.start+0x0007).Int; push(var);  //从RAM中读出变量y，入STA
    21                                  push(pop()+pop())
    35                                  data=pop();  //变量的值
                                        offset=pop();  //对应0x00820009
                                        addr=(offset>>0)&0xFFFF;  //变量的RAM地址
                                        if(addr&0x00800000) addr+=RTM.start;  //判断为局部变量
                                        len=(offset>>16)&0x7F;  //变量的长度
                                        memset(addr,data,len);  //赋值
    38                                  pop();  //?



#### if语句
addr 0x00005B:
if (x > 33) {
    0F 0500         push(RAM(RTM.start+0x0005).Int);  //从RAM中读出变量x，入STA
    4E 2100         push(0x0021<pop()?1:0);  //
    38              pop();
    39 730000       if(peek()==0){APP.seek(0x000073)}  //有条件跳转
x=17;
    03 05008200
    01 11
    35
    38
    3B 7C0000       goto 0x00007C;  //无条件跳转
}

addr 0x000073:
else {
y=19;
    03 07008200
    01 13
    35
    38
}

addr 0x00007C:
    ...
