#include "nwnx_redis"
#include "o_external"

int GetTagIsUUID(object oObject) {
  return 1;
}

string OrderReturnModuleName(){
  int zReturn = NWNX_Redis_HGET(SERVER_NAME,"ModuleName");
  string sReturn = NWNX_Redis_GetResultAsString(zReturn);
  return sReturn;
}

// add uuid to list of used id's
void OrderAddUUIDtoRedis(string uuid) {
  NWNX_Redis_HSET(OrderReturnModuleName()+":server:uuid", uuid, "1");
}

// confirm uuid does not exist already.
int OrderIsUUIDExists(string uuid);
int OrderIsUUIDExists(string uuid) {
  int nIsUnique = NWNX_Redis_HEXISTS(OrderReturnModuleName()+":server:uuid", uuid);
  int nState = NWNX_Redis_GetResultAsInt(nIsUnique);
  return nState;
}

// generate random letter or Number defined by us.
string OrderRandomLetterOrNumber();
string OrderRandomLetterOrNumber() {
  string sString = "abcdefghijklmnopqrstuvwxyz0123456789";
  int x = Random(34);
  string sLetter = GetSubString(sString, x, 1);
  return sLetter;
}

// generate a uuid
string OrderGenerateNewUUID() {
  string sUUID;
  int nUUIDgen;
  do {
    string zUUID = sUUID;
    sUUID = OrderRandomLetterOrNumber() + zUUID;
    nUUIDgen++;
  } while (nUUIDgen < 31);
  return sUUID;
}

// Call uuid generate, confirm it's not a duplicate
string OrderGetNewUUID() {
  string sUUID = OrderGenerateNewUUID();
  int nUnique = OrderIsUUIDExists(sUUID);
  while (nUnique == 0){
    string sUUID = OrderGetNewUUID();
  }
  OrderAddUUIDtoRedis(sUUID);
  return sUUID;
}
