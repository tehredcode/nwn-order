

string OrderPrettyClasses(int nClass) {
  switch(nClass) {
    case 0:  return "Barbarian";	
    case 3:  return "Druid";	
    case 4:  return "Fighter";	
    case 5:  return "Monk";	
    case 7:  return "Ranger";	
    case 8:  return "Rogue";	
    case 9:  return "Sorcerer";	
    case 10: return "Wizard";	
    case 11: return "Abberation";	
    case 12: return "Animal";	
    case 13: return "Construct";	
    case 14: return "Humanoid";	
    case 16: return "Elemental";	
    case 17: return "Fey";	
    case 18: return "Fragon";	
    case 19: return "Undead";	
    case 20: return "Commoner";	
    case 21: return "Beast";	
    case 22: return "Giant";	
    case 23: return "Magical Beast";	
    case 25: return "Shapechanger";	
    case 26: return "Vermin";		
    case 27: return "Shadowdancer";	
    case 28: return "Harper Scout";	
    case 29: return "Arcane Archer";	
    case 30: return "Assassin";	
    case 31: return "Blackguard";	
    case 32: return "Champion of Torm";	
    case 33: return "Weapon Master";	
    case 34: return "Pale Master";		
    case 35: return "Shifter";		
    case 36: return "Dwarven Defender";	
    case 37: return "Red Dragon Disciple";	
    case 41: return "Purple Dragon Knight";	
  }
  return "err"
}

// examine creature yo
void OrderExamineCreature(object oCreature){
    string sDescribe = GetDescription(oCreature, TRUE, TRUE);
    float fCR = GetChallengeRating(oCreature);

    string sOutput = "<cóó >CR Value:</c> " + FloatToString(fCR, 0, 2);
    sOutput += "\n";
    sOutput += "\n";
    sOutput += "Class: "     + OrderPrettyClasses(GetClassByPosition(1, oCreature)) + ": " + IntToString(GetLevelByPosition(1, oCreature))+"\n";
    sOutput += "Class: "     + OrderPrettyClasses(GetClassByPosition(2, oCreature)) + ": " + IntToString(GetLevelByPosition(2, oCreature))+"\n";
    sOutput += "Class: "     + OrderPrettyClasses(GetClassByPosition(3, oCreature)) + ": " + IntToString(GetLevelByPosition(3, oCreature))+"\n";
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
void OrderExamineItem(object oItem) {

  // just playing with some arrays and hashes stuff here. Won't work yet.  
  //int zCursor = NWNX_Redis_HKEYS(sServerStatHash);
  //string sOutput 
  //int i; for (i = 0; i < NWNX_Redis_GetArrayLength(zCursor); i++) { 
  //  int zEntry  = NWNX_Redis_GetArrayElement(zCursor, i); 
  //  string sField = NWNX_Redis_GetResultAsString(zEntry);
  //  int zValue = NWNX_Redis_HMGET(sServerStatHash,sField);
  //  string sValue = NWNX_Redis_GetResultAsString(zEntry);
  //  sOutput += sField +": "+sValue+"\n";
  //}
  //WriteTimestampedLogEntry(sOutput);
}

// examine door
void OrderExamineDoor(object oDoor) {

}

// examine placeable
void OrderExaminePlaceable(object oPlaceable) {

}

//  Set the description on the object
void SetDescription(object oObject);
void SetDescription(object oObject)
{
  // Determine item type
  int nType = GetObjectType(oObject);
  switch(nType) {
    // Creature
    case 1:
    OrderExamineCreature(oObject);
    // Item
    case 2:
    OrderExamineItem(oObject);
    // Door
    case 8:
    OrderExamineDoor(oObject);
    // Placeable
    case 64:
    OrderExaminePlaceable(oObject);

    // not special
    default:
        return;
  }
}