#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef char *string;

char pop();
void push(char value);

string stack;
// index of the top of the stack in the stack
int top;

int main() {
  top = 0;
  string expression = "5 8 +";
  stack = calloc(sizeof(char), strlen(expression));

  for (int i = 0; i < strlen(expression); i++) {
  }

  printf("len: %d\n", strlen(expression));

  free(stack);
  return 0;
}

char pop() { return stack[top--]; }

void push(char value) { stack[++top] = value; }
