import canvasUtils.configureCanvas
import kotlinx.browser.document
import kotlinx.browser.window
import org.w3c.dom.*
import state.State

var textWidth = 0.0

var applicationState = State()
var scoreElement = document.getElementById("score") as HTMLDivElement
var restartButton = document.getElementById("restart") as HTMLButtonElement

fun main() {
    val context = configureCanvas()
    val metrics = context.measureText("Text")
    textWidth = metrics.width + 5
    applicationState.configureMaxWidthAndHeight(context)
    applicationState.configureItems()
    applicationState.configureKeyboardCallbacks()
    applicationState.configureMouseCallbacks()

    renderLoop(context)
}

fun resetGame(context2D: CanvasRenderingContext2D) {
    applicationState.updateScore(-1 * applicationState.getScore())
    renderLoop(context2D)
}

fun renderLoop(context: CanvasRenderingContext2D) {
    window.requestAnimationFrame {
        context.clearRect(0.0, 0.0, applicationState.getMaxWidth(), applicationState.getMaxHeight())
        applicationState.updateItems()
        applicationState.renderItems(context)
        scoreElement.innerText = "Score: ${applicationState.getScore()}"

        renderLoop(context)
    }
}

// https://discuss.kotlinlang.org/t/setting-fillstyle-on-canvas-using-new-wasm-issue/26477
// Workaround for setting fillstyle etc.
@JsFun("value => value")
external fun toJsAny(value: String): JsAny
