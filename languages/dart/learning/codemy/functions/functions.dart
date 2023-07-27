void main(List<String> args) {
    print("Our args are $args");

    myFunc() {
        print("\nOi! Functions? We been having it!");
    }

    anotherFunc() {
        return "This is it!";
    }

    withAValueFunc(String variable) {
        print("The variable is `$variable`.\n");
    }
    
    withOptionalValue([String name = null]) {
        if (name != null) {
            print("Hello there $name!");
        } else {
            print("Uh, ok, keep your secrets then!");
        }
    }

    withDefaultValue([String person = "Dirkie Default"]) {
        print("Yolo $person!");
    }

    myFunc();
    print(anotherFunc());
    withAValueFunc("Boo!");
    withOptionalValue();
    withOptionalValue("Ben");
    withDefaultValue();
    withDefaultValue("Derrick");
}
