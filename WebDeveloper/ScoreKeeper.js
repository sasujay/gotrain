
var p1Button = document.getElementById("p1")
var p2Button = document.getElementById("p2")
var myButtons = document.querySelectorAll("button")
var resetButton = document.getElementById("reset")
var p1Score = 0
var p2Score = 0
var p1 = document.querySelector("#p1score")
var p2 = document.querySelector("#p2score")
var gameOver = false;
var winningScore = 5;
var numInput = document.querySelector("input")
var winningScoreDisplay = document.querySelector("p span")
console.log(numInput)
winningscore = numInput
console.log(winningscore)
//while (p1Score < 6) {
p1Button.addEventListener("click", function () {
    if (!gameOver) {
        p1Score++
        console.log(p1Score)
        p1.textContent = p1Score
        if (p1Score === winningScore) {
            p1.classList.add("winner")
            gameOver = true
        }
    }
})
//while (p1Score < 6) {
p2Button.addEventListener("click", function () {
    if (!gameOver) {
        p2Score++
        console.log(p2Score)
        p2.textContent = p2Score
    }
    if (p2Score === winningScore) {
        p2.classList.add("winner")
        gameOver = true
    }

})

resetButton.addEventListener("click", function () {
    reset()
})
function reset() {
    p1Score = 0
    p1.textContent = p1Score
    p1.classList.remove("winner")
    p2Score = 0
    p2.textContent = p2Score
    p2.classList.remove("winner")
    gameOver = false
}

numInput.addEventListener("change", function () {
    winningScoreDisplay.textContent = this.value
    winningScore = Number(this.value)
    reset()
})
console.log(myButtons)
for (var i = 0; i < myButtons.length; i++) {
    myButtons[i].addEventListener("mouseover", function () {
        this.style.color = "green"
    })
    myButtons[i].addEventListener("mouseout", function () {
        this.style.color = "black"
    })
}
