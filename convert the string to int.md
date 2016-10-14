## convert the string to int

### description

Input a string consists of numbers, please convert this to an int and output. For example, input the string '123' and output 123.

Actually, this will be easily implemented by javascript. As it has the feature to convert string to number. Hence, we will dicuss the solution by java.

### analysis and solution

Take '123' for example to explain the converting process:

1. When scan the first character '1', as it is the first one, so obtain the integer 1.
2. When scan the second character '2', as we have 1 before (which means 10), so obtatin 1*10+2=12.
3. Continue to scan the third character '3', as we have 12 before (which means 120), so obtain 12*10 + 3 = 123.

Hence, the idea of this solution is : scan the character from left to right, and multiply the before result with 10 and add the number of the current character.

```java
    public static int strToInt(char[] str) {
        int n = 0;
        for (int i = 0, len = str.length; i < len; i++) {
            int c = str[i] - '0';
            n = n*10 + c;
        }
        return n;
    }
```

But the above code overlooks several details:

* You'd better to judge if the input is empty.
* If the first character is '-', the result should be negative.
* The string should contains only the character of number.
* The input string cannot be too long, or the result will integer overflow.
