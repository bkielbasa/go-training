#include "test.h"

const char* get_string() {
	return "string sent from C";
}

void print_string(char* a) {
	printf("string sent from Go: %s\n", a);
}