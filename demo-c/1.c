#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int main() {
    struct data {
        int i;
        int j;
        int k;
    };

    struct data v1;
    struct data *dsptr;
    dsptr = (struct data*)malloc(sizeof(struct data));

    printf("Size of struct data=%d\n", sizeof(struct data));
    printf("Address of number int i=%zu\n", &(dsptr->i));
    printf("Address os number int j=%zu\n", &(dsptr->j));
    printf("Address of number int k=%zu\n", &(dsptr->k));
    printf("ddd k=%zu\n", &(dsptr->j))

    return 0;
}