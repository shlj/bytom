package main

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -ldl -L.
/*
#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>

typedef int (*CAC_FUNC)(int, int);
#define LINUX_LIB_CSIMDTS_PATH "./cSimdTs.so"

int linux_csimdts() {
    printf("---------xxx---------\n");

    void *handle;
    char *error;
    CAC_FUNC cac_func = NULL;

    //打开动态链接库
    handle = dlopen(LINUX_LIB_CSIMDTS_PATH, RTLD_LAZY);
    if (!handle) {
    fprintf(stderr, "%s\n", dlerror());
    exit(EXIT_FAILURE);
    }

    //clear previous error
    dlerror();

    //获取一个函数
    *(void **) (&cac_func) = dlsym(handle, "add");
    if ((error = dlerror()) != NULL)  {
    fprintf(stderr, "%s\n", error);
    exit(EXIT_FAILURE);
    }
    printf("add: %d\n", (*cac_func)(2,7));

    cac_func = (CAC_FUNC)dlsym(handle, "sub");
    printf("sub: %d\n", cac_func(9,2));

    cac_func = (CAC_FUNC)dlsym(handle, "mul");
    printf("mul: %d\n", cac_func(3,2));

    cac_func = (CAC_FUNC)dlsym(handle, "div");
    printf("div: %d\n", cac_func(8,2));

    //关闭动态链接库
    dlclose(handle);
    exit(EXIT_SUCCESS);
}
*/
import "C"

import(
    // "fmt"
    // "unsafe"
    // "reflect"
    // "time"
)

func main() {
    C.linux_csimdts()
}