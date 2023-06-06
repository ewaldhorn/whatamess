#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

int main() {
    printf("C Strings:\n\n");

    char firstName[] = {"Ben"};
    char lastName[] = {"Benjamin"};

    printf("All together now: %s %s\n", firstName, lastName);
    printf("%s is %lu characters long\n", firstName, strlen(firstName));

    printf("\n");
    return 0;
}