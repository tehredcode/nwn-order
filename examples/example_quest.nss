#include "order_quest"
#include "nwnx_time"

// this would be an internal name, no spaces, no capitals.
const string sQuest = "420blazeit";

const string sQuestObject1 = "magic_herb";
const string sQuestObject2 = "relic_of_cheech";
const string sQuestObject3 = "the_high_wizard";

int QuestOnePhaseOne(object oPC) {
  //seed some values into the quest
  // quest main hash stuff
  OrderQuestSetValue(oPC, sQuest, "Name", "The Weed Wizard's Coven");
  OrderQuestSetValue(oPC, sQuest, "MainObjective", "Get to the smokey mountains, attempt to gain entry to the high tower.");
  OrderQuestSetValue(oPC, sQuest, "Status", "1");
  OrderQuestSetValue(oPC, sQuest, "StartTime", IntToString(NWNX_Time_GetTimeStamp());
  OrderQuestSetValue(oPC, sQuest, "UUID", OrderObjectUUID(6));  
  OrderQuestSetValue(oPC, sQuest, "ReadableStatus", "You need more herb");
  
  // quest object hash stuff
  OrderQuestObjectSetValue(oPC,sQuest,sQuestItem1,"ObjectiveType","Obtain");
  OrderQuestObjectSetValue(oPC,sQuest,sQuestItem1,"status","1");
  OrderQuestObjectSetValue(oPC,sQuest,sQuestItem2,"ObjectiveType","Obtain");
  OrderQuestObjectSetValue(oPC,sQuest,sQuestItem2,"status","1");
  OrderQuestObjectSetValue(oPC,sQuest,sQuestItem3,"ObjectiveType","Kill");
  OrderQuestObjectSetValue(oPC,sQuest,sQuestItem3,"status","1");
  return 1;
}

int QuestOnePhaseTwo(object oPC) {
  OrderQuestAddValue(oPC, sQuest, "ReadableStatus", "How will I smoke all this herb");
  OrderQuestObjectAddValue(oPC,sQuest,sQuestItem1,"status","2");
  return 1;
}

int QuestOnePhaseThree(object oPC) {
  OrderQuestAddValue(oPC, sQuest, "ReadableStatus", "Who stole all my herb");
  OrderQuestObjectAddValue(oPC,sQuest,sQuestItem2,"status","2");
  return 1;
}

int QuestOnePhaseFour(object oPC) {
  OrderQuestAddValue(oPC, sQuest, "ReadableStatus", "Lit bro");
  OrderQuestObjectAddValue(oPC,sQuest,sQuestItem3,"status","2");
  GiveReward();
  return 1;
}

int ExampleQuest1Check(object oPC) {
  int nObjectCheck1 = OrderQuestObjectGetValueInt(oPC,sQuest, sQuestItem1, Status);
  if (nObjectCheck1 == 2 && nObjectCheck2 == 2 && nObjectCheck3 ==2) {
    return 4;
  }
  if (nObjectCheck1 == 2 && nObjectCheck2 == 2 && nObjectCheck3 ==1) {
    return 3;
    }
  if (nObjectCheck1 == 2 && nObjectCheck2 == 1 && nObjectCheck3 ==1) {
    return 2;
    }
  if (nObjectCheck1 == 1 && nObjectCheck2 == 1 && nObjectCheck3 ==1) {
    return 1;
  }
}