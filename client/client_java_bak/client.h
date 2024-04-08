#ifndef __CLIENT_H__
#define __CLIENT_H__

#include <stdint.h>

typedef unsigned long long uint64;
typedef unsigned int uint32;

typedef struct {
    uint64 term;
} CPFI_Message;

int Hello();

#endif