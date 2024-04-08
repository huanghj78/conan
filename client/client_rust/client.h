#ifndef __CLIENT_H__
#define __CLIENT_H__

#include <stdint.h>

typedef unsigned long long uint64;
typedef unsigned int uint32;

typedef struct {
    uint64 term;
} CPFIMessage;

int BeforeSendReq(CPFIMessage *msg);
int AfterSendReq(CPFIMessage *msg);
int BeforeSendAck(CPFIMessage *msg);
int AfterSendAck(CPFIMessage *msg);
int BeforeRecvAck(CPFIMessage *msg);
int BeforeRecvReq(CPFIMessage *msg);

#endif