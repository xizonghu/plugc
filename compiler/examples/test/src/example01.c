
struct ABC {
    char *str;
    int len;
};

static int funAdd(int a, int b) {
    if (a > 0) {
        return (a + b);
    }
    return a - b;
}

int strlens(char *str) {
    int len;
    while(*str++) {
        len++;
    }
    return len;
}

void changeABCname(struct ABC *def, char *name) {
    def->str = name;
    def->len = strlen(def->str);
}

void main() {
    int a = -2;
    char b = 4;
    struct ABC abc = {"hello", strlens("hello")};
    abc.len = 2;
    switch (abc.len) {
        case 2:
            abc.len = 3;
        case 3: {
            abc.len = 4;
            break;
        }
        default: {
            abc.len = 5;
        }
    }
    for (b=3; b < 6; b++) {
        if (b==4) continue;
        abc.len++;
    }
    printf("res=%d\n", funAdd(a, b));
    printf("len=%d, b=%d\n", abc.len, b);
    changeABCname(&abc, "hello, world");
    printf("abc=%s, %d\n", abc.str, abc.len);
}