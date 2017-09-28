# https://www.hackerrank.com/challenges/ctci-is-binary-search-tree/problem

""" Node is defined as
class node:
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None
"""
treeValues = []

def runTree(root):
    if root == None:
        return
    
    runTree(root.left)
    treeValues.append(root.data)
    runTree(root.right)


def checkBST(root):
    if root == None:
        return 1 # true

    runTree(root)
    for i in range(len(treeValues)):
        for j in range(i+1,len(treeValues)):
            if treeValues[i] >= treeValues[j]:
                return 0
    return 1