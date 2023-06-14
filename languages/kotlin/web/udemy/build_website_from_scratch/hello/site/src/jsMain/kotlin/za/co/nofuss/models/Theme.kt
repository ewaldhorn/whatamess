package za.co.nofuss.models

import com.varabyte.kobweb.compose.ui.graphics.Color.Companion.rgb
import org.jetbrains.compose.web.css.CSSColorValue

enum class Theme(val hex: String, val rgb: CSSColorValue) {
    Primary(hex = "#1d9c09", rgb = rgb(r = 29, g = 156, b = 9)),
    Secondary(hex = "121D34", rgb = rgb(r = 18, g = 29, b = 52)),
    DarkGray(hex = "#808080", rgb = rgb(r = 128, g = 128, b = 128)),
    Gray(hex = "#CFCFCF", rgb = rgb(r = 207, g = 207, b = 207)),
    LightGray(hex = "#EDEDED", rgb = rgb(r = 237, g = 237, b = 237)),
    LighterGray(hex = "#F9F9F9", rgb = rgb(r = 249, g = 249, b = 249))
}