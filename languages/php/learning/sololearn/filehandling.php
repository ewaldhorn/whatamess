<?php
    $myfile = fopen("file.txt", "w");
?>

<?php
    $myfile = fopen("names.txt", "w");

    $txt = "Bob\n";
    fwrite($myfile, $txt);
    $txt = "Ross\n";
    fwrite($myfile, $txt);

    fclose($myfile);

    /* File contains:
        Bob
        Ross
    */
?>

<?php
    // append to file
    $myFile = "test.txt";
    $fh = fopen($myFile, 'a');
    fwrite($fh, "Some text");
    fclose($fh);
?>

<?php
if(isset($_POST['text'])) {
  $name = $_POST['text'];
  $handle = fopen('names.txt', 'a');
  fwrite($handle, $name."\n");
  fclose($handle); 
}
?>
<form method="post">
  Name: <input type="text" name="text" /> 
  <input type="submit" name="submit" />
</form>