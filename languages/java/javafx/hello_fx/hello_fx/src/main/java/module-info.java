module za.co.nofuss {
    requires javafx.controls;
    requires javafx.fxml;

    requires transitive javafx.graphics; 

    opens za.co.nofuss to javafx.graphics, javafx.fxml;
    exports za.co.nofuss;
}
