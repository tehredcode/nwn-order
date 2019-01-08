// -- this is what is triggered via the order heartbeat tickers.
void OrderHeartbeat(string sTicker){
  //Log("heartbeat: "+sTicker,1);
  int nTicker = StringToInt(sTicker);
  switch (nTicker) {   
    case 0:
      // do the thing
    case 5:
      // do the thing
    case 60:
      // do the thing
    case 360:
      // do the thing
    case 720:
      // do the thing
    case 1440:
      // do the thing
    default:
      //Log("Error, ticker not recognized.",1);
    }
}