## the sort of odd and even

### description

For a given integer array, please adjust the sorting of the array so that all the odd numbers are in the front of the array, and the even numbers are in the end. The time complexity should be O(N).

### analysis and solution

The simplest solution to scan the array from the beginning to the end, when come across an even number, pick it up and move the number behind the even number forwards. Finally the even number will be put at the end of the array. As it needs to move O(n) numbers when coming across an even number. Hence the time complexity is O(n^2).

Actually, if the odd number can seen as small number, and the even number can be seen as big number. By the requirement of the problem, it means that small number should be in the front of the big number. Imagine the process of quick sorting, the array will be cut into two parts by the main role, the numbers less than the main role will be placed in the front, and the numbers larger than main role are placed at the end. And the particition process can be implemented by the following two methods.

1) The the head pointer and the end pointer scan towards the middle. If the number pointed by head pointer is larger than main role and the number pointed by end pointer is less than main role. Then swap the numbers pointed by the head pointer and the end pointer.

2) Let two pointers scan from the left to the right, and the two pointers are the front pointer and the behind pointer respectively. If the front pointer is less than the main role, then the behind pointer move towards right, and swap the numbers pointed by the two pointers.

Also, for this problem, we can take a similar method.

### method1: the front pointer and the end pointer scan to the middle

According to the above first thought, we may consider to maintain two pointers: one points to the first number of the array, called the front pointer and move it towards right; and the other pointer points to the final number of the array, called the end pointer, move it towards left.