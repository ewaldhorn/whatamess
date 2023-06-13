package za.co.nofuss.models

enum class Section(
    val id: String,
    val title: String,
    val subtitle: String,
    val path: String,
) {
    Home(
        id = "home",
        title = "Home",
        subtitle = "",
        path = "#home"
    ),
    About(
        id = "about",
        title = "About",
        subtitle = "",
        path = "#about"
    ),
    Services(
        id = "services",
        title = "Services",
        subtitle = "",
        path = "#services"
    ),
    Profile(
        id = "profile",
        title = "Profile",
        subtitle = "",
        path = "#profile"
    ),
    Contact(
        id = "contact",
        title = "Contact",
        subtitle = "",
        path = "#contact"
    ),
}