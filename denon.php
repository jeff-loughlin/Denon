<?php
$cmd = $_POST['cmd'];
$execute = "/usr/local/bin/denon $cmd";
exec($execute, $outputArray);
echo $outputArray[0];
?>

