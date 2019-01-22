#include "nwnx_webhook"

// Send webhook
void OrderSendWebhook(int nPermission, string sMessage, string sSendername) {
  object oMod = GetModule();
  switch (nPermission) {   
  case 0: {
    string sWebhookUrlDebug   = GetLocalString(oMod, "WEBHOOK_DEBUG");
    NWNX_WebHook_SendWebHookHTTPS("discordapp.com/api/webhooks/" + sWebhookUrlDebug + "/slack", sMessage, sSendername);
    break;
  }
  case 1: {
    string sWebhookUrlPrivate = GetLocalString(oMod, "WEBHOOK_PRIVATE");
    NWNX_WebHook_SendWebHookHTTPS("discordapp.com/api/webhooks/" + sWebhookUrlPrivate + "/slack", sMessage, sSendername);
    break;
  }
  case 2: {
    string sWebhookUrlPublic  = GetLocalString(oMod, "WEBHOOK_PUBLIC");
    NWNX_WebHook_SendWebHookHTTPS("discordapp.com/api/webhooks/" + sWebhookUrlPublic + "/slack", sMessage, sSendername);
    break;
    }
  default: {
    break;
    }
  }
}