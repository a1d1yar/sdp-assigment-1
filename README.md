# sdp-assigment-1


## Overview

This Go-based text manipulation tool is designed to perform various modifications on a given input text file and save the modified version to an output file. It follows coding best practices, includes unit tests, and can be audited by peers.
---
## Supported Modifications

1. Replace (hex) with the decimal equivalent of the preceding hexadecimal number.
2. Replace (bin) with the decimal equivalent of the preceding binary number.
3. Convert (up) to uppercase for the preceding word.
4. Convert (low) to lowercase for the preceding word.
5. Convert (cap) to capitalized form for the preceding word.
6. If a number follows (low), (up), or (cap), apply the conversion to the specified number of preceding words.
7. Ensure proper spacing around punctuation marks (., ,, !, ?, :, and ;) and handle special cases like ... and !?.
8. Place single quotes (`' '`) around words or phrases as needed.
9. Automatically change 'a' to 'an' when followed by a vowel or 'h'.

This tool empowers users to enhance the quality and readability of their text by applying these modifications efficiently.



<br>
<br>

---

## Clone the repo 
```bash
git clone git@git.01.alem.school:almaratov/go-reloaded.git
```

<br>

<br>
<br>

---

## How to run 

```bash
go run . sample.txt result.txt
```