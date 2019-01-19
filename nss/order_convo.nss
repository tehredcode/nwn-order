struct STRUCT_DLG_INFO dlg_GetDialogInfo(string sDialogTag)
{
    struct STRUCT_DLG_INFO strResult;
    string sSQL = "SELECT " + CS_DLG_ID + ", " + CS_DLG_TAG + ", " + CS_DLG_DESCRIPTION + " FROM ";
    sSQL += CS_DLG_MAINTABLE + " WHERE " + CS_DLG_TAG + " = ?";
    NWNX_SQL_PrepareQuery(sSQL);
    NWNX_SQL_PreparedString(0, sDialogTag);
    NWNX_SQL_ExecutePreparedQuery();

    if (NWNX_SQL_ReadyToReadNextRow())
    {
        NWNX_SQL_ReadNextRow();
        strResult.ID = StringToInt(NWNX_SQL_ReadDataInActiveRow(0));
        strResult.Tag = NWNX_SQL_ReadDataInActiveRow(1);
        strResult.Description = NWNX_SQL_ReadDataInActiveRow(2);
    }
    else
    {
        strResult.ID = 0;
    }

    return (strResult);
}

struct STRUCT_DLG_INFO dlg_GetDialogInfo(object oPC, string sDialogTag, int nTimeout="") {
  string sHash = OrderPlayerUUID(oPC)+":convo:"+sDialogTag);

  if (NWNX_Redis_HEXISTS(sHash,"UUID") != 1) NWNX_Redis_HMSET(sHash,"UUID",OrderGetNewUUID());

  if (nTimeout != "") NWNX_Redis_EXPIRE(sHash,nTimeout);

  int zValue = NWNX_Redis_HMGET(sHash,"UUID");
  string sValue = NWNX_Redis_GetResultAsString(zValue);
  strResult.UUID = sValue;

  zValue = NWNX_Redis_HMGET(sHash,"Tag");
  sValue = NWNX_Redis_GetResultAsString(zValue);
  strResult.Tag = sValue;

  zValue = NWNX_Redis_HMGET(sHash,"Description");
  sValue = NWNX_Redis_GetResultAsString(zValue);
  strResult.Description = sValue;
}

void dlg_StartConversation(string sDialogTag, object oNPC, object oPC, object oAdditionalObject=OBJECT_INVALID)
{
   object oMod = GetModule();
   int iDialog = GetLocalInt(oMod, CS_DLG_DIALOGOBJECT);
   switch (iDialog)
   {
      case 0: iDialog = 1;
           break;
      case CI_DLG_NUMBEROFDIALOGOBJECTS: iDialog = 1;
           break;
      default: iDialog++;
           break;
   }
   string sDialogToUse = "d_dlg_" + IntToString(iDialog);

   struct STRUCT_DLG_INFO strDialog = dlg_GetDialogInfo(sDialogTag);

   SetLocalInt(oPC, CS_DLG_PC_DIALOGID, strDialog.ID);
   SetLocalString(oPC, CS_DLG_PC_SCRIPT, strDialog.Tag);
   SetLocalString(oPC, CS_DLG_PC_NODE, "T");
   SetLocalInt(oPC, CS_DLG_PC_PAGE, 0);
   SetLocalInt(oPC, CS_DLG_PC_DIALOGOBJECT, iDialog);
   SetLocalInt(oPC, CS_DLG_PC_TOKEN, 0);
   SetLocalInt(oPC, CS_DLG_PC_FIRSTOPTIONFORPAGE + "0", 1);
   SetLocalObject(oPC, CS_DLG_PC_CONVERSATIONNPC, oNPC);
   if (GetIsObjectValid(oAdditionalObject))
       SetLocalObject(oPC, CS_DLG_PC_ADDITIONALOBJECT, oAdditionalObject);

   //SendMessageToAllDMs("Starting conversation '" + sDialogTag + "' (ID: #" + IntToString(strDialog.ID) + ")\n" + "Dialog Object: " + sDialogToUse);
   SetLocalInt(oMod, CS_DLG_DIALOGOBJECT, iDialog);

   AssignCommand(oNPC, ActionStartConversation(oPC, sDialogToUse, TRUE));
}