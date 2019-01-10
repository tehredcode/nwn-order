#include "nwnx_redis"

int IsCharacterValid() {

}

string OrderGetNewUUID() {

  string sRandomName = RandomName(Random(22));
  string sStringLength = GetStringLength(sRandomName);
  string sRandomLetter = GetStringLeft(RandomName(21), Random(sStringLength));
  
  NWNX_Redis_LINSERT(string key,string where,string pivot,string value);
  return sUUID;
}

// -- return or assign and return the oPC uuid.
string OrderGetUUIDPlayer(object oPC);
string OrderGetUUIDPlayer(object oPC) {
  // if the user has no uuid set
  if (GetTag(oPC) == "") {  
    object oMod = GetModule();

    // -- If in progress, else return ""                                                                                       
    if (nUuidInProgress != 1) {
      // get prepared uuid
      string sUUID = OrderGetNewUUID();

      //set oPC tag to uuid
      SetTag(oPC, sUUID);

      return sUUID;
    } else {
      // if no uuid can be grabbed, return "", which should be filtered from being saved. 
      return "";
    }
  } else {
      string sUUID = GetTag(oPC);
      return sUUID;
  }
}