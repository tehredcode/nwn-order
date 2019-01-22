void OrderIncomingDiscord(string sID); 
void OrderIncomingDiscord(string sID){
  object oObject = GetFirstObjectInArea("_discord_");
  while(GetIsObjectValid(oObject))
  {
    if(GetTag(oObject) == "chat")
    {
      int zSender = NWNX_Redis_HGET(OrderObjectEdge(4)+sID+":Sender",sKey);
      string sSender = NWNX_Redis_GetResultAsString(zSender);
      int zDiscordMessage = NWNX_Redis_HGET(OrderObjectEdge(4)+sID+":Message",sKey);
      string sDiscordMessage = NWNX_Redis_GetResultAsString(zDiscordMessage);
      
      // rename npc to match discord message sender
      SetName(oObject,sSender);      
      // make npc send message
      NWNX_Chat_SendMessage(2, sDiscordMessage, oObject);
      // reset name back
      SetName(oObject,"chat"); 
      break;
    }
    object oObject = GetNextObjectInArea(oArea);
  } 
}
