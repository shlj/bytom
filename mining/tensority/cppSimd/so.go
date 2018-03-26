package main

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -ldl -L.
/*
#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>

typedef int (*CSIMDTS_FUNC)(int, int);
#define LINUX_LIB_CSIMDTS_PATH "./cSimdTs.so"

int linux_csimdts() {
    printf("---------xxx---------\n");

    void *handle;
    char *error;
    CSIMDTS_FUNC csimdts_func = NULL;

    //open the dynamic lib
    handle = dlopen(LINUX_LIB_CSIMDTS_PATH, RTLD_LAZY);
    if (!handle) {
        fprintf(stderr, "%s\n", dlerror());
        exit(EXIT_FAILURE);
    }

    // clear previous error
    dlerror();

    // get the func
    *(void **) (&csimdts_func) = dlsym(handle, "add");
    if ((error = dlerror()) != NULL)  {
        fprintf(stderr, "%s\n", error);
        exit(EXIT_FAILURE);
    }
    printf("add: %d\n", (*csimdts_func)(2,7));

    csimdts_func = (CSIMDTS_FUNC)dlsym(handle, "sub");
    printf("sub: %d\n", csimdts_func(9,2));

    // close dynamic lib
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