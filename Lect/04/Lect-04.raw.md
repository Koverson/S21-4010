m4_include(../../setup.m4)

Lecture 4 - Hashing and Mining
====================

News
------

1. Janet Yellen offense stance.


[Hash and Merkel Trees - https://youtu.be/QgeJIpGkChw 4010-L04-pt1-Hash-Markle-Tree.mp4](https://youtu.be/QgeJIpGkChw 4010-L04-pt1-Hash-Markle-Tree.mp4)<br>
[Hash Function - https://youtu.be/h2QImL_X7S4](https://youtu.be/h2QImL_X7S4)<br>
[Mining - https://youtu.be/a4xNkrA_7aY](https://youtu.be/a4xNkrA_7aY)<br>
[Application of Hasing - https://youtu.be/KsMtapQnhws](https://youtu.be/KsMtapQnhws)<br>
[Merkle Trees Overview - https://youtu.be/Pqx0kmicwrc](https://youtu.be/Pqx0kmicwrc)<br>
[Code Walk high - https://youtu.be/FanMRbGy46g](https://youtu.be/FanMRbGy46g)<br>

From Amazon S3 - for download (same as youtube videos)

[Hash Function](http://uw-s20-2015.s3.amazonaws.com/4010-L04-pt2-Hash-Function.mp4)<br>
[Mining](http://uw-s20-2015.s3.amazonaws.com/4010-L04-pt3-Mining.mp4)<br>
[Application of Hasing](http://uw-s20-2015.s3.amazonaws.com/4010-L04-pt4-Application-of-Hashs.mp4)<br>
[Merkle Trees Overview](http://uw-s20-2015.s3.amazonaws.com/4010-L04-pt5-Merkle-Trees.mp4)<br>
[Code Walk high](http://uw-s20-2015.s3.amazonaws.com/4010-L04-pt6-Hash-Code-Walkthrough.mp4)<br>

What is a Hash
====================================================

A hash is a mapping from an input set to a unique number.
This is so that we have a high probability of a unique number for each different input.
The output will have a fixed length.


There are a number of different kind of hashes.

1. md5 a deprecated hash - it has weaknesses and should not be used.
2. sha1 some weaknesses.
3. sha256 
3. sha512 
3. sha3

hashes are used to validate passwords.  pbkdf and Ashley Madison.

Some hashes are specificly designed to be slow for passwords.

Some are build to be in hardware (sha1, sha256) or to not be in
hardware (Keccak used in Etherum).
 
Bitcoin uses sha256.

Example:

| Input             | Output                                                            |
|-------------------|-------------------------------------------------------------------|
| Hi                | c01a4cfa25cb895cdd0bb25181ba9c1622e93895a6de6f533a7299f70d6b0cfb  |
| Bitcoin           | deb10ca6fd85a5eba792ea8561da390635242f0c37c376f8eb7d7859adbffca9  |
| war-and-peace.txt | 6cbcc5ca5e590fb9ace161a5b93e4ecf280d7118104be0d63b686c004cfa70ae  |


Hashes for our purposes are unidirectional.  You can take the same input and get the same output, but you can't derive the input from the output.


```
// HashStrings hash a set of strings and return in hex-strings form
func HashStrings(a ...string) string {
	h := sha256.New()
	for _, z := range a {
		h.Write([]byte(z))
	}
	return fmt.Sprintf("%x", (h.Sum(nil)))
}
```

One of your first Assignmetns was to build a 'ksum' that works like sha256sum or md5sum and reads a file and calculates the hash for that file.

Once you have a hash you can prove that the original file has not changed.

This is the basis of lots of Blockchain metadata applications.



Blockchain and Mining
====================================================


What is Mining and How is it implemented.
-----

1. More on Go

Maps do not synchronize automatically.
So... Synchronization Primitives:

```Go
m4_include(Lect-04_05.go)
```

### A Go Core/Panic 

First the Code

```Go
m4_include(Lect-04_00.go)
```

Then the bad output.


```
m4_include(Lect_04_02.txt)
```

### Pseudo Code for Mining (Homework 02)


```Go
m4_include(Lect-04_01.go)
```


