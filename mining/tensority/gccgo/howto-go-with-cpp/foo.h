#ifndef _MY_PACKAGE_FOO_H_
#define _MY_PACKAGE_FOO_H_

#ifdef __cplusplus
extern "C" {
#endif
	#include <stdint.h>

	typedef void* Foo;
	Foo FooInit(void);
	void FooFree(Foo);
	void FooBar(Foo);
	void get(uint8_t*);

#ifdef __cplusplus
}
#endif

#endif
