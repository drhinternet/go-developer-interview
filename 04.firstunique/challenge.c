#include <strings.h>
#include <stdio.h>

char firstunique(char *str) {
    // todo: return the first unique (non-repeated, through the entire string) char in the string
    //       return 0 if there are no unique characters

    return 0;
}

void printExample(char *str) {
    char ch = firstunique(str);

    printf("in(%s) out(%c; %x)\n", str, ch, ch);
}

int main(int argc, char **argv) {
    printExample(""); // => 0
    printExample("x"); // => 'x'
    printExample("xxx"); // => 0
    printExample("food"); // => 'f'
    printExample("foof"); // => 0
    printExample("ffooood"); // => 'd'
}