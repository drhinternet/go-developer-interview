#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void convertToLowercase(char *str) {
    // todo, convert in-place
}

void printExample(char *in) {
    char *str = malloc(strlen(in) + 1);
    strcpy(str, in);

    printf("in(%s) ", str);
    convertToLowercase(str);
    printf("out(%s)\n", str);
}

int main(int argc, char **argv) {
    printExample("");
    printExample("foobar");
    printExample("FooBar");
    printExample("Hey this has TEXT! Also other symbols?");
}