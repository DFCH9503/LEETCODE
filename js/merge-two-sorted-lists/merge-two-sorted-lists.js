const util = require("util")

function ListNode(val, next) {
    this.val = (val===undefined ? 0 : val)
    this.next = (next===undefined ? null : next)
}

function mergeTwoLists(list1 , list2){

    let dummyNode = new ListNode(0)
    let tail = dummyNode

    while (list1 && list2){
        if (list1.val <= list2.val){
            tail.next = list1
            list1 = list1.next
        }else{
            tail.next = list2
            list2 = list2.next
        }
        tail = tail.next
    }
    tail.next = list1 || list2

    return dummyNode.next
}


let node3l1 = new ListNode(4)
let node2l1 = new ListNode(2, node3l1)
let node1l1 = new ListNode(1, node2l1)

let node3l2 = new ListNode(4)
let node2l2 = new ListNode(3, node3l2)
let node1l2 = new ListNode(1, node2l2)

let res = mergeTwoLists(node1l1, node1l2)

console.log("answer to the first example is:", res.val, res.next.val, res.next.next.val, res.next.next.next.val, res.next.next.next.next.val, res.next.next.next.next.next.val)

//this line allows to see the full object contained in res
console.log(util.inspect(res, false, null, true /* enable colors */))

