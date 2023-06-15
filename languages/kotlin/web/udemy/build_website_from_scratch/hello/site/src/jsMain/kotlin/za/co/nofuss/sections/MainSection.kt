package za.co.nofuss.sections

import androidx.compose.runtime.Composable
import com.varabyte.kobweb.compose.css.FontStyle
import com.varabyte.kobweb.compose.css.FontWeight
import com.varabyte.kobweb.compose.css.ObjectFit
import com.varabyte.kobweb.compose.foundation.layout.Arrangement
import com.varabyte.kobweb.compose.foundation.layout.Box
import com.varabyte.kobweb.compose.foundation.layout.Column
import com.varabyte.kobweb.compose.foundation.layout.Row
import com.varabyte.kobweb.compose.ui.Alignment
import com.varabyte.kobweb.compose.ui.Modifier
import com.varabyte.kobweb.compose.ui.modifiers.*
import com.varabyte.kobweb.compose.ui.toAttrs
import com.varabyte.kobweb.silk.components.graphics.Image
import com.varabyte.kobweb.silk.components.layout.SimpleGrid
import com.varabyte.kobweb.silk.components.layout.numColumns
import com.varabyte.kobweb.silk.components.style.breakpoint.Breakpoint
import com.varabyte.kobweb.silk.theme.breakpoint.rememberBreakpoint
import org.jetbrains.compose.web.css.Color
import org.jetbrains.compose.web.css.percent
import org.jetbrains.compose.web.css.px
import org.jetbrains.compose.web.dom.P
import org.jetbrains.compose.web.dom.Text
import za.co.nofuss.components.Header
import za.co.nofuss.components.SocialBar
import za.co.nofuss.models.Section
import za.co.nofuss.models.Theme
import za.co.nofuss.util.Constants
import za.co.nofuss.util.Resources

@Composable
fun MainSection() {
    val breakpoint = rememberBreakpoint()
    Box(
        modifier = Modifier.id(Section.Home.id).maxWidth(Constants.SECTION_WIDTH.px),
        contentAlignment = Alignment.TopCenter
    ) {
        MainBackground()
        MainContent(breakpoint)
    }
}

@Composable
fun MainBackground() {
    Image(
        modifier = Modifier.fillMaxSize().objectFit(ObjectFit.Cover),
        src = Resources.Images.BACKGROUND,
        desc = "Background image"
    )
}

@Composable
fun MainContent(breakpoint: Breakpoint) {
    Column(modifier = Modifier.fillMaxSize(), horizontalAlignment = Alignment.CenterHorizontally) {
        Header()
        Column(
            modifier = Modifier.fillMaxSize(),
            verticalArrangement = Arrangement.Bottom,
            horizontalAlignment = Alignment.CenterHorizontally
        ) {
            SimpleGrid(
                modifier = Modifier.fillMaxWidth(if (breakpoint >= Breakpoint.MD) 80.percent else 90.percent),
                numColumns = numColumns(base = 1, md = 2)
            ) {
                MainText(breakpoint)
            }
        }
    }
}

@Composable
fun MainText(breakpoint: Breakpoint) {
    Row(horizontalArrangement = Arrangement.Center, verticalAlignment = Alignment.CenterVertically) {
        if (breakpoint >= Breakpoint.MD) {
            SocialBar()
        }
        Column {
            P(
                attrs = Modifier.margin(topBottom = 0.px)
                    .fontFamily(Constants.FONT_FAMILY)
                    .fontSize(35.px)
                    .fontWeight(FontWeight.Normal)
                    .color(Theme.Primary.rgb)
                    .toAttrs()
            ) { Text("Howzit! It's me") }
            P(
                attrs = Modifier.margin(top = 10.px, bottom = 0.px)
                    .fontFamily(Constants.FONT_FAMILY)
                    .fontSize(60.px)
                    .fontWeight(FontWeight.Bold)
                    .color(Color.black)
                    .toAttrs()
            ) { Text("Ewald Horn") }
            P(
                attrs = Modifier.margin(top = 10.px, bottom = 5.px)
                    .fontFamily(Constants.FONT_FAMILY)
                    .fontSize(20.px)
                    .fontWeight(FontWeight.Bold)
                    .color(Color.black)
                    .toAttrs()
            ) { Text("Freelance Software Developer") }
            P(
                attrs = Modifier.margin(bottom = 25.px)
                    .fontFamily(Constants.FONT_FAMILY)
                    .maxWidth(400.px)
                    .fontSize(15.px)
                    .fontStyle(FontStyle.Italic)
                    .fontWeight(FontWeight.Normal)
                    .color(Color.black)
                    .toAttrs()
            ) { Text("I build bespoke software solutions for mavericks. If it's esoteric, bespoke or just plain crazy, we need to talk! If it's boring, mundane and tedious, we should not!") }
        }
    }
}