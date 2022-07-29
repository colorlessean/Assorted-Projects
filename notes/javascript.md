# JavaScript Notes
Collection of Language Notes for future reference.

## Basics
On page script
```
<script type="text/javascript"> ... </script>
```

Include external JS file
```
<script src="filename.js"></script>
```

Delay function
```
setTimeout(function () {}, 1000);
```

Functions
```
function addNumbers(a, b) {
    return a + b;
}
x = addNumbers(1, 2);
```

Edit a DOM element
```
document.getElementById("elementID").innerHTML = "Text!";
```

Output
```
console.log(a);             // browser console
document.write(a);          // write in HTML
alert(a);                   // output an alert box in browser
confirm("Confirmation?");   // yes/no dialog box returns true/false
prompt("Question?", "0");   // input dialog, second value is initial value
```

## Variables
```
var a;                          // variable
var b = "init";                 // string
var c = "Hi" + " " + "Joe";     // concat
var d = 1 + 2 + 3;              // addition opperand
var e = [2,3,5,8];              // array
var f = false;                  // boolean value
var g = /()/;                   // regex
var h = function() {}           // function object
const PI = 3.14;                // contant
var a = 1, b = 2, c = a + b;    // oneline initalization
let z = 'zzz';                  // block scope local variable
```

Strict mode to write secure code
```
"use strict";
x = 1;
```

Values
```
false, true                     // boolean values
18, 3.14, 0b10001, 0xF6, NaN    // numbers
"flower", 'John'                // strings
undefined, null, Infinity       // special
```

Operators
```
a = b + c - d;      // addition, subtraction
a = b * (c / d);    // mulitplication, division
x = 100 % 48;       // modulo i.e. remainder
a++; b--;           // postfix increment/decrement
```

Bitwise Operators
```
&   // AND
|   // OR
~   // NOT
^   // XOR
<<  // bitshift left
>>  // bitshift right
>>> // zero fill rightshift
```

Aritmatic

