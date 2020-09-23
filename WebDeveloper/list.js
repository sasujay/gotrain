let choice = prompt("Enter your choice:")
var mylist = []
while (choice != "quit") {
    if (choice == "new") {
        var myitem
        while (myitem != "done") {
            myitem = prompt("Enter item:")
            mylist.push(myitem)
        }
    } else if (choice == "list") {
        mylist.forEach(function (item) {
            console.log("*************")
            console.log(item)
        })
    } else if (choice == "delete") {
        var delitem
        index = prompt("Enter index item to delete:")
        mylist.splice(index, 1)
    }
    choice = prompt("Enter your choice:")
}
console.log("You have quit the App")