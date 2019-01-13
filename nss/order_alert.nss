#include "nwnx_webhook"

// Send webhook
void OrderSendbWebhook(int nPermission, string sMessage, string sSendername) {
  object oMod = GetModule();
  string sWebhookUrlPublic  = GetLocalString(oMod, "WEBHOOK_PUBLIC");
  string sWebhookUrlPrivate = GetLocalString(oMod, "WEBHOOK_PRIVATE");
  string sWebhookUrlDebug   = GetLocalString(oMod, "WEBHOOK_DEBUG");

  switch (nPermission) {   
  case 0:
    NWNX_WebHook_SendWebHookHTTPS("discordapp.com/api/webhooks/" + sWebhookUrlDebug + "/slack", sMessage, sSendername);
  case 1:
    NWNX_WebHook_SendWebHookHTTPS("discordapp.com/api/webhooks/" + sWebhookUrlPrivate + "/slack", sMessage, sSendername);
  case 2:
    NWNX_WebHook_SendWebHookHTTPS("discordapp.com/api/webhooks/" + sWebhookUrlPublic + "/slack", sMessage, sSendername);
  default:
  }
}