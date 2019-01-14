#include "order_inc"
#include "order_external"
#include "order_alert"

///////////////////
//  Will check length of the list and determine how many old entries to purge
///////////////////
void CleanLog() {
  string sKeyPath = RdsEdgeServer("server") + "Logs"; 
  int nListLength = NWNX_Redis_LLEN(sKeyPath);
  NWNX_Redis_LTRIM(sKeyPath,2500,nListLength);
}

///////////////////
//  Log to a redis list 
//  nForever defines if we keep this error forever
///////////////////
void RedisLog(string sMessage, int nLogLevel, int nForever=FALSE) {
  // get length of list
  int nListLength = NWNX_Redis_LLEN(RdsEdgeServer("server")+"Logs");
  // clean the temp logs if we need to
  if (nListLength > 3000) CleanLog();

  switch (nLogLevel) {   
  case 0: NWNX_Redis_RPUSH(RdsEdgeServer("server")+"Logs",sMessage);         break;
  case 1: NWNX_Redis_RPUSH(RdsEdgeServer("server")+"CriticalLogs",sMessage); break;
  default:
  }
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

  if (nWebhook) OrderSendbWebhook(nWebhookLevel, sMessage, "Order:Log: " + IntToString(nWebhookLevel));
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