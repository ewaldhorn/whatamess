import js

def button_clicked():
    js.console.log("Button clicked!")

def main():
    btn = js.document.getElementById("myButton")
    btn.addEventListener("click", button_clicked)

main()
