#include "nwnx_redis"
#include "order_inc"

void OrderQuestObjectAddValue(object oPC, string sQuest, string sResRef, string sEntry, string sValue);
void OrderQuestObjectRemoveValue(object oPC, string sQuest, string sResRef, string sEntry, string sValue);
string OrderQuestObjectGetValueString(object oPC, string sQuest, string sResRef, string sEntry);
int OrderQuestObjectGetValueInt(object oPC, string sQuest, string sResRef, string sEntry);
void OrderQuestAddValue(object oPC, string sQuest, string sEntry, string sValue);
void OrderQuestRemoveValue(object oPC, string sQuest, string sEntry, string sValue);
string OrderQuestGetValueString(object oPC, string sQuest, string sEntry);
int OrderQuestGetValueInt(object oPC, string sQuest, string sEntry);
void DeleteQuest(object oPC, string sQuestName);

// Add a quest object value
void OrderQuestObjectAddValue(object oPC, string sQuest, string sResRef, string sEntry, string sValue) {
  NWNX_Redis_HMSET(RdsEdgePlayer("player",oPC)+":quest:"+sQuest+":objects"+sResRef,sValue);
}

// Remove a quest object value
void OrderQuestObjectRemoveValue(object oPC, string sQuest, string sResRef, string sEntry, string sValue) {
  NWNX_Redis_HDEL(RdsEdgePlayer("player",oPC)+":quest:"+sQuest+":objects"+sResRef,sEntry);
}

// Get a quest object value as string
string OrderQuestObjectGetValueString(object oPC, string sQuest, string sResRef, string sEntry) {
  int zReturn = NWNX_Redis_HGET(RdsEdgePlayer("player",oPC)+":quest:"+sQuest+":objects"+sResRef,sEntry);
  string sReturn = NWNX_Redis_GetResultAsInt(zReturn);
  return sReturn;
}

// Get a quest object value as int
int OrderQuestObjectGetValueInt(object oPC, string sQuest, string sResRef, string sEntry) {
  int zReturn = NWNX_Redis_HGET(RdsEdgePlayer("player",oPC)+":quest:"+sQuest+":objects"+sResRef,sEntry);
  int sReturn = NWNX_Redis_GetResultAsInt(zReturn);
  return sReturn;
}

// Add a value to the core quest hash
void OrderQuestAddValue(object oPC, string sQuest, string sEntry, string sValue) {
  NWNX_Redis_HMSET(RdsEdgePlayer("player",oPC)+":quest"+sQuest,sEntry,sValue);
}

// Remove a value from the core quest hash
void OrderQuestRemoveValue(object oPC, string sQuest, string sEntry, string sValue) {
  NWNX_Redis_HDEL(RdsEdgePlayer("player",oPC)+":quest"+sQuest,sEntry);
}

// Get a string from the core quest hash
string OrderQuestGetValueString(object oPC, string sQuest, string sEntry) {
  int zReturn = NWNX_Redis_HGET(RdsEdgePlayer("player",oPC)+":quest:"+sQuest+":objects"+sResRef,sEntry);
  string sReturn = NWNX_Redis_GetResultAsString(zReturn);
  return sReturn;
}

// Get a int from the core quest hash
int OrderQuestGetValueInt(object oPC, string sQuest, string sEntry) {
  int zReturn = NWNX_Redis_HGET(RdsEdgePlayer("player",oPC)+":quest:"+sQuest+":objects"+sResRef,sEntry);
  string sReturn = NWNX_Redis_GetResultAsInt(zReturn);
  return sReturn;
}

// Remove a quest, Yes this deletes everything.
void DeleteQuest(object oPC, string sQuestName) {
  NWNX_Redis_HDEL(RdsEdgePlayer("player",oPC)+":quest", sQuestName);
}