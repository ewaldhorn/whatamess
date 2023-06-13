package za.co.nofuss.components

import androidx.compose.runtime.Composable
import com.varabyte.kobweb.compose.foundation.layout.Arrangement
import com.varabyte.kobweb.compose.foundation.layout.Row
import com.varabyte.kobweb.compose.ui.Alignment
import com.varabyte.kobweb.compose.ui.Modifier
import com.varabyte.kobweb.compose.ui.modifiers.fillMaxWidth
import com.varabyte.kobweb.compose.ui.modifiers.margin
import com.varabyte.kobweb.silk.components.graphics.Image
import org.jetbrains.compose.web.css.percent
import org.jetbrains.compose.web.css.px
import za.co.nofuss.util.Resources

@Composable
fun Header() {
    Row(
        modifier = Modifier.fillMaxWidth(80.percent).margin(topBottom = 50.px),
        horizontalArrangement = Arrangement.SpaceBetween,
        verticalAlignment = Alignment.CenterVertically
    ) {
        LeftSide()
    }
}

@Composable
fun LeftSide() {
    Row {
        Image(
            src = Resources.Images.NOFUSS_LOGO,
            desc = "NoFuss Logo"
        )
    }
}

@Composable
fun RightSide() {

}