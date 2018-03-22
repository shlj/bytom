#include <iostream>
#include <cstdio>
#include "foo.hpp"
#include "foo.h"
#include "BytomPoW.h"
#include "seed.h"

Foo FooInit() {
	cxxFoo * ret = new cxxFoo(666);
	return (void*)ret;
}

void FooFree(Foo f) {
	cxxFoo * foo = (cxxFoo*)f;
	delete foo;
}

void FooBar(Foo f) {
	cxxFoo * foo = (cxxFoo*)f;
	foo->Bar();
}

BytomMatList16* matList_int16;
uint8_t result[32] = {0};

uint8_t *get(uint8_t blockheader[32], uint8_t seed[32]){
    uint32_t exted[32];
    extend(exted, seed); // extends seed to exted
    Words32 extSeed;
    init_seed(extSeed, exted);

    matList_int16 = new BytomMatList16;
    matList_int16->init(extSeed);

    iter_mineBytom(blockheader, 32, result);

	return result;
}
