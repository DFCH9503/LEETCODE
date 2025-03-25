function ListNode(val, next) {
    this.val = (val===undefined ? 0 : val)
    this.next = (next===undefined ? null : next)
}

function addTwoNumbers(l1, l2){
    let dummy = new ListNode()
    let temp = dummy
    let carry = 0
    
    while (l1 !== null || l2 !== null || carry !== 0) {
        let val1 = l1 !== null ? l1.val : 0
        let val2 = l2 !== null ? l2.val : 0
        
        let sum = val1 + val2 + carry
        carry = Math.floor(sum / 10)
        temp.next = new ListNode(sum % 10)
        temp = temp.next
        
        if (l1 !== null) l1 = l1.next
        if (l2 !== null) l2 = l2.next
    }
    
    return dummy.next
}

let node3l1 = new ListNode(3)
let node2l1 = new ListNode(4, node3l1)
let node1l1 = new ListNode(2, node2l1)

let node3l2 = new ListNode(4)
let node2l2 = new ListNode(6, node3l2)
let node1l2 = new ListNode(5, node2l2)

let res = addTwoNumbers(node1l1, node1l2)

console.log("answer to the first example is:", res.val, res.next.val, res.next.next.val)
