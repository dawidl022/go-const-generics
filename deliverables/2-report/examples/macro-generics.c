#include <stdio.h>

#define ARRAY(TYPE, SIZE, NAME) typedef struct { \
    TYPE x[SIZE]; \
} NAME; \
\
TYPE NAME##_first(NAME s) { \
    return s.x[0]; \
}


ARRAY(int, 5, Foo)

ARRAY(char*, 2, Bar)

int main(void) {
    Foo f = { 1, 2, 3, 4, 5 };
    printf("%d\n", Foo_first(f));

    Bar b = { "hello", "world" };
    printf("%s\n", Bar_first(b));
}
