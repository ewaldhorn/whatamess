<?php

declare(strict_types=1);

function slices(string $digits, int $series): array
{
    $len = strlen($digits);

    if ($len < $series || $series == 0) {
        throw new Exception("No, can't help you.");
    }

    $ar = array();

    for ($index = 0; $index < $len; $index++) {
        $tmp = array();
        if ($index + $series <= $len) {
            for ($i = 0; $i < $series; $i++) {
                array_push($tmp, $digits[$index+$i]);
            }
            array_push($ar, implode($tmp));
        }
    }

    return $ar;
}

 
