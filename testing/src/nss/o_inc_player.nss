#include "nwnx_redis"
#include "o_inc"

// Add a player system value
void OrderPlayerSystemSetValue(object oPC, string sSystem, string sKey, string sValue) {
  NWNX_Redis_HMSET(OrderUniqueObjectEdge(oPC)+":"+sSystem,sKey,sValue);
}

// Remove a player system value value
void OrderPlayerSystemRemoveValue(object oPC, string sSystem, string sKey) {
  NWNX_Redis_HDEL(OrderUniqueObjectEdge(oPC)+":"+sSystem,sKey);
}

// Get a player system value as string
string OrderPlayerSystemGetValueString(object oPC, string sSystem, string sKey) {
  int zReturn = NWNX_Redis_HGET(OrderUniqueObjectEdge(oPC)+":"+sSystem,sKey);
  string sReturn = NWNX_Redis_GetResultAsString(zReturn);
  return sReturn;
}

// Get a player system value as int
int OrderPlayerSystemGetValueInt(object oPC, string sSystem, string sKey) {
  int zReturn = NWNX_Redis_HGET(OrderUniqueObjectEdge(oPC)+":"+sSystem,sKey);
  int sReturn = NWNX_Redis_GetResultAsInt(zReturn);
  return sReturn;
}

// Add a value to the core player hash
void OrderPlayerSetValue(object oPC, string sKey, string sValue) {
  NWNX_Redis_HMSET(OrderUniqueObjectEdge(oPC),sKey,sValue);
}

// Remove a value from the core player hash
void OrderPlayerRemoveValue(object oPC, string sKey) {
  NWNX_Redis_HDEL(OrderUniqueObjectEdge(oPC),sKey);
}

// Get a string from the core player hash
string OrderPlayerGetValueString(object oPC, string sKey) {
  int zReturn = NWNX_Redis_HGET(OrderUniqueObjectEdge(oPC),sKey);
  string sReturn = NWNX_Redis_GetResultAsString(zReturn);
  return sReturn;
}

// Get a int from the core player hash
int OrderPlayerGetValueInt(object oPC, string sKey) {
  int zReturn = NWNX_Redis_HGET(OrderUniqueObjectEdge(oPC),sKey);
  int nReturn = NWNX_Redis_GetResultAsInt(zReturn);
  return nReturn;
}

// Remove a player, Yes this deletes everything related to the character inside redis.
void OrderPlayerDeleteCharacter(object oPC) {
  NWNX_Redis_DEL(OrderUniqueObjectEdge(oPC));
}
