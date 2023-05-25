<?php

$names = array(
    "Lana Oberholzer", "Willie Wikkelspies", "Abel Neville", 
    "Elen Bark", "Denzel Braddingtong", "Olivia Bazaar"
);
sort($names, SORT_NATURAL | SORT_FLAG_CASE);
foreach ($names as $key => $val) {
    echo "names[" . $key . "] = " . $val . "\n";
}
