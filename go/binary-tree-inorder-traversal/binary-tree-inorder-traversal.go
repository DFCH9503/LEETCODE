package main

import(
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func inorderTraversal (root *TreeNode) []int{
	ret := []int{}

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode){
        if root == nil {
            return 
        }
        dfs(root.Left)
        ret = append(ret, root.Val)
        dfs(root.Right)
    }
    dfs(root)
    return ret
}

func main(){
	res := inorderTraversal()

	fmt.Println("answer to the first example is:", )
}