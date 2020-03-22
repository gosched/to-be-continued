#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

void calculate(char str[], int length, double *answer); // convert infix to postfix, calculate
void operate(char op, double num[], int *numtop);		// 運算數 + - * /

int checkOperatorOrder(char op); // 判斷運算符的優先順序
int divide;						 // check division by zero

int main()
{
	double answer = 0;

	char str[256]; // 表達式
	int length;	//  表達式長度

	while (1)
	{
		printf("輸入表達式 (輸入exit可結束程序):\n");

		scanf("%s", str);
		// fgets(str, 255, stdin);

		printf("\n");
		// fseek(stdin, 0, SEEK_END); // clear terminal

		length = strlen(str);

		if (str[length - 1] == '\n')
		{
			str[length - 1] = '\0';
			length--;
		}

		if (strcmp("exit", str) == 0)
		{
			return 0;
		}
		else if (length == 0)
		{
			printf("表達式無效\n");
		}
		else if (length > 250)
		{
			printf("表達式過長\n");
		}
		else
		{
			divide = 0;
			calculate(str, length, &answer);
			if (divide == 0)
			{
				printf("答案:%f\n\n", answer);
			}
			else
			{
				printf("\n");
				printf("不可除以0\n");
			}
		}
	}
	return 0;
}

int checkOperatorOrder(char op)
{
	int priority;
	switch (op)
	{
	case '*':
	case '/':
		priority = 2;
		break;
	case '+':
	case '-':
		priority = 1;
		break;
	default:
		priority = 0;
		break;
	}
	return priority;
}

void calculate(char str[], int length, double *answer)
{
	int i, j;
	int top = 0, numtop = 0; // top 記錄算子數目, numtop 記錄運算數數目
	double num[100];		 // 運算數 operand
	char stack[100];		 // 算子 operator

	for (i = 0; i < length; i++)
	{
		j = i;

		if ((str[i] <= '9' && str[i] >= '0') || str[i] == '.') // 運算數
		{
			// 考慮運算數 非個位數的情況
			while ((str[j] <= '9' && str[j] >= '0') || str[j] == '.')
			{
				j++;
			}

			numtop++;
			num[numtop] = atof(str + i);
			i = j - 1;
		}
		else // 算子
		{
			switch (str[i])
			{
			case '(':
				top++;
				stack[top] = str[i];
				break;

			case ')':
				while (stack[top] != '(')
				{
					operate(stack[top], num, &numtop);
					top--;
				}
				top--;
				break;

			case '+':
			case '-':
			case '*':
			case '/':
				while (checkOperatorOrder(stack[top]) >= checkOperatorOrder(str[i]))
				{
					operate(stack[top], num, &numtop);
					top--;
				}
				top++;
				stack[top] = str[i];
				break;

			case ' ':
				break;

			default:
				break;
			}
		}
		if (i == length - 1)
		{
			while (top != 0)
			{
				operate(stack[top], num, &numtop);
				top--;
			}
		}
	}
	if (divide == 0) // 除法無誤
	{
		*answer = num[1]; // 算出答案
	}
}

void operate(char op, double num[], int *numtop)
{
	switch (op)
	{
	case '+':
		num[*numtop - 1] = num[*numtop - 1] + num[*numtop];
		*numtop = *numtop - 1;
		break;
	case '-':
		num[*numtop - 1] = num[*numtop - 1] - num[*numtop];
		*numtop = *numtop - 1;
		break;
	case '*':
		num[*numtop - 1] = num[*numtop - 1] * num[*numtop];
		*numtop = *numtop - 1;
		break;
	case '/':
		num[*numtop - 1] = num[*numtop - 1] / num[*numtop];
		if (num[*numtop] == 0) // 分母為0
		{
			divide = 1;
		}
		*numtop = *numtop - 1;
		break;
	default:
		break;
	}
}