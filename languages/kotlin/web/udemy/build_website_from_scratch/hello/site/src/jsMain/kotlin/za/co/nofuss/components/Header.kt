package za.co.nofuss.components

import androidx.compose.runtime.Composable
import com.varabyte.kobweb.compose.css.FontWeight
import com.varabyte.kobweb.compose.css.ObjectFit
import com.varabyte.kobweb.compose.css.TextDecorationLine
import com.varabyte.kobweb.compose.css.VerticalAlign
import com.varabyte.kobweb.compose.foundation.layout.Arrangement
import com.varabyte.kobweb.compose.foundation.layout.Row
import com.varabyte.kobweb.compose.ui.Alignment
import com.varabyte.kobweb.compose.ui.Modifier
import com.varabyte.kobweb.compose.ui.modifiers.*
import com.varabyte.kobweb.silk.components.graphics.Image
import com.varabyte.kobweb.silk.components.icons.fa.FaBars
import com.varabyte.kobweb.silk.components.icons.fa.IconSize
import com.varabyte.kobweb.silk.components.navigation.Link
import com.varabyte.kobweb.silk.components.style.breakpoint.Breakpoint
import com.varabyte.kobweb.silk.components.style.toModifier
import com.varabyte.kobweb.silk.theme.breakpoint.rememberBreakpoint
import org.jetbrains.compose.web.css.percent
import org.jetbrains.compose.web.css.px
import za.co.nofuss.models.Section
import za.co.nofuss.models.Theme
import za.co.nofuss.styles.LogoStyle
import za.co.nofuss.styles.NavigationItemStyle
import za.co.nofuss.util.Constants
import za.co.nofuss.util.Resources

@Composable
fun Header() {
    val breakpoint = rememberBreakpoint()
    Row(
        modifier = Modifier.fillMaxWidth(if (breakpoint > Breakpoint.MD) 80.percent else 90.percent)
            .margin(topBottom = 50.px),
        horizontalArrangement = Arrangement.SpaceBetween,
        verticalAlignment = Alignment.CenterVertically
    ) {
        LeftSide(breakpoint)
        if (breakpoint > Breakpoint.SM) {
            RightSide(breakpoint)
        }
    }
}

@Composable
fun LeftSide(breakpoint: Breakpoint) {
    Row(
        verticalAlignment = Alignment.CenterVertically
    ) {
        if (breakpoint <= Breakpoint.SM) {
            FaBars(
                modifier = Modifier.margin(right = 15.px),
                size = IconSize.XL
            )
        }
        Image(
            modifier = LogoStyle.toModifier().fillMaxSize(percent = 70.percent).objectFit(ObjectFit.Cover),
            src = Resources.Images.NOFUSS_LOGO,
            desc = "NoFuss Logo"
        )
    }
}

@Composable
fun RightSide(breakpoint: Breakpoint) {
    Row(
        modifier = Modifier.fillMaxWidth(if (breakpoint > Breakpoint.MD) 50.percent else 70.percent)
            .borderRadius(r = 50.px).backgroundColor(Theme.LightGray.rgb)
            .padding(all = 20.px),
        horizontalArrangement = Arrangement.Center
    ) {
        Section.values().take(5).forEach { section ->
            Link(
                modifier = NavigationItemStyle.toModifier()
                    .padding(right = 25.px)
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