#include "nwnx_time"
#include "nwnx_chat"
#include "o_inc"

void InitOrder() {
  NWNX_Chat_RegisterChatScript(o_chat);

  // random server values we want
  string sBootTime = NWNX_Time_GetSystemTime();
  string sBootDate = NWNX_Time_GetSystemDate();
  string sNwserver = GetModuleName();

  NWNX_Redis_HMSET(OrderObjectEdge(1),"BootTime",sBootTime);
  NWNX_Redis_HMSET(OrderObjectEdge(1),"BootDate",sBootDate);
  NWNX_Redis_HMSET(OrderObjectEdge(1),"ModuleName",sServerName);
  NWNX_Redis_HMSET(OrderObjectEdge(1),"Online","0");

  int zCursor = NWNX_Redis_HKEYS(OrderObjectEdge(1));
  int i; for (i = 0; i < NWNX_Redis_GetArrayLength(zCursor); i++) {
    int zEntry  = NWNX_Redis_GetArrayElement(zCursor, i);
    string sField = NWNX_Redis_GetResultAsString(zEntry);
    int zValue = NWNX_Redis_HMGET(OrderObjectEdge(1),sField);
    string sValue = NWNX_Redis_GetResultAsString(zEntry);
    WriteTimestampedLogEntry(sField +": "+sValue);
  }
}
