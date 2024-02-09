<?php

declare(strict_types=1);

function distance(string $strandA, string $strandB): int
{
    if (strlen($strandA) != strlen($strandB)) {
        throw new InvalidArgumentException('DNA strands must be of equal length.');
    }
    $count = 0;
    $len = strlen($strandA);

    for ($index = 0; $index < $len; $index++) {
        if ($strandA[$index] != $strandB[$index]) {
            $count++;
        }
    }

    return $count;
}

