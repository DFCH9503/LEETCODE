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

let node3l1 = new ListNode(3)
let node2l1 = new ListNode(4, node3l1)
let node1l1 = new ListNode(2, node2l1)

let node3l2 = new ListNode(4)
let node2l2 = new ListNode(6, node3l2)
let node1l2 = new ListNode(5, node2l2)

let res = inorderTraversal(root)

console.log("answer to the first example is:", res.val, res.next.val, res.next.next.val)