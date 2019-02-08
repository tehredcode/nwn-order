#include "nwnx_chat"
#include "o_inc"
#include "o_uuid"

string ChatChannel(int nChannel)
  switch(nChannel) {
    case 1:  return PLAYER_TALK;
    case 2:  return PLAYER_SHOUT;
    case 3:  return PLAYER_WHISPER;
    case 4:  return PLAYER_TELL;
    case 5:  return SERVER_MSG;
    case 6:  return PLAYER_PARTY;
    case 14: return PLAYER_DM;
    case 17: return DM_TALK;
    case 18: return DM_SHOUT;
    case 19: return DM_WHISPER;
    case 20: return DM_TELL;
    case 22: return DM_PARTY;
    case 30: return DM_DM;
  }
)

void main() {
  object oSender = NWNX_Chat_GetSender();
  if (GetTag(oSender) == "chat") {
    NWNX_Chat_SkipMessage();
  }

  int nChannel = NWNX_Chat_GetChannel();
  string sMessage = NWNX_Chat_GetMessage();
  object oTarget = NWNX_Chat_GetTarget();
  string sUUID = OrderGetNewUUID();
  string sChatEdge = OrderObjectEdge(int nType)+":chat:"+sUUID;
  NWNX_Redis_HMSET(sChatEdge, "message",sMessage);
  NWNX_Redis_HMSET(sChatEdge, "sender",GetName(oSender));
  NWNX_Redis_HMSET(sChatEdge, "target",GetName(oTarget));
  NWNX_Redis_HMSET(sChatEdge, "channel",ChatChannel(nChannel));
  NWNX_Redis_HMSET(sChatEdge, "npc",IntToString(GetIsPC(oSender));
  NWNX_Redis_PUBLISH("Discord:Out",sUUID);
}