#include "nwnx_events"

void main() {
  object oMod = GetModule();

  // set the examine events here
  NWNX_Events_SubscribeEvent("NWNX_ON_EXAMINE_OBJECT_BEFORE", "order_examine");

  // other seed data will be put here.

}