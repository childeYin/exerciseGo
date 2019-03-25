# cha1

1. func 首字母大写 外部可引用，否则外部不可用，只能内部调用
2. 声明变量，赋值，如果只声明了，未赋值则会自动采用零值机制。记int 0, string "", bool false
3. 指针 局部变量的地址是安全的
4. go GC mark && sweep, 
    - mark： mark the object's status
    - sweep: if object's status is unreached, they will be recycled
    - GCPercent [CPU & Memory]
    - MaxHeap []
5. go %取模，符号综合被除数保持一致
6. 字符串作为只读的类型，我们并不会直接向字符串直接追加元素改变其本身的内存空间，
   所有追加的操作都是通过拷贝来完成的。新的字符串其实就是一片新的内存空间，
   与原来的字符串没有任何关联。指针地址并不会改
    
