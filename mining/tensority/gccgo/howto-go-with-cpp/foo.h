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
	uint8_t *get(uint8_t bh[32], uint8_t seed[32], uint8_t result[32]);

#ifdef __cplusplus
}
#endif

#endif
