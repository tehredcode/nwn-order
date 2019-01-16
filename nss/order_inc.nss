#include "nwnx_redis"

string OrderObjectUUID(int nEdgeType);
string OrderObjectUUID(int nEdgeType) {
  string sNwserver = GetModuleName();
  switch (nEdgeType) {
    case 1: { return sNwserver+":server"; }        
    case 2: { return sNwserver+":item"; }
    case 3: { return sNwserver+":creature"; }
    case 4: { return sNwserver+":door"; }
    case 5: { return sNwserver+":placeable"; } 
    case 6: { return sNwserver+":quests"; }     
    default:{ return "err"; }
  }
}

int OrderIsUUIDExists(string uuid);
int OrderIsUUIDExists(string uuid) {
  int nIsUnique = NWNX_Redis_HEXISTS(RdsEdgeServer(1)+":uuid", uuid);
  int nState = NWNX_Redis_GetResultAsInt(nIsUnique);
  return nState;
}

void OrderAddUUIDtoRedis(string uuid) {
  NWNX_Redis_HSET(RdsEdgeServer("server")+":uuid", uuid, "1");
}

string OrderRandomLetterOrNumber();
string OrderRandomLetterOrNumber() {
  string sString = "abcdefghijklmnopqrstuvwxyz0123456789";
  int x = Random(34);
  string sLetter = GetSubString(sString, x, 1);
  return sLetter;
}

string OrderGenerateNewUUID() {
  string sUUID;
  int nUUIDgen;
  do {
    string zUUID = sUUID;
    string sUUID = OrderRandomLetterOrNumber() + zUUID;
    nUUIDgen++;
  } while (nUUIDgen < 31);
  return sUUID;
}

string OrderGetNewUUID() {
  string sUUID = OrderGenerateNewUUID();
  int nUnique = OrderIsUUIDExists(sUUID);
  while (nUnique == 1){
    string sUUID = OrderGetNewUUID();
  }
  OrderAddUUIDtoRedis(sUUID);
  return sUUID;
}

// -- return or assign and return the oPC uuid.
string OrderGetUUIDPlayer(object oPC);
string OrderGetUUIDPlayer(object oPC) {
  // if the user has no uuid set
  if (GetTag(oPC) == "") {  
    object oMod = GetModule();
    string sUUID = OrderGetNewUUID();

    //set oPC tag to uuid
    SetTag(oPC, sUUID);
    return sUUID;
  } else {
    string sUUID = GetTag(oPC);
    return sUUID;
  }
}

string OrderPlayerUUID(object oPC);
string OrderPlayerUUID(object oPC) {
  string Nwserver = GetModuleName();
  string UUID     = OrderGetUUIDPlayer(oPC);
  return sNwserver+":player:"+UUID;
}