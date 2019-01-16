#include "nwnx_time"
#include "nwnx_events"

#include "order_inc"
#include "order_alert"
#include "nwnx_chat"

void main() {
  object oMod = GetModule();

  // set the examine events here
  NWNX_Events_SubscribeEvent("NWNX_ON_EXAMINE_OBJECT_BEFORE", "order_examine");
  //NWNX_Events_SubscribeEvent("NWNX_ON_EXAMINE_OBJECT_AFTER",  "order_examine_a");

  // other seed data will be put here.

}