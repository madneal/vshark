##  the including of string

### description

Given a long string a and a short string b. So, how to judge the a contains the short string b totallly. Please write a function `stringContain(a, b)` to implement this.

For simplicity,we assume that the string only contains capital characters. For example,

* Assume a is 'ABCD', b is 'BAD', the answer is true, as a contains all the characters of b, or we can say b is the true subset of a.
* Assume a is 'ABCD', and  b is 'BCE', the answer is false, as the character 'E' is not in the a.
* Assume a is 'ABCD', and b is 'AA', and the answer is true, because the character 'A' is in b.


### method 1 poll violently

To judge the characters of b is contained in the long string a, the intuitive method is : judge each character in b, and compare it with the each characters in a to see the result.

```javascript
function stringContain(a, b) {
  var aLen = a.length;
  var bLen = b.length;
  a = a.split('');
  b = b.split('');
  for (var i = 0; i < bLen; ++ i) {
      for (var j = 0;j < aLen && a[j] !== b[j]; ++ j)
      ;
      if (j >= aLen) {
        return false;
      }
  }
  return true;
}
```

If the length of long string a is n, and the length of short string b is m. Hence, the algorithm will need O(nm) times comparison. Obviously, if n and m is large, the time cost will be huge, a better method is preferred.



### method 2 poll after sorted 

If sorting is allowed, we can poll after sorting. For example, we can sort the characters in the two strings, the roll in the two strings respectively.

Generally, the sorting of the two strings need O(mlogm) + O(nlogn) times operations, and the later linear scan need O(m+n) times operations. For the sort method, can adopt quick sort.

```javascript
function stringContain(a, b) {
	var aLen = a.length;
	var bLen = b.length;
	a = a.split('');
	b = b.split('');
	a.sort();
	b.sort();
	for (var pa = 0, pb = 0; b < bLen;) {
		while (pa < aLen && a[pa] < a[pb]) {
			++ pa;
		}
		if (pa > aLen || a[pa] > b[pb]) {
			return false;
		}
		++ pb;
	}
	return false;
}
```



There are two another methods in the book. But they seem uninteresting to me. So I will not mention here.

