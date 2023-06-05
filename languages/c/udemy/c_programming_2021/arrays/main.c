#include <stdlib.h>
#include <stdio.h>

int main() {

    // 1d array
    int nums[10];

    nums[0] = 1; // 1st element
    nums[9] = 10; // last element

    printf("\n1d Array of numbers:\n");
    for (int i = 0; i < 10; i++) {
        printf("Position %2d contains %d\n", i, nums[i]);
    }

    return 0;
}