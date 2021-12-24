// 宏定义常量
#define FILE_MAX_LEN 0x22334455
#define PI 3.1415926
#define GO_GREETING "Hello, Gopher"
#define A_CHAR 'a'

// const 关键字修饰的标识符（本质依然是常量）
const int size = 5;
int a[size] = {1,2,3,4,5}; // size 本质不是常量，这将导致编译器错误

// 枚举类型
enum Weekday {
    SUNDAY,
    MONDAY,
    TUESDAY,
    WEDNESDAY,
    THURSDAY,
    FRIDAY,
    SATURDAY
};

int main() {
    enum Weekday d = SATURDAY;
    printf("%d\n", d); // 6
}