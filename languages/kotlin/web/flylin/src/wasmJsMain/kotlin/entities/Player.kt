package entities

import org.w3c.dom.CanvasRenderingContext2D
import toJsAny

class Player(
    private var x: Double,
    private var y: Double,
    private val minimumX: Double,
    private val maximumX: Double,
    private var xSpeed: Double = 0.0,
    private var ySpeed: Double = 0.0,
    private var width: Double = 50.0,
    private var height: Double = 50.0,
) {
    private var bullets = mutableListOf<PlayerBullet>()

    fun render(context2d: CanvasRenderingContext2D) {

        for (b in bullets) {
            b.render(context2d)
        }

        context2d.strokeStyle = toJsAny("yellow")
        context2d.beginPath()
        context2d.moveTo(x + (width / 2), y - height / 2)
        context2d.lineTo(x + width, y + height / 2)
        context2d.lineTo(x, y + height / 2)
        context2d.lineTo(x + width / 2, y - height / 2)
        context2d.stroke()

        context2d.beginPath()
        context2d.strokeStyle = toJsAny("red")
        context2d.lineWidth = 8.0
        context2d.moveTo(x + (width / 2), y - (height / 2) - 5)
        context2d.lineTo(x + width / 2, y)
        context2d.stroke()
    }

    fun moveLeft() {
        xSpeed = -10.0
    }

    fun stopLeft() {
        xSpeed = 0.0
    }

    fun stopRight() {
        xSpeed = 0.0
    }

    fun moveRight() {
        xSpeed = 10.0
    }

    fun shoot() {
        bullets.add(PlayerBullet((x + (width / 2.0)), y - 2))
    }

    fun getBullets() = bullets

    fun update() {
        if (x >= minimumX && x <= (maximumX - width - 10)) {
            x += xSpeed
        } else {
            xSpeed = 0.0
        }

        if (x < minimumX) {
            x = minimumX
        }

        if (x > maximumX - width - 10) {
            x = maximumX - width - 10
        }

        for (b in bullets) {
            b.update()
        }

        bullets = bullets.filter { it.y > 5.0 }.toMutableList()
    }
}