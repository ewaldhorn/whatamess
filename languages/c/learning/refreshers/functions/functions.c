#include "functions.h"

int addUp(int left, int right) {
    return left + right;
}

int main() {
    printf("Adding up 1 and 2 gives %d.\n", addUp(1,2));
    return 0;
}