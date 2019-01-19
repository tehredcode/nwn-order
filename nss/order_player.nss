#include "order_inc"

void OrderBountySetValue(object oPC, int nValue) {
  NWNX_Redis_HMSET(OrderPlayerUUID(oPC)+":stats","bounty", IntToString(nValue));
}

void OrderBountyAddValue(object oPC, int nValue) {
  int sInitialValue = OrderBountGetValue(oPC);
  int nValue = nValue + nFirstValue;
  NWNX_Redis_HMSET(OrderPlayerUUID(oPC)+":stats","bounty", IntToString(nValue));
}

int OrderBountyGetValue(object oPC) {
  int zValue = NWNX_Redis_HMGET(OrderPlayerUUID(oPC)+":stats","bounty");
  int nValue = NWNX_Redis_GetResultAsInt(zValue);
  return nValue;
}

int OrderBountyHasBounty(object oPC) {
  int nValue = NWNX_Redis_HEXISTS(OrderPlayerUUID(oPC)+":stats","bounty");  
  return nValue;
}
