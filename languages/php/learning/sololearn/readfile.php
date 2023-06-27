<?php
    $read = fopen('names.txt');
    foreach ($read as $line) {
        echo $line .", ";
    }
?>

<?php
    $myfile = fopen("names.txt", "w");

    $txt = "John\n";
    fwrite($myfile, $txt);
    $txt = "David\n";
    fwrite($myfile, $txt);

    fclose($myfile);

    $read = file('names.txt');
    $count = count($read);
    $i = 1;
    foreach ($read as $line) {
        echo $line;
        if($i < $count) {
            echo ', ';
        }
        $i++;
    }
?>