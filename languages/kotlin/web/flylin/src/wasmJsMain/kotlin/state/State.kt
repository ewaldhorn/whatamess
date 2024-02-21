package state

import canvasUtils.configureCommonStyle
import entities.Enemy
import entities.Player
import kotlinx.browser.document
import kotlinx.browser.window
import org.w3c.dom.BEVEL
import org.w3c.dom.CanvasLineJoin
import org.w3c.dom.CanvasRenderingContext2D
import org.w3c.dom.events.Event
import org.w3c.dom.events.KeyboardEvent
import org.w3c.dom.events.MouseEvent
import toJsAny
import kotlin.random.Random

enum class Phase {
    WAITING, PLAYING, REPLAY
}

class State {
    private var phase = Phase.WAITING
    lateinit var player: Player
    private var maxWidth: Double = 0.0
    private var maxHeight: Double = 0.0
    private var score = 0
    private var oldScore = 0
    private var enemies = mutableListOf<Enemy>()
    private var mouseStatus = false

    fun configureMaxWidthAndHeight(context2D: CanvasRenderingContext2D) {
        maxWidth = context2D.canvas.width.toDouble()
        maxHeight = context2D.canvas.height.toDouble()
    }

    fun getMaxWidth() = maxWidth
    fun getMaxHeight() = maxHeight

    fun getScore() = score
    fun updateScore(value: Int) {
        score += value
    }

    fun configureItems() {
        player = Player(maxWidth / 2.0 - (25.0), maxHeight - 55, 25.0, maxWidth)

        for (i in 1..5) {
            enemies.add(Enemy(30.0 + Random.nextDouble(maxWidth / 2.0), i * 60.0, 25.0, maxWidth, maxHeight))
        }
    }

    fun updateItems() {
        player.update()

        for (bullet in player.getBullets()) {
            for (enemy in enemies) {
                val damage = bullet.testForHit(enemy.getX(), enemy.getY(), enemy.getWidth(), enemy.getHeight())
                if (damage > 0) {
                    enemy.takeHit(damage)
                    updateScore(damage * 4)
                }
            }
        }

        for (enemy in enemies) {
            enemy.update()
        }

        enemies = enemies.filter { it.getHealth() > 0 }.toMutableList()

        if (enemies.isEmpty()) {
            oldScore = score
            score = 0
            phase = Phase.REPLAY
            configureItems()
        }
    }

    fun renderItems(context2D: CanvasRenderingContext2D) {
        if (phase == Phase.PLAYING) {

            configureCommonStyle(context2D)

            context2D.shadowColor = "#953"
            for (enemy in enemies) {
                enemy.render(context2D)
            }

            context2D.shadowColor = "yellow"
            player.render(context2D)
        } else {
            renderHelpScreen(context2D)
        }
    }

    private fun renderHelpScreen(context2d: CanvasRenderingContext2D) {
        context2d.shadowColor = "black"
        context2d.shadowBlur = 0.0
        context2d.lineJoin = CanvasLineJoin.BEVEL
        context2d.lineWidth = 2.0

        context2d.font = "normal 26pt Arial"
        val text = "Use arrows to move and space to shoot. Press ENTER to begin!"
        val tmp = context2d.measureText(text).width
        context2d.strokeStyle = toJsAny("white")
        context2d.strokeText(text, (maxWidth / 2.0) - (tmp / 2.0), (maxHeight / 2.0) + 30.0)

        if (phase == Phase.REPLAY) {
            context2d.font = "normal 22pt Arial"
            val scoreText = "You scored $oldScore points!"
            val scoreWidth = context2d.measureText(scoreText).width
            context2d.strokeStyle = toJsAny("white")
            context2d.strokeText(scoreText, (maxWidth / 2.0) - (scoreWidth / 2.0), (maxHeight / 2.0) + 70.0)
        }

    }

    fun configureKeyboardCallbacks() {
        document.addEventListener("keydown", fun(e: Event) {
            val tmp = e as KeyboardEvent

            when (tmp.code) {
//                "ArrowUp" -> println("Arrow UP")
//                "ArrowDown" -> println("Arrow DOWN")
                "ArrowLeft" -> player.moveLeft()
                "ArrowRight" -> player.moveRight()
            }
        })
        document.addEventListener("keyup", fun(e: Event) {
            val tmp = e as KeyboardEvent

            when (tmp.code) {
//                "ArrowUp" -> println("Arrow UP")
//                "ArrowDown" -> println("Arrow DOWN")
                "ArrowLeft" -> player.stopLeft()
                "ArrowRight" -> player.stopRight()
                "Space" -> {
                    if (phase == Phase.PLAYING) {
                        player.shoot()
                        updateScore(-1)
                    }
                }

                "Enter" -> {
                    phase = Phase.PLAYING
                }
            }
        })
    }

    fun configureMouseCallbacks() {
        window.addEventListener("mousedown", fun(e: Event) {
            val tmp = e as MouseEvent
            if (tmp.button == 0.toShort()) {
                mouseStatus = true
            }
            e.preventDefault()
        })

        window.addEventListener("mouseup", fun(e: Event) {
            mouseStatus = false
            e.preventDefault()
        })

    }
}