package za.co.nofuss.sections

import androidx.compose.runtime.Composable
import com.varabyte.kobweb.compose.css.*
import com.varabyte.kobweb.compose.foundation.layout.Arrangement
import com.varabyte.kobweb.compose.foundation.layout.Box
import com.varabyte.kobweb.compose.foundation.layout.Column
import com.varabyte.kobweb.compose.foundation.layout.Row
import com.varabyte.kobweb.compose.ui.Alignment
import com.varabyte.kobweb.compose.ui.Modifier
import com.varabyte.kobweb.compose.ui.graphics.Colors
import com.varabyte.kobweb.compose.ui.modifiers.*
import com.varabyte.kobweb.compose.ui.toAttrs
import com.varabyte.kobweb.silk.components.graphics.Image
import com.varabyte.kobweb.silk.components.layout.SimpleGrid
import com.varabyte.kobweb.silk.components.layout.numColumns
import com.varabyte.kobweb.silk.components.navigation.Link
import com.varabyte.kobweb.silk.components.style.breakpoint.Breakpoint
import com.varabyte.kobweb.silk.theme.breakpoint.rememberBreakpoint
import org.jetbrains.compose.web.css.Color
import org.jetbrains.compose.web.css.percent
import org.jetbrains.compose.web.css.px
import org.jetbrains.compose.web.dom.Button
import org.jetbrains.compose.web.dom.P
import org.jetbrains.compose.web.dom.Text
import za.co.nofuss.components.Header
import za.co.nofuss.components.SocialBar
import za.co.nofuss.models.Section
import za.co.nofuss.models.Theme
import za.co.nofuss.util.Constants
import za.co.nofuss.util.Links
import za.co.nofuss.util.Resources
import za.co.nofuss.util.StaticText

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
    Column(
        modifier = Modifier.fillMaxSize(),
        horizontalAlignment = Alignment.CenterHorizontally,
        verticalArrangement = Arrangement.SpaceBetween
    ) {
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
                MainImage(breakpoint)
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
                    .fontSize(if (breakpoint >= Breakpoint.LG) 35.px else 25.px)
                    .fontWeight(FontWeight.Normal)
                    .color(Theme.Primary.rgb)
                    .toAttrs()
            ) { Text("Howzit! It's me") }
            P(
                attrs = Modifier.margin(top = 10.px, bottom = 0.px)
                    .fontFamily(Constants.FONT_FAMILY)
                    .fontSize(if (breakpoint >= Breakpoint.LG) 60.px else 45.px)
                    .fontWeight(FontWeight.Bold)
                    .color(Color.black)
                    .toAttrs()
            ) { Text("Ewald Horn") }
            P(
                attrs = Modifier.margin(top = 10.px, bottom = 10.px)
                    .fontFamily(Constants.FONT_FAMILY)
                    .fontSize(if (breakpoint >= Breakpoint.LG) 20.px else 15.px)
                    .fontWeight(FontWeight.Bold)
                    .color(Color.black)
                    .toAttrs()
            ) { Text("Your next freelance software developer") }
            P(
                attrs = Modifier.margin(bottom = 25.px)
                    .fontFamily(Constants.FONT_FAMILY)
                    .maxWidth(400.px)
                    .fontSize(if (breakpoint >= Breakpoint.LG) 15.px else 12.px)
                    .fontStyle(FontStyle.Italic)
                    .fontWeight(FontWeight.Normal)
                    .color(Color.black)
                    .toAttrs()
            ) { Text(StaticText.MAIN_SECTION_SERVICE_DESCRIPTION) }
            Button(
                attrs = Modifier.height(40.px).border(width = 0.px).borderRadius(5.px)
                    .backgroundColor(Theme.Primary.rgb).color(Colors.White)
                    .toAttrs()
            ) {
                Link(
                    path = Links.BOOK_CALL,
                    modifier = Modifier.color(Colors.White)
                        .textDecorationLine(TextDecorationLine.None)
                        .fontWeight(FontWeight.Bold)
                        .padding(left = 5.px, right = 5.px),
                    text = "Book a call"
                )
            }
        }
    }
}

@Composable
fun MainImage(breakpoint: Breakpoint) {
    Column(
        modifier = Modifier.fillMaxSize(80.percent).margin(top = 50.px, bottom = 50.px),
        verticalArrangement = Arrangement.Bottom
    ) {
        Image(
            src = Resources.Images.EWALD_MAIN,
            modifier = Modifier.fillMaxWidth(85.percent).objectFit(ObjectFit.Cover),
            desc = "Ewald Horn Profile Photo"
        )
        P(
            attrs = Modifier.margin(bottom = 10.px)
                .fontFamily(Constants.FONT_FAMILY)
                .fillMaxWidth(85.percent)
                .fontSize(if (breakpoint >= Breakpoint.LG) 12.px else 10.px)
                .fontStyle(FontStyle.Italic)
                .fontWeight(FontWeight.Normal)
                .color(Color.black)
                .textAlign(TextAlign.End)
                .toAttrs()
        ) { Text("Speaking at DevConf 2023. Obviously preaching to my cult.") }

    }
}