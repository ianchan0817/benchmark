<?php
  $current_time = microtime(true);
  echo "********* Bechmark start *********\n";
  $largest_number = 0;
  for ($x = 0; $x < 10000000; $x++) {
    if($largest_number < $x) {
      $largest_number = $x;
    }
  }
  $duration = (microtime(true)-$current_time)*100;
  echo "It takes $duration ms\n**********************************\n";

  // ********* Bechmark start *********
  // It takes 65.749716758728 ms
  // **********************************
?>
