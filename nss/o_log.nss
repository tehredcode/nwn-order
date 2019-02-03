#include "o_inc"
#include "o_external"
#include "o_webhook_out"

// channels we send log messages to
const string sLogDebug     = "Log:Debug";
const string sLogInfo      = "Log:Warning";
const string sDebugWarning = "Log:Fatal";
const string sDebugFatal   = "Log:Fatal";

///////////////////
//  Will check length of the list and determine how many old entries to purge
///////////////////
void CleanLog() {
  string sKeyPath = OrderObjectEdge(1) + "Logs";
  int nListLength = NWNX_Redis_LLEN(sKeyPath);
  NWNX_Redis_LTRIM(sKeyPath,2500,nListLength);
}

///////////////////
//  Log to a redis list
//  nForever defines if we keep this error forever
///////////////////
void RedisLog(string sMessage, int nLogLevel, int nForever=FALSE) {
  // get length of list
  int nListLength = NWNX_Redis_LLEN(OrderObjectEdge(1)+"Logs");
  // clean the temp logs if we need to
  if (nListLength > 3000) CleanLog();

  switch (nLogLevel) {
  case 0: NWNX_Redis_RPUSH(OrderObjectEdge(1)+"Logs",sMessage);         break;
  case 1: NWNX_Redis_RPUSH(OrderObjectEdge(1)+"CriticalLogs",sMessage); break;
  default:
  }
}

//
void OrderLogDebug(string sMessage){
  EOrderLogDebug(sMessage);
  NWNX_Redis_PUBLISH(sLogDebug, sMessage);
}

//
void OrderLogInfo(string sMessage){
  EOrderLogInfo(sMessage);
  NWNX_Redis_PUBLISH(sLogInfo, sMessage);
}

//
void OrderLogWarning(string sMessage){
  EOrderLogWarning(sMessage);
  NWNX_Redis_PUBLISH(sDebugWarning, sMessage);
}

// WARNING, THIS WILL SHUT DOWN ORDER
void OrderLogFatal(string sMessage){
  EOrderLogFatal(sMessage);
  NWNX_Redis_PUBLISH(sDebugFatal, sMessage);
}

///////////////////
// Log levels:
// 0:Debug
// 1:Info
// 2:Warning
// 3:Fatal
// Webhook Levels:
// 0:debug
// 1:private
// 2:public
///////////////////
void orderLog(string sMessage, int nLogLevel, int nWebhook=1, int nWebhookLevel=0, int nSaveError=1, int nPermSaveError=0) {

  if (nWebhook) OrderSendWebhook(nWebhookLevel, sMessage, "Order:Log: " + IntToString(nWebhookLevel));
  if (nSaveError) RedisLog(sMessage,nLogLevel,nPermSaveError);

  NWNX_Redis_PUBLISH("Log:"+IntToString(nLogLevel),sMessage);

  switch (nLogLevel) {
  case 0:  OrderLogDebug(sMessage);                     break;
  case 1:  OrderLogInfo(sMessage);                      break;
  case 2:  OrderLogWarning(sMessage);                   break;
  case 3:  OrderLogFatal(sMessage);                     break;
  default: OrderLogWarning("Log level not recognized"); break;
  }
}
