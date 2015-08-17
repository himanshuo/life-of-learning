/*
 * 1) convert input into zigzag.
 * zigzag always takes the form of taking the input and moving it
 * down the column for nRows  spaces. Then, cross it to the next to next
 * column. In the cross, you will end up hitting the center columns.
 * For nRows=3, you hit only the middle row in the crossing
 * for nRows=4, you hit 
 * 
 * 
 * A A
 * AAA
 * AAA
 * A A

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:

string convert(string text, int nRows);
convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR"
 */ 
package leetcode
