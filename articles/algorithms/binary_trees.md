Binary Trees
======
A binary tree is when you have a node that can have 0, 1, or 2 child nodes.  

### Binary Tree vs. Binary Search Tree
Binary tree: each node has max of 2 children   

Binary Search Tree: left nodes <= cur < right nodes, still max of 2 children

### balanced vs. unbalanced
balanced: depth of subtrees do not differ by more than one

* should know avg and worst case time of algorithm

### Full and Complete
full: a node has either 0 or exactly 2 children  

![full binary tree](binary-trees-images/Full_binary.jpg)

complete: all levels of tree are completed, expect possible the last. The last level must have nodes as far to the left as possible.

![complete binary tree](binary-trees-images/Complete_binary.jpg)

For a tree to be full and complete, it must have `2^n - 1` nodes
