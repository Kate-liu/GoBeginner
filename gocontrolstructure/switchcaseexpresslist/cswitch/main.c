#include <stdio.h>

// switch case 语句多个条件匹配，实现同一个逻辑
void check_work_day(int a) {
    switch (a) {
        case 1:
        case 2:
        case 3:
        case 4:
        case 5:
            printf("it is a work day\n");
            break;
        case 6:
        case 7:
            printf("it is a weekend day\n");
            break;
        default:
            printf("do you live on earth?\n");
    }
}