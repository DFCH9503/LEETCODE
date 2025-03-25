function ListNode(val, next) {
    this.val = (val===undefined ? 0 : val)
    this.next = (next===undefined ? null : next)
}

function deleteDuplicates(head){
    let res = head

    while (head && head.next) {
        if (head.val === head.next.val) {
            head.next = head.next.next
        } else {
            head = head.next
        }
    }

    return res 
}

let node3Head = new ListNode(2)
let node2Head = new ListNode(1, node3Head)
let node1Head = new ListNode(1, node2Head)


let res = deleteDuplicates(node1Head)

console.log("answer to the first example is:", res)