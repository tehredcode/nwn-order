#include "nwnx_redis"

int OrderIsUUIDExists(string uuid);
int OrderIsUUIDExists(string uuid) {
  int nIsUnique = NWNX_Redis_HEXISTS(RdsEdgeServer("server")+":uuid", uuid);
  int nState = NWNX_Redis_GetResultAsInt(nIsUnique);
  return IntToString(nState);
}

void OrderAddUUIDtoRedis(string uuid) {
  NWNX_Redis_HSET(RdsEdgeServer("server")+":uuid", uuid, 1);
}

string OrderRandomLetterOrNumber();
string OrderRandomLetterOrNumber() {
  string sLetter;
  string sString = "abcdefghijklmnopqrstuvwxyz0123456789";
  int x = Random(34);

  string sLetter = GetSubString(sString, x, 1);
  return sLetter;
}

string OrderGenerateNewUUID() {
  string sUUID;
  int nUUIDgen;
  for (nUUIDgen = 0; nUUIDgen < 31; nUUIDgen++) {
    string sUUID =  sUUID + OrderRandomLetterOrNumber;
    nUUIDgen = nUUIDgen+1;
  }
  return IntToString(sUUID);
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