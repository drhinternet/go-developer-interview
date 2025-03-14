#include <stdio.h>
#include <stdlib.h>

int *generatePrimes(int n) {
    // todo: generate the first n primes, return the pointer to the array
}

int main(int argc, char **argv) {
    int *primes;

    for ( int n = 0; n <= 10; ++n ) {
        int *primes = generatePrimes(n);

        printf("%2d primes:", n);

        for ( int pn = 0; pn < n; ++pn ) {
            printf(" %d", primes[pn]);
        }

        printf("\n");
    }
}
