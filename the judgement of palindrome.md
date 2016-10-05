## the judgement of a palindrome

### description

For the given string, how to judge if the string is a palindrome?

### analysis and solution

A palindrome is a word that is the same wheter you read it backward or forward. Just like 'madam', 'refer'. So how to judge if a string is a palindrome.

### method1: scan from the two sides to the middle

For the given string, define two pointers point to the head and end of the string. Then let the two pointers scan towards middle. During the process of scan, if the characters pointed to by the two pointers are same all. The string is a palindrome.

```javascript
function isPlaindrome(s) {
	if (s === '') {
		return false;
	}
	s = s.split('');
	var front = 0;
	var end = s.length - 1;
	while(front < end) {
		if (s[front] !== s[end]) {
			return false;
		}
		++ front;
		-- end;
	}
	return true;
}
```

The time complexity is O(n), and the room complexity is O(1).



### method2: scan from the middle to the two sides

Actually, besides the method1 , we can also let a pointer scan from the middle to the two sides and compare the corresponding characters.

```javascript
function isPalindrome(s) {
	if (s === '') {
		return false;
	}
	s = s.split('');
	var n = s.length;
	var m = n & 1 ? (n-1)/2 : n / 2 -1;
	var first = m;
	var second = n - 1 - m;
	while (first >= 0) {
		if (s[first] !== s[second]) {
			return false;
		}
		-- first;
		++ second;
	}
	return true;
}
```



