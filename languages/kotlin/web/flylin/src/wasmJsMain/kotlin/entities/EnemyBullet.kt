package entities

import org.w3c.dom.CanvasRenderingContext2D
import toJsAny

class EnemyBullet(var x: Double, var y: Double, val speed: Double = 8.0, val damage: Int = 1) {
    fun render(context2d: CanvasRenderingContext2D) {
        context2d.strokeStyle = toJsAny("orange")
        context2d.strokeRect(x, y, 3.0, 3.0)
        context2d.fillStyle = toJsAny("green")
        context2d.fillRect(x, y, 3.0, 3.0)
    }

    fun update() {
        y += speed
    }

    // Test if we hit something.  If we did, return the amount of damage and set the y to -1, so we can clear the bullet
    fun testForHit(xPos: Double, yPos: Double, width: Double, height: Double): Int {
        return if ((x in xPos..xPos) && (y in yPos..yPos)) {
            y = 5000.0
            damage
        } else {
            0
        }
    }
}