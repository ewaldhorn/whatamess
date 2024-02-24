package entities

import org.w3c.dom.CanvasRenderingContext2D
import toJsAny

class PlayerBullet(
    private var x: Double,
    var y: Double,
    private val speed: Double = 10.0,
    private val damage: Int = 1
) {
    fun render(context2d: CanvasRenderingContext2D) {
        if (y > 1.0) {
            context2d.strokeStyle = toJsAny("red")
            context2d.strokeRect(x - 1, y, 3.0, 3.0)
            context2d.fillStyle = toJsAny("orange")
            context2d.fillRect(x - 1, y, 3.0, 3.0)
        }
    }

    fun update() {
        y -= speed
    }

    // Test if we hit something.  If we did, return the amount of damage and set the y to -1, so we can clear the bullet
    fun testForHit(xPos: Double, yPos: Double, width: Double, height: Double): Int {
        return if ((x >= (xPos - 3) && (x <= xPos + 3 + width)) && ((y >= (yPos - 3) && (y <= yPos + 3 + height)))) {
            y = -1.0
            damage
        } else {
            0
        }
    }
}