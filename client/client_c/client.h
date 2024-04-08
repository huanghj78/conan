#ifndef __CLIENT_H__
#define __CLIENT_H__

#include <stdint.h>

typedef unsigned long long uint64;
typedef unsigned int uint32;

typedef struct {
    uint64 term;
} CPFI_Message;

int BeforeSendReq(CPFI_Message *msg);
int AfterSendReq(CPFI_Message *msg);
int BeforeSendAck(CPFI_Message *msg);
int AfterSendAck(CPFI_Message *msg);
int BeforeRecvAck(CPFI_Message *msg);
int BeforeRecvReq(CPFI_Message *msg);

#endif