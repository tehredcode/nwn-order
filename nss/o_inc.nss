const string sServerName = GetModuleName();

string OrderObjectEdge(int nType);
string OrderObjectEdge(int nype) {
  switch (nType) {
    case 1: { return sNwserver+":server"; }        
    case 2: { return sNwserver+":door"; }
    case 3: { return sNwserver+":placeable"; }  
    case 4: { return sNwserver+":chat"}   
    default:{ return "err"; }
  }
}

// assign and/or return the object uuid.
// currently we support oPC and items.
string OrderUniqueObjectEdge(object oObject);
string OrderUniqueObjectEdge(object oObject){
  // Determine item type
  int nType = GetObjectType(oObject);
  switch(nType) {
    // Creature
    case 1: {}
      // if the user has no uuid set
      if (GetTagIsUUID(oObject) == 1) {  
        string sUUID = OrderGetNewUUID();
        SetTag(oObject, sUUID);
        return sUUID;
      } else {
        string sUUID = GetTag(oObject);
        return sServerName+":Player:"+sUUID;
      }

    // Item
    case 2:
      // if the item has no uuid set
      if (GetTagIsUUID(oObject) == 1) {  
        string sUUID = OrderGetNewUUID();
        OrderItemAddValue(oObject, "oldTag", GetTag(oObject));
        SetTag(oObject, sUUID);
        return sServerName+":Item"+sUUID;
      } else {
        string sUUID = GetTag(oObject);
        return sServerName+":Item"+sUUID;
      }
    default:
    // log("a UUID cannot be added to: " + GetName(oObject) + ":" + GetResRef(oObject));
    return "";
  }
}