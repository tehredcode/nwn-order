#include "order_quest"

string ExamineDM(object oObject, int nType) {
  string sOutput;
  sOutput += "DM Only Information:\n";
  sOutput += "UUID: "+GetTag(oObject);
  sOutput += "ResRef: "+GetResRef(oObject);
           
  string sHash = sObjectUUID(nType,oObject);
  int zCursor = NWNX_Redis_HKEYS(sHash);
  int i; for (i = 0; i < NWNX_Redis_GetArrayLength(zCursor); i++) { 
    int zEntry  = NWNX_Redis_GetArrayElement(zCursor, i);
    string sField = NWNX_Redis_GetResultAsString(zEntry);

    int zValue = NWNX_Redis_HMGET(oObject,sField);
    string sValue = NWNX_Redis_GetResultAsString(zEntry);

    sOutput += GetName(oObject)":"+sField +": "+sValue+"\n";
  }
  sOutput += "\n";
  sOutput += "\n";
  return sOutput;
}

// should return a list of quests on the object 
string ExamineQuest(string sResref, object oPC) {
  string sOutput;
  sOutput += "Quests:\n";

  RdsEdgePlayer("player",oPC)+":quest"+sField+":"+sResRef);

  int zCursor = NWNX_Redis_HKEYS(OrderPlayerUUID(oPC)+":quest");
  int i;
  for (i = 0; i < NWNX_Redis_GetArrayLength(zCursor); i++) {
    int zEntry  = NWNX_Redis_GetArrayElement(zCursor, i);
    string sField = NWNX_Redis_GetResultAsString(zEntry);
    
    int err = NWNX_Redis_EXISTS(RdsEdgePlayer("player",oPC)+":quest"+sQuest+":objects"+sResRef,"ReadableStatus");
    if (err != 0) {
      int zValue = NWNX_Redis_HMGet(RdsEdgePlayer("player",oPC)+":quest"+sQuest+":objects"+sResRef,"ReadableStatus");
      string sValue = NWNX_Redis_GetResultAsString(zValue);
      int zName = NWNX_Redis_HMGet(RdsEdgePlayer("player",oPC)+":quest"+sQuest+":objects"+sResRef,"Name");
      sOutput += "  " + sName + ": "+ sValue;
    }
  }
  sOutput += "\n";
  sOutput += "\n";
  return sOutput;
}

// examine creature yo
void OrderExamineCreature(object oCreature,object oPC){
    string sDescribe = GetDescription(oCreature, TRUE, TRUE);
    float fCR = GetChallengeRating(oCreature);

    string sOutput = GetName(oCreature);
        // debug information
        if (IsoPCDM(oPC) == 1) {
            sOutput += ExamineDM(); 
        }

        sOutput += ExamineQuest(GetResRef(oCreature));

        // creature information
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
    SetDescription(oCreature, sOutput, TRUE);
}

// examine item
void OrderExamineItem(object oItem, object oPC) {

}

// examine door
void OrderExamineDoor(object oDoor,object oPC) {

}

// examine placeable
void OrderExaminePlaceable(object oPlaceable,object oPC) {

}

//  main event when a player opens the description on an object.
void main()
{
  // examiner
  object oPC = OBJECT_SELF;
  
  // examinee
  object oExaminee = NWNX_Object_StringToObject(NWNX_Events_GetEventData("EXAMINEE_OBJECT_ID"));

  // Determine item type
  int nType = GetObjectType(oExaminee);
  switch(nType) {
    // Creature
    case 1:
    OrderExamineCreature(GetResRef(oExaminee),oPC);
    // Item
    case 2:
    OrderExamineItem(oExaminee,oPC);
    // Door
    case 8:
    OrderExamineDoor(oExaminee,oPC);
    // Placeable
    case 64:
    OrderExaminePlaceable(oExaminee,oPC);

    // not special
    default:
        return;
  }
}