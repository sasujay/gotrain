var numSquares = 6
var colors = generateRandomColors(numSquares)
var squares = document.querySelectorAll(".square")
var pickedColor = pickColor()
var colorDisplay = document.getElementById("rgb")
var message = document.querySelector("#message")
var replay = document.querySelector("#replay")
var h1 = document.querySelector("h1")
var hard = document.getElementById("hard")
var easy = document.getElementById("easy")

replay.addEventListener("click", function () {
    colors = generateRandomColors(numSquares)
    pickedColor = pickColor()
    colorDisplay.textContent = pickedColor;
    this.textContent = "New Colors"
    message.textContent = ""
    for (var i = 0; i < squares.length; i++) {
        squares[i].style.background = colors[i]
    }
    h1.style.background = "steelblue"
})
easy.addEventListener("click", function () {
    hard.classList.remove("on")
    easy.classList.add("on")
    numSquares = 3
    colors = generateRandomColors(numSquares)
    pickedColor = pickColor()
    colorDisplay.textContent = pickedColor;
    for (var i = 0; i < squares.length; i++) {
        if (colors[i]) {
            squares[i].style.background = colors[i]
        } else {
            squares[i].style.display = "none"
        }
    }

})
hard.addEventListener("click", function () {
    easy.classList.remove("on")
    hard.classList.add("on")
    numSquares = 6
    colors = generateRandomColors(numSquares)
    pickedColor = pickColor()
    colorDisplay.textContent = pickedColor;
    for (var i = 0; i < squares.length; i++) {
        squares[i].style.background = colors[i]
        squares[i].style.display = "block"
    }
})

colorDisplay.textContent = pickedColor;

for (var i = 0; i < squares.length; i++) {
    //add initial colors to squares
    squares[i].style.background = colors[i]
    //add click listerners to squres
    squares[i].addEventListener("click", function () {
        //grab color or clicked square
        var clickedColor = this.style.background;
        if (this.style.background == pickedColor) {
            message.textContent = "Correct!!"
            resetColor(clickedColor)
            h1.style.background = clickedColor
            replay.textContent = "PLAY AGAIN? "
        } else {
            this.style.background = "#232322"
            message.textContent = "Try Again"
        }
    })
}

function resetColor(color) {
    for (var i = 0; i < squares.length; i++) {
        squares[i].style.background = color
    }
}
function pickColor() {
    var random = Math.floor(Math.random() * colors.length)
    console.log(random)
    return colors[random]
}
function generateRandomColors(num) {
    var arr = []
    for (var i = 0; i < num; i++) {
        arr.push(randomColor())
    }
    return arr
}
function randomColor() {
    //Pick a Red 0 - 255
    var r = Math.floor(Math.random() * 256)
    //Pick a Blue 0 - 255
    var b = Math.floor(Math.random() * 256)
    //Pick a Gree 0 - 255
    var g = Math.floor(Math.random() * 256)
    //rbg(r, b, g)
    return "rgb(" + r + ", " + b + ", " + g + ")"

}
function resetGame() {
    for (var j = 0; j < squares.length; j++) {
        //add initial colors to squares
        squares[i].style.background = colors[j]
    }
}