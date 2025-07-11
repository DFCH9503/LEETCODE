function ListNode(val, next) {
    this.val = (val===undefined ? 0 : val)
    this.next = (next===undefined ? null : next)
}

function swapPairs(head){
    let dummy = new ListNode(0, head)
    let prev = dummy, cur = head

    while (cur && cur.next) {
        let npn = cur.next.next
        let second = cur.next

        second.next = cur
        cur.next = npn
        prev.next = second

        prev = cur
        cur = npn
    }

    return dummy.next
}

let head = new ListNode(1)
let nodeA = new ListNode(2)
let nodeB = new ListNode(3)
let nodeC = new ListNode(4)

head.next = nodeA
nodeA.next = nodeB
nodeB.next = nodeC

res = swapPairs(head)

console.log(res)