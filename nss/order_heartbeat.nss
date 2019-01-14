#include "order_external"

// -- this is what is triggered via the order heartbeat tickers.
void OrderHeartbeat(string sTicker){
  //Log("heartbeat: "+sTicker,1);
  int nTicker = StringToInt(sTicker);
  switch (nTicker) {   
    case 0:
      OrderHeartbeat1();
      break;
    case 5:
      OrderHeartbeat5();
      break;
    case 60:
      OrderHeartbeat60();
      break;
    case 360:
      OrderHeartbeat360();
      break;
    case 720:
      OrderHeartbeat720();
      break;
    case 1440:
      OrderHeartbeat1440();
      break;
    default:
      //Log("Error, ticker not recognized.",1);
      break;
    }
}