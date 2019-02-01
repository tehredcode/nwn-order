void InitOrder() {
  // chat
  NWNX_Chat_RegisterChatScript(o_chat);

  // random server values we want
  string sBootTime = NWNX_Time_GetSystemTime();
  string sBootDate = NWNX_Time_GetSystemDate();
  string sNwserver = GetModuleName();

  NWNX_Redis_HMSET(OrderObjectEdge(1),"BootTime",sBootTime);
  NWNX_Redis_HMSET(OrderObjectEdge(1),"BootDate",sBootDate);   
  NWNX_Redis_HMSET(OrderObjectEdge(1),"ModuleName",sNwserver);
  NWNX_Redis_HMSET(OrderObjectEdge(1),"Online","0");
    
  int zCursor = NWNX_Redis_HKEYS(sServerStatHash);
  int i; for (i = 0; i < NWNX_Redis_GetArrayLength(zCursor); i++) { 
    int zEntry  = NWNX_Redis_GetArrayElement(zCursor, i); 
    string sField = NWNX_Redis_GetResultAsString(zEntry);
    int zValue = NWNX_Redis_HMGET(sServerStatHash,sField);
    string sValue = NWNX_Redis_GetResultAsString(zEntry);
    WriteTimestampedLogEntry(sField +": "+sValue);
  }
}