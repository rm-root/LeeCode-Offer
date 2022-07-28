package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) SetValue(val int) {
	t.Val = val
}

func CreateTreeNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil}
}

//func buildTree(preorder []int, inorder []int) *TreeNode {
//	if len(preorder) == 0 {
//		return nil
//	}
//	root := &TreeNode{preorder[0], nil, nil}
//	i := 0
//	for ; i < len(inorder); i++ {
//		if inorder[i] == preorder[0] {
//			break
//		}
//	}
//	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
//	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
//	return root
//}

func buildTree(preorder []int, inorder []int) *TreeNode {
	for i := range inorder {
		if inorder[i] == preorder[0] { //中序遍历 root (index=k)
			return &TreeNode{
				Val: preorder[0],
				//Val: inorder[k],
				Left:  buildTree(preorder[1:i+1], inorder[0:i]),
				Right: buildTree(preorder[i+1:], inorder[i+1:]),
			}
		}
	}
	return nil
}

//func main() {
//	preorder := []int{3, 9, 20, 15, 7}
//	inorder := []int{9, 3, 15, 20, 7}
//	buildTree(preorder, inorder)
//}
