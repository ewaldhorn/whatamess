<?php
    function spitStats() {
        echo $_SERVER['SCRIPT_NAME'];
    }

    echo $_SERVER['HTTP_HOST'];
    $host = $_SERVER['HTTP_HOST'];
    $image_path = $host.'/images/';

    require 'config.php';
    echo '<img src="'.$image_path.'header.png" />';
?>