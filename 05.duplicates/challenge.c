#include <stdio.h>
#include <stdlib.h>

struct duplicateNumbers {
    int *data;
    int num;
};

struct duplicateNumbers findDuplicates(int in[], int num) {
    // todo: return a list of numbers that are duplicated in the +in+ array
    //       the duplicates need not be consecutive

    struct duplicateNumbers result = {
        .data = NULL,
        .num = 0,
    };

    return result;
}

void printExample(int in[], int num) {
    printf("in(");
    for (int ii = 0; ii < num; ++ii ) {
        if ( ii > 0 ) printf(",");
        printf("%d", in[ii]);
    }
    printf(")");

    struct duplicateNumbers dups = findDuplicates(in, num);

    printf(" out(");
    for (int ii = 0; ii < dups.num; ++ii) {
        if (ii > 0) printf(",");
        printf("%d", dups.data[ii]);
    }
    printf(")");

    printf("\n");
}

int main(int argc, char **argv) {
    int arr1[] = {0};
    printExample(arr1, 0);

    int arr2[] = {1,2,3};
    printExample(arr2, sizeof(arr2)/sizeof(int));

    int arr3[] = {1,1,1};
    printExample(arr3, sizeof(arr3)/sizeof(int));

    int arr4[] = {5, 4, 5, 3, 2, 1, 4};
    printExample(arr4, sizeof(arr4)/sizeof(int));
}