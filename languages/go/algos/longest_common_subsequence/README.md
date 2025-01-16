# Longest Common Subsequence (LCS)

The Longest Common Subsequence (LCS) problem is a classic problem in computer science and mathematics. It involves finding the longest contiguous or non-contiguous substring that is common to two or more sequences.

## Formal Definition

  Given two sequences:

    X = (x1, x2, ..., xm)

    Y = (y1, y2, ..., yn)

  The LCS problem is to find the longest sequence Z = (z1, z2, ..., zk) such that:

    Z is a subsequence of both X and Y

    k is the length of the LCS


## Example
  Suppose we have two sequences:

    X = (A, G, G, T, A, B)

    Y = (G, X, T, X, A, Y, B)

The LCS of X and Y is (G, T, A, B), which has a length of 4.

## Applications
  The LCS problem has numerous applications in:
    - Data comparison: LCS is used to compare two or more sequences, such as DNA or protein sequences, to identify similarities and differences.
    - Text editing: LCS is used in text editing algorithms, such as diff and merge, to identify common subsequences between two versions of a text.
    - Bioinformatics: LCS is used in bioinformatics to compare and analyze biological sequences, such as DNA, RNA, and protein sequences.

## Algorithmic Solutions
There are several algorithmic solutions to the LCS problem, including:
- Dynamic programming: This approach uses a 2D array to store the lengths of common subsequences and has a time complexity of O(m * n).
- Recursive approach: This approach uses recursion to find the LCS and has a time complexity of O(2^(m+n)).


The dynamic programming approach is generally more efficient and is widely used in practice.
