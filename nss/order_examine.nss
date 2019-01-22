#include "order_quest"
#include "order_players"

////////////////////////////////////////////////////////////////
// should return a list of quests on oPC, intended for dm use //
////////////////////////////////////////////////////////////////
string OrderDescriptionQuestList(object oPC) {
  string sOutput;
  int zCursor = NWNX_Redis_HKEYS(OrderPlayerUUID(oPC)+":quest");
  int i;
  for (i = 0; i < NWNX_Redis_GetArrayLength(zCursor); i++) {
    int zEntry  = NWNX_Redis_GetArrayElement(zCursor, i);
    string sField = NWNX_Redis_GetResultAsString(zEntry);{
    int zValue = NWNX_Redis_HMGet(RdsEdgePlayer("player",oPC)+":quest"+sQuest+":objects"+sResRef,"ReadableStatus");
    string sValue = NWNX_Redis_GetResultAsString(zValue);
    int zName = NWNX_Redis_HMGet(RdsEdgePlayer("player",oPC)+":quest"+sQuest+":objects"+sResRef,"Name");
    string sName = NWNX_Redis_GetResultAsString(zName);
    // example output "  The Weed Wizard's Coven: You need more herb"
    sOutput += "  " + sName + ": "+ sValue;
    }
  }
  sOutput += "\n";
  sOutput += "\n";
  return sOutput;
}

///////////////////////////////////////
// creature object dm examine return //
///////////////////////////////////////
string ExamineDMObjectCreature() {
  string sOutput;
  sOutput += "DM information:\n";
  sOutput += "UUID: "+GetTag(oObject)+"\n";
  sOutput += "ResRef: "+GetResRef(oObject)+"\n";
  sOutput += "Current Quests:\n";
  sOutput += OrderDescriptionQuestList();
  return sOutput;
}

///////////////////////////////////
// item object dm examine return //
///////////////////////////////////
string ExamineDMObjectItem() {
  string sOutput;

  return sOutput;
}

///////////////////////////////////
// door object dm examine return //
///////////////////////////////////
string ExamineDMObjectDoor(object oDoor) {
  string sOutput;
  string sLocked = GetLockKeyTag(oDoor);
  if (sLocked != "") {}
    sOutput =+ "Key needed to unlock";
    sOutput =+ "  "+sLocked+"\n\n";
  return sOutput;
}

////////////////////////////////////////
// placeable object dm examine return //
////////////////////////////////////////

string ExamineDMObjectPlaceable() {
  string sOutput;

  return sOutput;
}

////////////////////////////
// main dm examine return //
////////////////////////////

string ExamineDMObject(object oObject, int nType) {
  string sOutput;

  switch(nType) {
    // Creature
    case 1:
    return OrderExamineCreatureObject(oExaminee,oExaminer,nEnemy,nDM);
    // Item
    case 2:
    return OrderExamineItem(oExaminee,oExaminer,nDM);
    // Door
    case 8:
    return OrderExamineDoor(oExaminee,oExaminer,nDM);
    // Placeable
    case 64:
    return OrderExaminePlaceable(oExaminee,oExaminer,nDM);

    // not special
    default:
        return;
  }
}

///////////////////////////////////////////
// return a list of quests on the object //
///////////////////////////////////////////
string ExamineQuestObject(string sResref, object oPC) {
  string sOutput;
  sOutput += "Quests:\n";

  int zCursor = NWNX_Redis_HKEYS(RdsEdgePlayer("player",oPC)+":quest"+sField+":"+sResRef);
  int i;
  for (i = 0; i < NWNX_Redis_GetArrayLength(zCursor); i++) {
    int zEntry  = NWNX_Redis_GetArrayElement(zCursor, i);
    string sField = NWNX_Redis_GetResultAsString(zEntry);
    
    int err = NWNX_Redis_EXISTS(RdsEdgePlayer("player",oPC)+":quest"+sQuest+":objects"+sResRef,"ReadableStatus");
    if (err != 0) {
      int zValue = NWNX_Redis_HMGet(RdsEdgePlayer("player",oPC)+":quest"+sQuest+":objects"+sResRef,"ReadableStatus");
      string sValue = NWNX_Redis_GetResultAsString(zValue);
      int zName = NWNX_Redis_HMGet(RdsEdgePlayer("player",oPC)+":quest"+sQuest+":objects"+sResRef,"Name");
      string sValue = NWNX_Redis_GetResultAsString(zName);
      sOutput += "  " + sName + ": "+ sValue;
    }
  }
  sOutput += "\n";
  sOutput += "\n";
  return sOutput;
}

/////////////////////////////
// examine creature object //
/////////////////////////////
string ExamineNPC(object oCreature, int nEnemy) {
  switch(nEnemy) {
    case 0: {
      // npc friendly

    }
    case 1: {
      // npc hostile
      float fCR = GetChallengeRating(oCreature);
      string sOutput = GetName(oCreature);
      sOutput += ExamineQuest(GetResRef(oCreature));
      sOutput += "Class: "     + GetStringByStrRef(StringToInt(Get2DAString("Classes","Label",GetClassByPosition(1, oCreature)))); + ": " + IntToString(GetLevelByPosition(1, oCreature))+"\n";
      sOutput += "Class: "     + GetStringByStrRef(StringToInt(Get2DAString("Classes","Label",GetClassByPosition(2, oCreature))); + ": " + IntToString(GetLevelByPosition(2, oCreature))+"\n";
      sOutput += "Class: "     + GetStringByStrRef(StringToInt(Get2DAString("Classes","Label",GetClassByPosition(3, oCreature)))); + ": " + IntToString(GetLevelByPosition(3, oCreature))+"\n";
      sOutput += "STR: "       + IntToString(GetAbilityScore(oCreature, ABILITY_STRENGTH))+"\n";
      sOutput += "DEX: "       + IntToString(GetAbilityScore(oCreature, ABILITY_DEXTERITY))+"\n";
      sOutput += "CON: "       + IntToString(GetAbilityScore(oCreature, ABILITY_CONSTITUTION))+"\n";
      sOutput += "INT: "       + IntToString(GetAbilityScore(oCreature, ABILITY_INTELLIGENCE))+"\n";
      sOutput += "WIS: "       + IntToString(GetAbilityScore(oCreature, ABILITY_WISDOM))+"\n";
      sOutput += "CHA: "       + IntToString(GetAbilityScore(oCreature, ABILITY_CHARISMA))+"\n";
      sOutput += "AC: "        + IntToString(GetAC(oCreature))+"\n";
      sOutput += "HP: "        + IntToString(GetCurrentHitPoints(oCreature)) + "/" + IntToString(GetMaxHitPoints(oCreature))+"\n";
      sOutput += "BAB: "       + IntToString(GetBaseAttackBonus(oCreature))+"\n";
      sOutput += "Fortitude: " + IntToString(GetFortitudeSavingThrow(oCreature))+"\n";
      sOutput += "Reflex: "    + IntToString(GetReflexSavingThrow(oCreature))+"\n";
      sOutput += "Will: "      + IntToString(GetWillSavingThrow(oCreature))+"\n";
      sOutput += "SR: "        + IntToString(GetSpellResistance(oCreature))+"\n";
      sOutput += "\n";
      sOutput += "\n";
      sOutput += sDescribe;
      return sOutput;
    }
  }
}

/////////////////////
// examine Player //
////////////////////

string ExaminePlayer(object oCreature, int nEnemy) {
  string sOutput;
  switch(nEnemy) {
    case 0: {
      // pc friendly

    }
    case 1: {
      // pc hostile
      float fCR = GetChallengeRating(oCreature);
      string sOutput = GetName(oCreature);
      return sOutput;
    }
  }
  // bounty
  if OrderBountyHasBounty(oPC){
    string sBountyAmmount = IntToString(OrderBountyGetValue(object oPC));
    sOutput += "Bounty: " + sBountyAmmount + "g"
  }

  return sOutput;
}

///////////////////////
// examine creature //
//////////////////////
string OrderExamineCreatureObject(object oExaminee,object oExaminer, int nEnemy, int nDM){
  switch(nEnemy) {
    case 0: sOutput += ExaminePlayer(nEnemy);
    case 1: sOutput += ExamineNPC(nEnemy);
  }
  if (nDM) sOutput += ExamineDMObject();
  return sOutput;
}

/////////////////////////
// examine item object //
/////////////////////////

string ExamineUnidentifiedItem(object oItem, object oPC) {

}

string ExamineIdentifiedItem(object oItem, object oPC) {
  
}

string OrderExamineItem(object oItem, object oPC) {
  string sOutput;

  if (GetStolenFlag(oItem) == 1) {sOutput += "Item Stolen \n\n"} 

  int nIdentified = GetIdentified(oItem);
  switch(nIdentified) {
    case 0: sOutput += ExamineUnidentifiedItem(oItem);
    case 1: sOutput += ExamineIdentifiedItem(oItem);
  }
  return sOutput;
}

//////////////////////////////
// examine door object
//////////////////////////////
void OrderExamineDoor(object oDoor,object oPC, int nDM) {
  string sOutput;
  return sOutput;
}

//////////////////////////////
// examine placeable object
//////////////////////////////
void OrderExaminePlaceable(object oPlaceable,object oPC) {
  string sOutput;
  return sOutput;
}

///////////////////////////////////////////////////////////////////
//  main event when a player opens the description on an object. //
///////////////////////////////////////////////////////////////////
void main()
{
  // examiner
  object oExaminer = OBJECT_SELF;
  
  // examinee
  object oExaminee = NWNX_Object_StringToObject(NWNX_Events_GetEventData("EXAMINEE_OBJECT_ID"));

  // is the object our friend?
  int nDM = GetIsDM(oExaminer);

  // original description
  string sOriginalDescription = GetDescription(oExaminee, TRUE, TRUE);

  string sOutput;

  // Determine item type
  int nType = GetObjectType(oExaminee);
  switch(nType) {
    // Creature
    case 1:
    int nEnemy = GetIsEnemy(oExaminee,oExaminer);
    sOutput += OrderExamineCreatureObject(oExaminee,oExaminer,nEnemy,nDM);
    // Item
    case 2:
    sOutput += OrderExamineItem(oExaminee,oExaminer,nDM);
    // Door
    case 8:
    sOutput += OrderExamineDoor(oExaminee,oExaminer,nDM);
    // Placeable
    case 64:
    sOutput += OrderExaminePlaceable(oExaminee,oExaminer,nDM);
    // not special
    default:
    sOutput += "No additional information available"
  }
  SetDescription(oExaminee, sOriginalDescription+"\n\n"+sOutput, TRUE);
}