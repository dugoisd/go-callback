#include <stdio.h>
#include "../libtest.h"
#include <unistd.h>

void a( int * ctx, const char * a ){
  printf("Hello from GO : %s ,%d\r\n",a,*ctx);
}

int main() {
  printf("C says: about to call Go...\n");
  int b = 125425;
  RegisterCallBack(&a,&b);
  Start();
  sleep(5); // wait 5s 
  printf("C says: Hello \r\n");
  return 0;
}