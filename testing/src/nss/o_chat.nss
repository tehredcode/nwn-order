#include "nwnx_chat"
#include "nwnx_redis"

int GetIsPlayer(object oTarget) {
    if (GetIsPC(oTarget) == TRUE) { return TRUE; 
    } else if (GetIsDM(oTarget) == TRUE)  { return TRUE; 
    } else return FALSE;
}

string OrderReturnChatChannel(int nChannel) {
  switch(nChannel) {
    case 1:  return "PC-Talk"; 
    case 2:  return "PC-Shout";     
    case 3:  return "PC-Whisper"; 
    case 4:  return "PC-Tell"; 
    case 5:  return "Server-Message"; 
    case 6:  return "PC-Party"; 
    case 14: return "PC-DM";
    case 17: return "DM-Talk";
    case 18: return "DM-Shout";
    case 19: return "DM-Whisper";
    case 20: return "DM-Tell";
    case 22: return "DM-Party";
    case 30: return "DM-DM";
  }
  return "Err";
}

void main() {   
    string sChat;

    // channel
    int nChannel = NWNX_Chat_GetChannel();
    string sChannel = OrderReturnChatChannel(nChannel);

    // break if it's the server talking
    if (nChannel == 5) return;

    // confirm the object calling the chat is a player
    if (GetIsPlayer(OBJECT_SELF) == FALSE) return;

    // sender
    object oNameSender = NWNX_Chat_GetSender();
    string sNameSender = GetName(oNameSender);
    if (NWNX_Chat_GetSender() == OBJECT_INVALID) return;

    // receiver
    object oNameReceiver = NWNX_Chat_GetTarget();
    string sNameReceiver = GetName(oNameReceiver);

    // get the message sent
    string sMessage = NWNX_Chat_GetMessage();

    // make a string order can parse and log
    sChat = sNameSender+":"+sNameReceiver+":"+sChannel+":"+sMessage;
    NWNX_Redis_PUBLISH("discord:out",sChat);
}