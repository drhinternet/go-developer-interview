#include <strings.h>
#include <stdio.h>

// Return the first unique (non-repeated) character in the string
char firstunique(char *str) {
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