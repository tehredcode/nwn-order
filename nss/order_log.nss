#include "order_inc"

void orderLog(string sMessage, string sLogLevel) {
  // call the external scripts first, seems best for right now
  if (sLogLevel == "Debug")   { OrderLogDebug(sMessage); }
  if (sLogLevel == "Info")    { OrderLogInfo(sMessage); }
  if (sLogLevel == "Warning") { OrderLogWarning(sMessage); }
  if (sLogLevel == "Fatal")   { OrderLogFatal(sMessage); }

  //////////////////////
  //  Pubsub channels we sub to in ORDER
  //  Log:Debug 
  //  Log:Info
  //  Log:Warning
  //  Log:Fatal
  //////////
  WNX_Redis_PUBLISH("Log:"+sLogLevel,sMessage);
}