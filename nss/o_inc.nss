#include "nwnx_redis"
#include "o_inc_uuid"

const string sServerName = "order";

string OrderObjectEdge(int nType);
string OrderObjectEdge(int nType) {
  switch (nType) {
    case 1: { return sServerName+":server"; }
    case 2: { return sServerName+":door"; }
    case 3: { return sServerName+":placeable"; }
    case 4: { return sServerName+":chat"; }
    default:{ return "err"; }
  }
  return "";
}

// assign and/or return the object uuid.
// currently we support oPC and items.
string OrderUniqueObjectEdge(object oObject);
string OrderUniqueObjectEdge(object oObject){
  // Determine item type
  int nType = GetObjectType(oObject);
  switch(nType) {
    // Creature
    case 1: {
      // if the user has no uuid set
      if (GetTagIsUUID(oObject) == 1) {
        string sUUID = OrderGetNewUUID();
        SetTag(oObject, sUUID);
        return sUUID;
      } else {
        string sUUID = GetTag(oObject);
        return sServerName+":Player:"+sUUID;
      }
    }
    // Item
    case 2: {
      // if the item has no uuid set
      if (GetTagIsUUID(oObject) == 1) {
        string sUUID = OrderGetNewUUID();
        NWNX_Redis_HMSET(sServerName+":Item"+sUUID,"oldTag", GetTag(oObject));
        SetTag(oObject, sUUID);
        return "";
      } else {
        string sUUID = GetTag(oObject);
        return sServerName+":Item"+sUUID;
      }
    }
    default:
    // log("a UUID cannot be added to: " + GetName(oObject) + ":" + GetResRef(oObject));
    return "";
  }
  return "";
}

void main() {}










