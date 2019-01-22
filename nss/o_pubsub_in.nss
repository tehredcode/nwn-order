#include "nwnx_redis_ps"
#include "order_return"
#include "order_github"
#include "order_heartbeat"

int nMessageConvert(string sMessage) {
  if (sMessage == "webhook:github") return 0;
  if (sMessage == "heartbeat") return 1;
  if (sMessage == "discord:in") return 2;
  else return 99;
}

void main() {
  struct NWNX_Redis_PubSubMessageData data = NWNX_Redis_GetPubSubMessageData();
  int nMessage = nMessageConvert(data.channel);
  switch(nMessage) {
    case 0: 
      OrderGithub(data.message);
      break;
    case 1: 
      int nTicker = StringToInt(data.message);
      OrderHeartbeat(nTicker);
      break;
    case 2: 
      OrderIncomingDiscord(data.message);
      break;
    default:
      break;
  }
}