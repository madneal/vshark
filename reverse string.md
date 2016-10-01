## reverse string

### description

Given a string, it isrequired to move the several characters in the front to the tail of the string. For example, move the first three characters 'a', 'b', 'c' of the string 'abcdef' to the tail of the string. So the string will change to 'deface'. Please write a function to implement this.

#### method 1

At first, we may think of the method to move the character to the end one by one.  Given a string s, and the length of the string is n. So we can write a function `leftShiftOne(s, n)` to move one character to the end of the string.

```javascript
function leftShiftOne(s, n) {
    s = s.split('');
    var firstChar = s[0];
    for (var i = 1;i < n;i++) {
         s[i-1] = s[i];
    }
    s[n-1] = firstChar;
    return s.join('');
}
```

And then we call `leftShiftOne` m times, to make the first m characters to the end of the string.

```javascript
function leftRotateString(s, n, m) {
	while (m--) {
		s = leftShiftOne(s, n);
	}
	return s;
}
```

Hereby we can satify the requirement to move the first few characters to the end. Then we can analysis the time complexity and room complextity of the algorithm. For the string of n,  assume that to move m characters to the end, and m*n times operations should be finished. Meanwhile we need a variable to save the first character. Hence, the time complexity if O(mn), and the room complexity is O(1). Is there better way to reduce the time complexity?