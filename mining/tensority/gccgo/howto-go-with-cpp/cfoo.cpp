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

void get(uint8_t *in){

	for(int i=0; i<32; ++i) {
		printf("0x%02x\n", (*in+i));
	}

	return;
}
