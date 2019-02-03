#include "nwnx_redis"
#include "o_inc"

// Add a quest object value
void OrderQuestObjectSetValue(object oPC, string sQuest, string sResRef, string sKey, string sValue) {
  NWNX_Redis_HMSET(OrderUniqueObjectEdge(oPC)+":quest:"+sQuest+":objects:"+sResRef,sKey, sValue);
}

// Remove a quest object value
void OrderQuestObjectRemoveValue(object oPC, string sQuest, string sResRef, string sEntry) {
  NWNX_Redis_HDEL(OrderUniqueObjectEdge(oPC)+":quest:"+sQuest+":objects:"+sResRef,sEntry);
}

// Get a quest object value as string
string OrderQuestObjectGetValueString(object oPC, string sQuest, string sResRef, string sEntry) {
  int zReturn = NWNX_Redis_HGET(OrderUniqueObjectEdge(oPC)+":quest:"+sQuest+":objects:"+sResRef,sEntry);
  string sReturn = NWNX_Redis_GetResultAsString(zReturn);
  return sReturn;
}

// Get a quest object value as int
int OrderQuestObjectGetValueInt(object oPC, string sQuest, string sResRef, string sEntry) {
  int zReturn = NWNX_Redis_HGET(OrderUniqueObjectEdge(oPC)+":quest:"+sQuest+":objects:"+sResRef,sEntry);
  int sReturn = NWNX_Redis_GetResultAsInt(zReturn);
  return sReturn;
}

// Add a value to the core quest hash
void OrderQuestSetValue(object oPC, string sQuest, string sEntry, string sValue) {
  NWNX_Redis_HMSET(OrderUniqueObjectEdge(oPC)+":quest:"+sQuest,sEntry,sValue);
}

// Remove a value from the core quest hash
void OrderQuestRemoveValue(object oPC, string sQuest, string sEntry, string sValue) {
  NWNX_Redis_HDEL(OrderUniqueObjectEdge(oPC)+":quest:"+sQuest,sEntry);
}

// Get a string from the core quest hash
string OrderQuestGetValueString(object oPC, string sQuest, string sEntry) {
  int zReturn = NWNX_Redis_HGET(OrderUniqueObjectEdge(oPC)+":quest:"+sQuest,sEntry);
  string sReturn = NWNX_Redis_GetResultAsString(zReturn);
  return sReturn;
}

// Get a int from the core quest hash
int OrderQuestGetValueInt(object oPC, string sQuest, string sEntry) {
  int zReturn = NWNX_Redis_HGET(OrderUniqueObjectEdge(oPC)+":quest:"+sQuest,sEntry);
  int nReturn = NWNX_Redis_GetResultAsInt(zReturn);
  return nReturn;
}

// Remove a quest, Yes this deletes everything.
void OrderQuestDeletet(object oPC, string sQuest) {
  NWNX_Redis_DEL(OrderUniqueObjectEdge(oPC)+":quest:"+sQuest);
}
