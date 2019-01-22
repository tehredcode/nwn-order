#include "nwnx_redis"
#include "order_inc"

// item system functions
void OrderPlayerSystemSetValue(object oPC, string sSystem, string sKey, string sValue);
void OrderPlayerSystemRemoveValue(object oPC, string sSystem, string sKey);
string OrderPlayersystemObjectGetValueString(object oPC, string sSystem, string sKey);
int OrderPlayerSystemGetValueInt(object oPC, string sSystem, string sKey);
// main Item functions
void OrderItemAddValue(object oItem, string sKey, string sValue);
void OrderItemRemoveValue(object oItem, string sKey);
string OrderItemGetValueString(object oItem, string sKey);
int OrderItemGetValueInt(object oItem, string sKey);
void DeleteItem(object oItem);

// Add an Item system value
void OrderPlayerSystemSetValue(object oPC, string sSystem, string sKey, string sValue) {
  NWNX_Redis_HMSET(OrderUniqueObjectEdge(oPC)+":"+sSystem,sKey,sValue);
}

// Remove an Item system value value
void OrderPlayerSystemRemoveValue(object oPC, string sSystem, string sKey) {
  NWNX_Redis_HDEL(OrderUniqueObjectEdge(oPC)+":"+sSystem,sKey);
}

// Get an Item system value as string
string OrderPlayersystemObjectGetValueString(object oPC, string sSystem, string sKey) {
  int zReturn = NWNX_Redis_HGET(OrderUniqueObjectEdge(oPC)+":"+sSystem,sKey);
  string sReturn = NWNX_Redis_GetResultAsInt(zReturn);
  return sReturn;
}

// Get an Item system value as int
int OrderPlayerSystemGetValueInt(object oPC, string sSystem, string sKey) {
  int zReturn = NWNX_Redis_HGET(OrderUniqueObjectEdge(oPC)+":"+sSystem,sKey);
  int sReturn = NWNX_Redis_GetResultAsInt(zReturn);
  return sReturn;
}

// Add a value to the core quest hash
void OrderItemAddValue(object oItem, string sKey, string sValue) {
  NWNX_Redis_HMSET(OrderUniqueObjectEdge(oItem),sKey,sValue);
}

// Remove a value from the core quest hash
void OrderItemRemoveValue(object oItem, string sKey) {
  NWNX_Redis_HDEL(OrderUniqueObjectEdge(oItem),sKey);
}

// Get a string from the core quest hash
string OrderItemGetValueString(object oItem, string sKey) {
  int zReturn = NWNX_Redis_HGET(OrderUniqueObjectEdge(oItem),sKey);
  string sReturn = NWNX_Redis_GetResultAsString(zReturn);
  return sReturn;
}

// Get a int from the core quest hash
int OrderItemGetValueInt(object oItem, string sKey) {
  int zReturn = NWNX_Redis_HGET(OrderUniqueObjectEdge(oItem),sKey);
  string sReturn = NWNX_Redis_GetResultAsInt(zReturn);
  return sReturn;
}

// Remove an item
void DeleteItem(object oItem) {
  NWNX_Redis_HDEL(OrderUniqueObjectEdge(oItem));
}