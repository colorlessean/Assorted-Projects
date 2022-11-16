// Hello world! Cplayground is an online sandbox that makes it easy to try out
// code.

#include <stdio.h>
#include <stdlib.h>

// // lamdba at addr 0x1000
// void lamdba(int str) {
//     printf("%d\n", str);
// }

// void func(void (*lamda)(int), int val) {
//     (*lamda)(val);
// }

// int main() {
//     // lambda_ptr will hold address 0x1000 and tell us we need to provide a single int parameter on the stack
//     void (*lamdba_ptr)(int)  = &lamdba;
    
//     (*lamdba_ptr)(1);
    
//     func(lamdba_ptr, 10);
//     func(lamdba_ptr, 20);
    
//     return 0;
// }

bool less_than(int x, int y) {
    return x < y;
}

bool greater_than(int x, int y) {
    return x > y;
}

void sort(int * list, size_t size, bool (*compare)(int, int)) {
    for (int i = 0; i < size; i++) {
        for (int j = 0; j < size; j++) {
            if ((*compare)(list[i], list[j])) {
                int temp = list[i];
                list[i] = list[j];
                list[j] = temp;
            }
        }
    }
}

void print(int * list, size_t size) {
    for (int i = 0; i < size; i++) {
        printf("%d\n", list[i]);
    }
}

int main() {
    int list[] = {4, 2, 1, 0};
    
    bool (*compare)(int, int) = &less_than;
    
    sort(list, 4, compare);
    print(list, 4);
    
    bool (*compare1)(int, int) = &greater_than;
    
    sort(list, 4, compare1);
    print(list, 4);
    
}

