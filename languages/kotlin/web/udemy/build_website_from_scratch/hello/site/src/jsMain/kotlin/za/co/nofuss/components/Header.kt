package za.co.nofuss.components

import androidx.compose.runtime.Composable
import com.varabyte.kobweb.compose.css.FontWeight
import com.varabyte.kobweb.compose.css.ObjectFit
import com.varabyte.kobweb.compose.css.TextDecorationLine
import com.varabyte.kobweb.compose.foundation.layout.Arrangement
import com.varabyte.kobweb.compose.foundation.layout.Row
import com.varabyte.kobweb.compose.ui.Alignment
import com.varabyte.kobweb.compose.ui.Modifier
import com.varabyte.kobweb.compose.ui.modifiers.*
import com.varabyte.kobweb.silk.components.graphics.Image
import com.varabyte.kobweb.silk.components.navigation.Link
import com.varabyte.kobweb.silk.components.style.toModifier
import org.jetbrains.compose.web.css.percent
import org.jetbrains.compose.web.css.px
import za.co.nofuss.models.Section
import za.co.nofuss.models.Theme
import za.co.nofuss.styles.NavigationItemStyle
import za.co.nofuss.util.Constants
import za.co.nofuss.util.Resources

@Composable
fun Header() {
    Row(
        modifier = Modifier.fillMaxWidth(80.percent).margin(topBottom = 50.px),
        horizontalArrangement = Arrangement.SpaceBetween,
        verticalAlignment = Alignment.CenterVertically
    ) {
        LeftSide()
        RightSide()
    }
}

@Composable
fun LeftSide() {
    Row {
        Image(
            modifier = Modifier.fillMaxSize().objectFit(ObjectFit.Cover),
            src = Resources.Images.NOFUSS_LOGO,
            desc = "NoFuss Logo"
        )
    }
}

@Composable
fun RightSide() {
    Row(
        modifier = Modifier.fillMaxWidth(50.percent).borderRadius(r = 50.px).backgroundColor(Theme.LightGray.rgb)
            .padding(all = 20.px),
        horizontalArrangement = Arrangement.End
    ) {
        Section.values().take(7).forEach { section ->
            Link(
                modifier = NavigationItemStyle.toModifier()
                    .padding(right = 30.px)
                    .fontFamily(Constants.FONT_FAMILY)
                    .fontSize(18.px)
                    .fontWeight(FontWeight.Normal)
                    .textDecorationLine(TextDecorationLine.None),
                path = section.path,
                text = section.title
            )
        }
    }
}