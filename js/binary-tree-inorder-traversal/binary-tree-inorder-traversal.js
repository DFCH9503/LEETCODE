function TreeNode(val, left, right) {
    this.val = (val===undefined ? 0 : val)
    this.left = (left===undefined ? null : left)
    this.right = (right===undefined ? null : right)
}

function inorderTraversal(root){

    const res = []

    function inorder(node) {
        if (!node) {
            return
        }
        inorder(node.left)
        res.push(node.val)
        inorder(node.right)
    }

    inorder(root)
    return res 

}

let root = new TreeNode(1)
let nodeA = new TreeNode(null)
let nodeB = new TreeNode(2)
let nodeC = new TreeNode(3)

root.left = nodeA
root.right = nodeB
nodeB.left = nodeC

let res = inorderTraversal(root)
res.shift()  //remove the element before root

console.log("answer to the first example is:", res)