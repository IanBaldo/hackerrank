# https://www.hackerrank.com/challenges/ctci-is-binary-search-tree/problem

""" Node is defined as
class node:
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None
"""
# Checks if Left tree if a BST
def deepL(node, val):
    if node == None:
        return 1 # true

    if (node.left and (node.left.data >= node.data )) or (node.right and (node.right.data <= node.data or node.right.data >= val)):
        return 0 # false
    
    return deepL(node.left, node.data) and deepR(node.right, val)

# Checks if Right tree is BST
def deepR(node, val):
    if node == None:
        return 1 # true
    
    if (node.left and (node.left.data >= node.data or node.left.data <= val )) or (node.right and (node.right.data <= node.data)):
        return 0 # false
    
    return deepL(node.left, val) and deepR(node.right, node.data)

def checkBST(root):
    if root == None:
        return 1 # true
    return deepL(root.left,root.data) and deepR(root.right,root.data)