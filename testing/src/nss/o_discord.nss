//#include "nwnx_redis"
//#include "nwnx_chat"
//#include "o_inc"
//
//void OrderIncomingDiscord(string sID);
//void OrderIncomingDiscord(string sID){
//    object oArea = GetFirstArea();
//    while(GetIsObjectValid(oArea))
//    {
//        if (GetTag(oArea) == "_discord_") {
//
//          object oObject = GetFirstObjectInArea(oArea); // _discord_ is the area we are checking in
//          while(GetIsObjectValid(oObject)){
//          if(GetTag(oObject) == "chat"){ // creature tag here
//            int zSender = NWNX_Redis_HGET(OrderObjectEdge(4)+sID+":Sender",sKey);
//            string sSender = NWNX_Redis_GetResultAsString(zSender);
//            int zDiscordMessage = NWNX_Redis_HGET(OrderObjectEdge(4)+sID+":Message",sKey);
//            string sDiscordMessage = NWNX_Redis_GetResultAsString(zDiscordMessage);
//            // rename npc to match discord message sender
//            SetName(oObject,sSender);
//            // make npc send message
//            NWNX_Chat_SendMessage(2, sDiscordMessage, oObject);
//            break;
//          }
//          object oObject = GetNextObjectInArea(oArea);
//        }
//      }
//    oArea = GetNextArea();
//  }
//}
