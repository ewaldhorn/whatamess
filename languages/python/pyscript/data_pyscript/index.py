from pyscript.web import br, button, div, option, page, select, span

page.append(
    div(
        div("Hello!", classes="a-css-class", id="hello"),
        select(
            option("apple", value=1),
            option("pear", value=2),
            option("orange", value=3),
        ),
        div(
            button(span("Hello! "), span("World!"), id="my-button"),
            br(),
            button("Click me!"),
            classes=["css-class1", "css-class2"],
            style={"background-color": "red"},
        ),
        div(
            children=[
                button(children=[span("Hello! "), span("Again!")], id="another-button"),
                br(),
                button("b"),
            ],
            classes=["css-class1", "css-class2"],
        ),
    )
)
