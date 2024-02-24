package entities

import org.w3c.dom.CanvasRenderingContext2D
import toJsAny
import kotlin.random.Random

class Enemy(
    private var x: Double,
    private var y: Double,
    private val minimumX: Double,
    private val maximumX: Double,
    private val maximumY: Double,
    private var damaged: Int = 0,
    private var health: Int = 10,
    private val maxHealth: Int = 10,
    private var xSpeed: Double = 0.0,
    private var ySpeed: Double = 0.0,
    private var width: Double = 50.0,
    private var height: Double = 50.0,
) {
    private var bullets = mutableListOf<EnemyBullet>()

    init {
        val tmp = Random.nextDouble(1.0, 3.0)
        xSpeed = if (Random.nextBoolean()) tmp else (tmp * -1.0)
    }

    fun render(context2d: CanvasRenderingContext2D) {
        for (b in bullets) {
            b.render(context2d)
        }

        if (damaged > 0) {
            context2d.strokeStyle = toJsAny("red")
            damaged -= 1
        } else {
            context2d.strokeStyle = toJsAny("green")
        }
        context2d.strokeRect(x, y, width, height)
        context2d.fillStyle = toJsAny("black")
        context2d.fillRect(x, y, width, height)

        context2d.font = "normal 16pt Arial"
        val tmp = context2d.measureText("$health").width
        context2d.strokeStyle = toJsAny("white")
        context2d.strokeText("$health", x + 25.0 - (tmp / 2.0), y + 30.0)
    }

    fun shoot() {
        bullets.add(EnemyBullet((x + (width / 2.0)), y + 2))
    }

    fun update() {
        if (x <= minimumX + 5 || x >= (maximumX - width - 10)) {
            xSpeed *= -1.0
            if (health < maxHealth) {
                health += 1
            }
        }

        x += xSpeed

        for (b in bullets) {
            b.update()
        }

        bullets = bullets.filter { it.y >= maximumY }.toMutableList()
    }

    fun getHealth() = health
    fun getX() = x
    fun getY() = y
    fun getWidth() = width
    fun getHeight() = height
    fun takeHit(damage: Int) {
        damaged += 10;
        health -= damage
    }
}