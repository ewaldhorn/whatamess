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
    Articles(
        id = "articles",
        title = "Articles",
        subtitle = "",
        path = "#articles"
    ),
    Blog(
        id = "blog",
        title = "Blog",
        subtitle = "",
        path = "#blog"
    ),
    Contact(
        id = "contact",
        title = "Contact",
        subtitle = "",
        path = "#contact"
    ),
}