#include "nwnx_redis_ps"
#include "order_return"
#include "order_github"
#include "order_heartbeat"

int nMessageConvert(string sMessage) {
  if (sMessage == "github") return 0;
  if (sMessage == "heartbeat") return 1;
  if (sMessage == "discord:in") return 2;
  else return 99;
}

void main() {
  struct NWNX_Redis_PubSubMessageData data = NWNX_Redis_GetPubSubMessageData();
  int nMessage = nMessageConvert(data.channel);
  switch(nMessage) {
    case 0: 
      OrderReturn(data.message);
    case 1: 
      OrderGithub(data.message);
    case 2: 
      OrderHeartbeat(data.message);
    default:
  }
}