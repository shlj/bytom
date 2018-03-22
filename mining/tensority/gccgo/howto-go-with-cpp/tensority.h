#ifndef _TENSORITY_H_
#define _TENSORITY_H_

#ifdef __cplusplus
extern "C" {
#endif
	#include <stdint.h>

	uint8_t *get(uint8_t blockheader[32], uint8_t seed[32]);

#ifdef __cplusplus
}
#endif

#endif
