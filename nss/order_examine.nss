#include "order_quest"
#include "order_players"

// should return a list of quests on oPC, intended for dm use
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

ExamineDMObjectCreature() {
  string sOutput;
  sOutput += "DM information:\n";
  sOutput += "UUID: "+GetTag(oObject)+"\n";
  sOutput += "ResRef: "+GetResRef(oObject)+"\n";
  sOutput += "Current Quests:\n";
  sOutput += OrderDescriptionQuestList();

  return sOutput;
}

ExamineDMObjectItem() {}

ExamineDMObjectDoor() {}

ExamineDMObjectPlaceable() {}

string ExamineDMObject(object oObject, int nType) {
  string sOutput;

  switch(nType) {
    // Creature
    case 1:
    OrderExamineCreatureObject(oExaminee,oExaminer,nEnemy,nDM);
    // Item
    case 2:
    OrderExamineItem(oExaminee,oExaminer,nDM);
    // Door
    case 8:
    OrderExamineDoor(oExaminee,oExaminer,nDM);
    // Placeable
    case 64:
    OrderExaminePlaceable(oExaminee,oExaminer,nDM);

    // not special
    default:
        return;
  }
  return sOutput;
}

// should return a list of quests on the object 
string ExamineQuestObject(string sResref, object oPC) {
  string sOutput;
  sOutput += "Quests:\n";

  RdsEdgePlayer("player",oPC)+":quest"+sField+":"+sResRef);

  int zCursor = NWNX_Redis_HKEYS(OrderPlayerUUID(oPC)+":quest");
  int i;
  for (i = 0; i < NWNX_Redis_GetArrayLength(zCursor); i++) {
    int zEntry  = NWNX_Redis_GetArrayElement(zCursor, i);
    string sField = NWNX_Redis_GetResultAsString(zEntry);
    sOutput += 
    }
  }
  sOutput += "\n";
  sOutput += "\n";
  return sOutput;
}

//////////////////////////////
// examine item object
//////////////////////////////
void OrderExamineItem(object oItem, object oPC) {
  string sDescribe = GetDescription(oExaminee, TRUE, TRUE);
  int nEnemy = GetIsEnemy(oExaminee,oExaminer);
  switch(nEnemy) {
    case 0: sOutput += ExaminePlayer();
    case 1: sOutput += ExamineNPC(nEnemy);
    }
  }
  SetDescription(oExaminee, sOutput, TRUE);
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


//////////////////////////////
// examine creature object
//////////////////////////////
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

string ExaminePlayer(object oCreature, int nEnemy) {
  string sOutput;

  // bounty
  if OrderBountyHasBounty(oPC){
    string sBountyAmmount = IntToString(OrderBountyGetValue(object oPC));
    sOutput += "Bounty: " + sBountyAmmount + "g"
  }

  return sOutput;
}

// examine creature yo
void OrderExamineCreatureObject(object oExaminee,object oExaminer, int nEnemy, int nDM){
  string sDescribe = GetDescription(oExaminee, TRUE, TRUE);
  switch(nEnemy) {
    case 0: sOutput += ExaminePlayer();
    case 1: sOutput += ExamineNPC(nEnemy);
  }
  if (nDM) sOutput += ExamineDMObject();
  SetDescription(oExaminee, sOutput, TRUE);
}

//////////////////////////////
//  main event when a player opens the description on an object.
//////////////////////////////
void main()
{
  // examiner
  object oExaminer = OBJECT_SELF;
  
  // examinee
  object oExaminee = NWNX_Object_StringToObject(NWNX_Events_GetEventData("EXAMINEE_OBJECT_ID"));

  // is the object our friend?
  int nDM = GetIsDM(oExaminer);

  // Determine item type
  int nType = GetObjectType(oExaminee);
  switch(nType) {
    // Creature
    case 1:
    int nEnemy = GetIsEnemy(oExaminee,oExaminer);
    OrderExamineCreatureObject(oExaminee,oExaminer,nEnemy,nDM);
    // Item
    case 2:
    OrderExamineItem(oExaminee,oExaminer,nDM);
    // Door
    case 8:
    OrderExamineDoor(oExaminee,oExaminer,nDM);
    // Placeable
    case 64:
    OrderExaminePlaceable(oExaminee,oExaminer,nDM);

    // not special
    default:
        return;
  }
}