#include <stdio.h>
#include <stdlib.h>

int main() {
    int x = 10;
    double y = 3.0;
    x=y;

    printf("x=%d\n",x);
    printf("y=%lf\n",y);

    printf("\nThe address of x is %d", &x);

    return 0;
}