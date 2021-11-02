x := linspace(0, 2*pi, 20)
y1 := [sin(x) for i <- x]
y2 := [cos(x) for i <- x]
plot(x, y1, x, y2)
