#include <stdio.h>

int main() {
    int i;
    int sum = 0;
    for (i = 0; i < 10; i++) {
        sum += i;
    }
    printf("%d\n", sum);  // 45
}