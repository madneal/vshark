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



### method 2

Address this problem, we may think this problem in another side. Hence the problem requires that to move the fornt part characters to the end, we can handle the string requires movement and these don't respectively. For example,we can cut the string into two parts, then reverse the two parts respectively, finally reverse the whole string, and the problem will be solves.

Take the example of string 'abcdef', if we need to move 'def' to the front of 'abc', we only need three steps.

1. Cut the original string into two parts, X means 'abc', and Y means 'def'.
2. Reverse the characters of X which means reverse 'abc' to get 'cba'; then reverse the characters of Y, which will obtain 'fed' from 'def'.
3. Finally, reverse the result above, that is to reverse the whole string "cbafed" to obtain 'defabc'. Hereby, the reverse string is implemented.

```javascript
function reverseString(s, from, to) {
	s = s.split('');
	while (from < to) {
		var t = s[from];
		s[from++] = s[to];
		s[to--] = t;
	}
	return s.join('');
}

function leftRotateString(s, n, m) {
	m %= n;
	s = reverseString(s, 0, m-1);
	s = reverseString(s, m, n-1);
	s = reverseString(s, 0, n-1);
	return s;
}
```

Hereby, the time complexity is O(n), and the room complexity is O(1).