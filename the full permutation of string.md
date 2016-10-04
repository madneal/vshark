## the full permutation of string

### description

Input a string, print the full permutation of the string. For example, input the string 'abc', hence the full permutation of the string is 'abc', 'acb', 'bac', 'bca', 'cab', 'cba'.

### analysis and solution

### method 1 :recursive implementation

Select one character from the string to be the first character, and the make the full permutation of the other characters. And continue recursive operation, and obtain the full permutation of the string. Take the string 'abc' for example like the following steps.

* Let 'a' be the first one, and obtatin the permutation of 'bc', 'abc' and 'acb'.

* Let 'b' be the first one, and obtatin the permutation of 'ac', 'bac' and 'bca'.

* Let 'c' be the first one,  and obtain the permutation of 'bc', 'cba' and 'cab'

  ```java
  public static void calAllpremutation(char[] perm, int from, int to) {
          if (to <= 1) {
              return;
          }
          if (from == to) {
              for (int i = 0;i <= to; i++) {
                  System.out.print(perm[i]);
              }
              System.out.println();
          } else {
              for (int j = from; j <= to; j++) {
                  swap(perm, j, from);
                  calAllpremutation(perm, from + 1, to);
                  swap(perm, j, from);
              }
          }

      }

      public static void swap(char[] arr, int a, int b) {
          char temp = arr[a];
          arr[a] = arr[b];
          arr[b] = temp;
      }
  ```

  ​

### 

