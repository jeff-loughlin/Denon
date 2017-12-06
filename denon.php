<?php
$cmd = $_POST['cmd'];
$execute = "/usr/bin/ssh jeff@daffy /home/jeff/denon/denon $cmd";
exec($execute, $outputArray);
?>

