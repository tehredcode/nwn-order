//#include "order_inc"
//
//void ItemUnaquireTracking(){
//
//}
//
//void ItemAquireTracking() {
//
//}
//
//void initItemTracking(object oItem) {
//  string sItemUUID = sItemUUID(oItem);
//  string sItemType = GetBaseItemType(oItem);
//  string sItemDescription = GetDescription(oItem,0,1);
//  string sOriginalTag = GetTag(oItem);
//
//
//  itemproperty ipLoop = GetFirstItemProperty(oItem);
//  while (GetIsItemPropertyValid(ipLoop)) {
//
//    ipLoop=GetNextItemProperty(oItem);
//  }
//}
//
//string sItemUUID(object oItem);
//string sItemUUID(object oItem) {
//  string sLocalUUID = GetLocalString(oItem, "uuid");
//  if (sLocalUUID != "") {
//    return sLocalUUID;
//  } else {
//    string sNewUUID = OrderGetUUID();
//    SetLocalString(oItem, "uuid", sNewUUID);
//    return sNewUUID;
//  }
//}
//
//int nItemIsMeleeWeapon(object oItem);
//int nItemIsMeleeWeapon(object oItem) {
//    if (IPGetIsMeleeWeapon(oItem)) {
//        return 1;
//    }
//    return 0;
//}
//
//int nItemRangedWeapon(object oItem);
//int nItemRangedWeapon(object oItem) {
//    if (GetWeaponRanged(oItem)) {
//        return 1;
//    }
//    return 0;
//}
//
//Extensive knowledge of infrastructure best practices and customer relations.
//Earlier career was focused on Windows environments. But I would love to transition to a Linux based work environment and get the exposure I need to grow.
//
//Personal experience has always been Linux and programming. I deploy and maintain multiple projects in my homelab, using Docker/Kubernetes for applications, and Ceph for storage. All with build processes using circle and and a personal program I use to detect image changes and deploy them safely. 
//
//I also have programming knowledge using golang, ruby, python, and bash.
//
//I would love to join a work environment where I can grow and flourish and help the business do the same thing.