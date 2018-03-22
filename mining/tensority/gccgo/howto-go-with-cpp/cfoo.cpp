#include "foo.hpp"
#include "foo.h"
#include <iostream>
#include <cstdio>

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

uint8_t *get(uint8_t bh[32], uint8_t seed[32], uint8_t result[32]){
	printf("%s\n", "---------------");
	for(int i=0; i<32; ++i) {
		printf("0x%02x, ", bh[i]);
	}
	printf("\n%s\n", "---------------");
	for(int i=0; i<32; ++i) {
		printf("0x%02x, ", seed[i]);
	}
	printf("\n%s\n", "---------------");

	return result;
}
