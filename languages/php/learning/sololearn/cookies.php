<?php
    setcookie(name, value, expire, path, domain, secure, httponly);
?>

<?php
    $value = "John";
    setcookie("user", $value, time() + (86400 * 30), '/'); 

    if (isset($_COOKIE['user'])) {
        //Outputs "Value is: John"
        echo "Value is: ". $_COOKIE['user'];
    }
    
?>