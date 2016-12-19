$deers = cat input |
           %{$_ -replace "(.*) can fly (.*) km/s for (.*) seconds, but then must rest for (.*) seconds.",
             '$1,$2,$3,$4'}|ConvertFrom-Csv -Header Name,Speed,FlyTime,RestTime

$deers | %{ $_.Speed = [int]$_.Speed; $_.FlyTime = [int]$_.FlyTime; $_.RestTime = [int]$_.RestTime }

$distance = @{}; $points = @{}
foreach($i in (0..2502)){
    foreach($deer in $deers){
        if($i % ($deer.FlyTime + $deer.RestTime) -lt $deer.FlyTime){
            $distance[$deer] += $deer.Speed
        }
    }

    $currWinner = [pscustomobject]@{ Name = ""; Distance = 0 }
    foreach($deer in $deers){
        if($distance[$deer] -gt $currWinner.Distance){
            $currWinner = [pscustomobject]@{Name = $deer.Name; Distance = $distance[$deer]}
        }
    }

    $points[$currWinner.Name] += 1
}

$points.Values | measure -max | %{ "Highest points: $($_.Maximum)" }
$distance.Values | measure -max | %{ "Highest distance: $($_.Maximum)" }
