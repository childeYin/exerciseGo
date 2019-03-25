# cha1

1. func 首字母大写 外部可引用，否则外部不可用，只能内部调用
2. 声明变量，赋值，如果只声明了，未赋值则会自动采用零值机制。记int 0, string "", bool false
3. 指针 局部变量的地址是安全的
4. go GC mark && sweep, 
    - mark： mark the object's status
    - sweep: if object's status is unreached, they will be recycled
    - GCPercent [CPU & Memory]
    - MaxHeap []
    
