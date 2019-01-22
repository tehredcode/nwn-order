void main() {
  // examinee
  object oExaminee = NWNX_Object_StringToObject(NWNX_Events_GetEventData("EXAMINEE_OBJECT_ID"));

  // set original description
  string sOriginalDescription = GetDescription(oExaminee, TRUE, TRUE);
  SetDescription(oExaminee, sOriginalDescription, TRUE);
}