package za.co.nofuss.styles

import com.varabyte.kobweb.compose.ui.Modifier
import com.varabyte.kobweb.compose.ui.modifiers.color
import com.varabyte.kobweb.silk.components.style.ComponentStyle
import com.varabyte.kobweb.silk.components.style.anyLink
import com.varabyte.kobweb.silk.components.style.hover
import za.co.nofuss.models.Theme

val NavigationItemStyle by ComponentStyle {
    base {
        Modifier.color(Theme.Secondary.rgb)
    }

    anyLink {
        Modifier.color(Theme.Secondary.rgb)
    }

    hover {
        Modifier.color(Theme.Primary.rgb)
    }
}