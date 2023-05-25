<?php
class SampleClass
{
	function __call($function_name, $arguments)
	{
		$count = count($arguments);

		if ($function_name == 'add') {
			if ($count == 2) {
				return array_sum($arguments);
			} else if ($count == 3) {
				return array_sum($arguments) > 10 ? 10 : array_sum($arguments);
			}
		}
	}
}

$sampleObject = new SampleClass;
echo $sampleObject->add(12, 12) . PHP_EOL; // Outputs 24 
echo $sampleObject->add(12, 2, 6) . PHP_EOL; // Outputs 10