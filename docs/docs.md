# Nwscript functions
## [o_inc.nss](https://github.com/Urothis-nwn-Order/nwn-order/blob/master/nss/o_inc.nss)
```c
string OrderObjectEdge(int nType)
> DockerDemo:Server
```
Returns a string that is the redis hash/location to store none unique object information.

| nType | Object Type |
| :--- | :--- |
| 1 | `server` |
| 2 | `door` |
| 3 | `placeable` |
| 4 | `chat` |

```c
string OrderUniqueObjectEdge(object oObject)
> DockerDemo:Player:6fc7438a87d42b2dec552b4fb81b75a2
```
Returns a string that is the redis hash/location to store unique object information.
***
## [o_inc_uuid.nss](https://github.com/Urothis-nwn-Order/nwn-order/blob/master/nss/o_inc_uuid.nss)
```c
string OrderGetNewUUID()
> 6fc7438a87d42b2dec552b4fb81b75a2
```
Returns a UUID, confirmed not conflicting with existing.
***

## Data manipulation
### [o_inc_player.nss](https://github.com/Urothis-nwn-Order/nwn-order/blob/master/nss/o_inc_player.nss)

> player functions
```c
void OrderPlayerSetValue(object oPC, string sKey, string sValue);
> // OrderPlayerSetValue(oPC,"FavoriteColor", "purple");
``` 
```c
void OrderPlayerRemoveValue(object oPC, string sKey);
> // OrderPlayerRemoveValue(oPC,"FavoriteColor");
```
```c
string OrderPlayerGetValueString(object oPC, string sKey);
> purple // OrderPlayerGetValueString(oPC, "FavoriteColor");
```
```c
int OrderPlayerGetValueInt(object oPC, string sKey);
> 128 // OrderPlayerGetValueInt(oPC, "Currency")
```
```c
void OrderPlayerDeleteCharacter(object oPC);
> // OrderPlayerDeleteCharacter(oPC); will delete the players core redis hash
```

***
> player system functions 

```c
void OrderPlayerSystemSetValue(object oPC, string sSystem, string sKey, string sValue);
> // OrderPlayerSystemSetValue(oPC, "smurfConversion", "prettyStatus", "My hand is blue bro")
```
```c
void OrderPlayerSystemRemoveValue(object oPC, string sSystem, string sKey);
> // this will remove sKey on the sSystem.
```
```c
string OrderPlayersystemObjectGetValueString(object oPC, string sSystem, string sKey);
> "My hand is blue bro" //OrderPlayersystemObjectGetValueString(oPC, "smurfConversion", "prettyStatus");
```
```c
int OrderPlayerSystemGetValueInt(object oPC, string sSystem, string sKey);
> 12 // OrderPlayersystemObjectGetValueString(oPC, "smurfConversion", "smurfPercentage");
```
***

### [o_inc_item.nss](https://github.com/Urothis-nwn-Order/nwn-order/blob/master/nss/o_inc_item.nss)

> item functions
```c
void OrderItemSetValue(object oItem, string sKey, string sValue);
> // OrderItemAddValue(oObject, "oldTag", GetTag(oObject));
```
```c
void OrderItemRemoveValue(object oItem, string sKey);
> // Will remove a value from the base item hash
```
```c
string OrderItemGetValueString(object oItem, string sKey);
> "longsword_1" // OrderItemGetValueString(object oItem, "oldTag");
```
```c
int OrderItemGetValueInt(object oItem, string sKey);
> 624 // OrderItemGetValueInt(oItem, "creaturesSlain");
```
```c
void OrderItemDelete(object oItem);
>  // Deletes all redis information about the item
```

***
> item system functions
```c
void OrderPlayerSystemSetValue(object oPC, string sSystem, string sKey, string sValue);
> // OrderPlayerSystemSetValue(oPC, "soulsBound", "soulCount", "123");
```
```c
void OrderPlayerSystemRemoveValue(object oPC, string sSystem, string sKey);
> // OrderPlayerSystemRemoveValue(object oPC, "soulsBound", "soulCount");
```
```c
string OrderPlayerSystemGetValueString(object oPC, string sSystem, string sKey);
> "2d12" // OrderPlayerSystemGetValueString(oPC, "upgrades", "FireUpgrade");
```
```c
int OrderPlayerSystemGetValueInt(object oPC, string sSystem, string sKey);
> 123 // OrderPlayerSystemGetValueInt(object oPC, "soulsBound", "soulCount");
```

***
### [o_inc_quest.nss](https://github.com/Urothis-nwn-Order/nwn-order/blob/master/nss/o_inc_quest.nss)
Works slightly differently
We are using the resref here as a reference to the hash name relative to the quest. [Example](https://github.com/Urothis-nwn-Order/nwn-order/blob/master/examples/example_quest.nss)

> main quest functions
```c
void OrderQuestSetValue(object oPC, string sQuest, string sEntry, string sValue);
> // OrderQuestSetValue(oPC, sQuest, "Name", "The Pickle Wizard's Coven");
```
```c
void OrderQuestRemoveValue(object oPC, string sQuest, string sEntry);
> // OrderQuestRemoveValue(oPC, sQuest, "Name");
```
```c
string OrderQuestGetValueString(object oPC, string sQuest, string sEntry);
> "The Pickle Wizard's Coven" // OrderQuestGetValueString(oPC, sQuest, "Name");
```
```c
int OrderQuestGetValueInt(object oPC, string sQuest, string sEntry);
> 1 // OrderQuestGetValueInt(oPC, sQuest, "status");
```
```c
void OrderQuestDelete(object oPC, string sQuestName);
> // OrderQuestDelete(object oPC, string sQuestName);
```
***

> quest object functions
```c
void OrderQuestObjectSetValue(object oPC, string sQuest, string sResRef, string sEntry, string sValue);
> // OrderQuestObjectSetValue(oPC,sQuest,"dog","wantPets","1");
```
```c
void OrderQuestObjectRemoveValue(object oPC, string sQuest, string sResRef, string sEntry);
> // OrderQuestObjectRemoveValue(oPC,sQuest,"dog","wantPets");
```
```c
string OrderQuestObjectGetValueString(object oPC, string sQuest, string sResRef, string sEntry);
> "pickles" // OrderQuestObjectGetValueString(oPC,sQuest,"dog","name");
```
```c
int OrderQuestObjectGetValueInt(object oPC, string sQuest, string sResRef, string sEntry);
> 1 // OrderQuestObjectGetValueInt(oPC,sQuest,"dog","wantPets");
```
