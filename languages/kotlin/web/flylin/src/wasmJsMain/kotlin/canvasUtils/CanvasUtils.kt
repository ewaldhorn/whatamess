package canvasUtils

import kotlinx.browser.document
import kotlinx.browser.window
import org.w3c.dom.BEVEL
import org.w3c.dom.CanvasLineJoin
import org.w3c.dom.CanvasRenderingContext2D
import org.w3c.dom.HTMLCanvasElement
import toJsAny

fun configureCanvas(): CanvasRenderingContext2D {
    val canvas = document.getElementById("canvas") as HTMLCanvasElement
    val context = canvas.getContext("2d") as CanvasRenderingContext2D

    canvas.width = window.innerWidth - 20
    canvas.height = window.innerHeight - 80

    return context
}

fun configureCommonStyle(context2D: CanvasRenderingContext2D) {
    context2D.shadowColor = "#953"
    context2D.shadowBlur = 20.0
    context2D.lineJoin = CanvasLineJoin.BEVEL
    context2D.lineWidth = 5.0
}