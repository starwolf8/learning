/*
    Program to calculate the 200th triangular number 
    Introduction of the for statement
*/

#include <stdio.h>

int main(int argc, char *argv[]) {

    int n, triangularNumber;

    triangularNumber = 0;

    for (n = 1; n <= 200; n = n + 1 ){
        triangularNumber = triangularNumber + n;

    }

    int count = 1;

    while(count <= 5) {
        printf("%i\n", count);
        ++count;
    }

    printf("The 200th triangular number is %i\n", triangularNumber);

    return 0;
}