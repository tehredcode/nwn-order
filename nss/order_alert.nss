#include "nwnx_webhook"
#include "order_log"

// Send webhook
void OrderSendbWebhook(int nPermission, string sMessage, string sSendername) {
  object oMod = GetModule();
  string sWebhookUrlPublic  = GetLocalString(oMod, "WEBHOOK_PUBLIC");
  string sWebhookUrlPrivate = GetLocalString(oMod, "WEBHOOK_PRIVATE");
  string sWebhookUrlDebug   = GetLocalString(oMod, "WEBHOOK_DEBUG");

  switch (nPermission) {   
  case 0:
    NWNX_WebHook_SendWebHookHTTPS("discordapp.com/api/webhooks/" + sWebhookUrlDebug + "/slack", sMessage, sSendername);
    break;
  case 1:
    NWNX_WebHook_SendWebHookHTTPS("discordapp.com/api/webhooks/" + sWebhookUrlPrivate + "/slack", sMessage, sSendername);
    break;
  case 2:
    NWNX_WebHook_SendWebHookHTTPS("discordapp.com/api/webhooks/" + sWebhookUrlPublic + "/slack", sMessage, sSendername);
    break;
  default:
    orderLog(sMessage, 2);
    break;
  }
}