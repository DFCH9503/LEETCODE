function ListNode(val, next) {
    this.val = (val===undefined ? 0 : val)
    this.next = (next===undefined ? null : next)
}

function removeNthFromEnd(head, n){
    let slow = head, fast = head
    for (let i = 0; i < n; i++) fast = fast.next
    if (!fast) return head.next
    while (fast.next) {
        fast = fast.next
        slow = slow.next
    }
    slow.next = slow.next.next
    return head
}

let node5 = new ListNode(5)
let node4 = new ListNode(4, node5)
let node3 = new ListNode(3, node4)
let node2 = new ListNode(2, node3)
let node1 = new ListNode(1, node2)

let res = removeNthFromEnd(node1, 2)

console.log(`answer to the first example is: [${res.val}, ${res.next.val}, ${res.next.next.val}, ${res.next.next.next.val}]`)

