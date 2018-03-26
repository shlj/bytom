#ifndef _C_LOADLIB_H_
#define _C_LOADLIB_H_

#ifdef __cplusplus
extern "C" {
#endif
    #define LIB_CACULATE_PATH "./libcaculate.so" //动态链接库路径
    typedef int (*CAC_FUNC)(int, int); //函数指针

    int dododo();
    int fuck();

#ifdef __cplusplus
}
#endif

#endif
